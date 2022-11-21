// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dynamodbgraph

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	sdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/schema/field"
)

type (
	// CreateSpec holds the information for creating
	// a node in the graph.
	CreateSpec struct {
		Table  string
		ID     *FieldSpec
		Fields []*FieldSpec
		Edges  []*EdgeSpec
	}

	// FieldSpec holds the information for updating a field in the database.
	FieldSpec struct {
		Key   string
		Type  field.Type
		Value interface{}
	}

	// EdgeTarget holds the information for the target nodes of an edge.
	EdgeTarget struct {
		Nodes  []interface{}
		IDSpec *FieldSpec
	}

	// EdgeSpec holds the information for updating a field in the database.
	EdgeSpec struct {
		Rel        Rel
		Inverse    bool
		Table      string
		Attributes []string
		Bidi       bool        // bidirectional edge.
		Target     *EdgeTarget // target nodes.
	}

	// EdgeSpecs used for perform common operations on list of edges.
	EdgeSpecs []*EdgeSpec

	// NodeSpec defines the information for querying and
	// decoding nodes in the graph.
	NodeSpec struct {
		Table string
		Keys  []string
		ID    *FieldSpec
	}

	// Rel is a relation type of edge.
	Rel int

	// BatchCreateSpec holds the information for creating
	// multiple nodes in the graph.
	BatchCreateSpec struct {
		Nodes []*CreateSpec
	}
)

// Relation types.
const (
	_   Rel = iota // Unknown.
	O2O            // One to one / has one.
	O2M            // One to many / has many.
	M2O            // Many to one (inverse perspective for O2M).
	M2M            // Many to many.
)

// String returns the relation name.
func (r Rel) String() (s string) {
	switch r {
	case O2O:
		s = "O2O"
	case O2M:
		s = "O2M"
	case M2O:
		s = "M2O"
	case M2M:
		s = "M2M"
	default:
		s = "Unknown"
	}
	return s
}

type (
	graph struct {
		tx dialect.ExecQuerier
		dynamodb.RootBuilder
	}

	creator struct {
		graph
		*CreateSpec
		*BatchCreateSpec
	}
)

// CreateNode applies the CreateSpec on the graph.
func CreateNode(ctx context.Context, drv dialect.Driver, spec *CreateSpec) (err error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return err
	}
	gr := graph{tx: tx, RootBuilder: dynamodb.RootBuilder{}}
	cr := &creator{
		CreateSpec: spec,
		graph:      gr,
	}
	if err = cr.node(ctx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// BatchCreate applies the BatchCreateSpec on the graph.
func BatchCreate(ctx context.Context, drv dialect.Driver, spec *BatchCreateSpec) error {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return err
	}
	gr := graph{tx: tx}
	cr := &creator{BatchCreateSpec: spec, graph: gr}
	if err := cr.nodes(ctx, tx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// node is the controller to create a single node in the graph.
func (c *creator) node(ctx context.Context) (err error) {
	batchWrite := dynamodb.BatchWriteItem()
	edges, fields := EdgeSpecs(c.CreateSpec.Edges).GroupRel(), c.CreateSpec.Fields
	itemData := make(map[string]types.AttributeValue)
	// ID field is not included in CreateSpec.Fields
	if c.CreateSpec.ID != nil {
		fields = append(fields, c.CreateSpec.ID)
	}
	if err = c.insert(ctx, c.CreateSpec.Table, batchWrite, fields, edges, itemData); err != nil {
		return err
	}
	if err = c.graph.addFKEdges(ctx, []interface{}{c.ID.Value}, append(edges[O2M], edges[O2O]...)); err != nil {
		return err
	}
	if err = c.graph.addM2MEdges(ctx, []interface{}{c.ID.Value}, edges[M2M], batchWrite); err != nil {
		return err
	}
	op, input := batchWrite.Op()
	if err := c.tx.Exec(ctx, op, input, &sdk.BatchWriteItemOutput{}); err != nil {
		return fmt.Errorf("create node failed: %w", err)
	}
	return nil
}

func (c *creator) nodes(ctx context.Context, tx dialect.ExecQuerier) error {
	if len(c.Nodes) == 0 {
		return nil
	}
	batchWrite := dynamodb.BatchWriteItem()
	for i, node := range c.Nodes {
		if i > 0 && node.Table != c.Nodes[i-1].Table {
			return fmt.Errorf("more than 1 table for batch insert: %q != %q", node.Table, c.Nodes[i-1].Table)
		}
		if node.ID.Value == nil {
			return fmt.Errorf("the value of id field is required")
		}
		edges, fields := EdgeSpecs(node.Edges).GroupRel(), node.Fields
		itemData := make(map[string]types.AttributeValue)
		// ID field is not included in CreateSpec.Fields
		fields = append(fields, node.ID)
		if err := c.insert(ctx, node.Table, batchWrite, fields, edges, itemData); err != nil {
			return err
		}
	}

	for _, node := range c.Nodes {
		edges := EdgeSpecs(node.Edges).GroupRel()
		if err := c.graph.addM2MEdges(ctx, []interface{}{node.ID.Value}, edges[M2M], batchWrite); err != nil {
			return err
		}
		if err := c.graph.addFKEdges(ctx, []interface{}{node.ID.Value}, append(edges[O2M], edges[O2O]...)); err != nil {
			return err
		}
	}

	op, args := batchWrite.Op()
	if err := tx.Exec(ctx, op, args, &sdk.BatchWriteItemOutput{}); err != nil {
		return fmt.Errorf("create nodes failed: %w", err)
	}

	return nil
}

// insert put an item with all attributes it owns to the table,
// it does not handle adding foreign keys fields in other tables.
func (c *creator) insert(
	ctx context.Context,
	table string,
	batchWrite *dynamodb.BatchWriteItemBuilder,
	fields []*FieldSpec,
	edges map[Rel][]*EdgeSpec,
	itemData map[string]types.AttributeValue) (err error) {
	putItemBuilder := c.PutItem(table)
	if err = c.setItemAttributes(fields, edges, itemData); err != nil {
		return err
	}
	putItemBuilder.SetItem(itemData)
	batchWrite.Append(table, putItemBuilder)
	return nil
}

func (c *creator) setItemAttributes(fields []*FieldSpec, edges map[Rel][]*EdgeSpec, itemData map[string]types.AttributeValue) (err error) {
	for _, f := range fields {
		if itemData[f.Key], err = attributevalue.Marshal(f.Value); err != nil {
			return err
		}
	}
	for _, e := range edges[M2O] {
		if itemData[e.Attributes[0]], err = attributevalue.Marshal(e.Target.Nodes[0]); err != nil {
			return err
		}
	}
	for _, e := range edges[O2O] {
		if e.Inverse || e.Bidi {
			if itemData[e.Attributes[0]], err = attributevalue.Marshal(e.Target.Nodes[0]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *graph) addFKEdges(ctx context.Context, ids []interface{}, edges []*EdgeSpec) (err error) {
	if len(ids) > 1 && len(edges) != 0 {
		// O2M and O2O edges are defined by a FK in the "other" collection.
		// Therefore, ids[i+1] will override ids[i] which is invalid.
		return fmt.Errorf("unable to link FK edge to more than 1 node: %v", ids)
	}
	id := ids[0]
	for _, edge := range edges {
		if edge.Rel == O2O && edge.Inverse {
			continue
		}
		for _, n := range edge.Target.Nodes {
			keyVal, err := attributevalue.Marshal(n)
			if err != nil {
				return fmt.Errorf("key type not supported: %v has type %T", n, n)
			}
			query, err := g.Update(edge.Table).
				WithKey(edge.Target.IDSpec.Key, keyVal).
				Set(edge.Attributes[0], id).
				BuildExpression(types.ReturnValueAllNew)
			if err != nil {
				return fmt.Errorf("build update query for table %s: %w", edge.Table, err)
			}
			op, input := query.Op()
			var output sdk.UpdateItemOutput
			if err := g.tx.Exec(ctx, op, input, &output); err != nil {
				return fmt.Errorf("add %s edge for table %s: %w", edge.Rel, edge.Table, err)
			}
		}
	}
	return nil
}

func (g *graph) clearFKEdges(ctx context.Context, ids []interface{}, edges []*EdgeSpec) error {
	id := ids[0]
	for _, edge := range edges {
		if edge.Rel == O2O && edge.Inverse {
			continue
		}
		var inverseQuery *dynamodb.Selector
		if (edge.Rel == O2O && edge.Bidi) || edge.Rel == O2M {
			inverseQuery = dynamodb.Select().From(edge.Table).Where(dynamodb.EQ(edge.Attributes[0], id)).BuildExpressions()
		} else {
			inverseQuery = dynamodb.Select().From(edge.Table).Where(dynamodb.EQ(edge.Target.IDSpec.Key, id)).BuildExpressions()
		}
		op, input := inverseQuery.Op()
		var output sdk.ScanOutput
		if err := g.tx.Exec(ctx, op, input, &output); err != nil {
			return fmt.Errorf("remove %s edge for table %s: %w", edge.Rel, edge.Table, err)
		}
		for _, item := range output.Items {
			query, err := g.Update(edge.Table).
				WithKey(edge.Target.IDSpec.Key, item[edge.Target.IDSpec.Key]).
				Remove(edge.Attributes[0]).
				BuildExpression(types.ReturnValueAllNew)
			if err != nil {
				return fmt.Errorf("remove %s edge for table %s: %w", edge.Rel, edge.Table, err)
			}
			op, input := query.Op()
			var output sdk.UpdateItemOutput
			if err := g.tx.Exec(ctx, op, input, &output); err != nil {
				return fmt.Errorf("remove %s edge for table %s: %w", edge.Rel, edge.Table, err)
			}
		}

	}
	return nil
}

func (g *graph) addM2MEdges(ctx context.Context, ids []interface{}, edges []*EdgeSpec, batchWrite *dynamodb.BatchWriteItemBuilder) (err error) {
	if len(edges) == 0 {
		return nil
	}
	for _, e := range edges {
		m2mTable := e.Table
		fromIds, toIds := e.Target.Nodes, ids
		if e.Inverse {
			fromIds, toIds = toIds, fromIds
		}
		toAttr, fromAttr := e.Attributes[0], e.Attributes[1]
		for _, fromId := range fromIds {
			for _, toId := range toIds {
				item := make(map[string]types.AttributeValue)
				if item[toAttr], err = attributevalue.Marshal(toId); err != nil {
					return fmt.Errorf("add m2m edge: %w", err)
				}
				if item[fromAttr], err = attributevalue.Marshal(fromId); err != nil {
					return fmt.Errorf("add m2m edge: %w", err)
				}
				batchWrite.Append(m2mTable, g.PutItem(m2mTable).SetItem(item))
				if e.Bidi {
					reverseItem := make(map[string]types.AttributeValue)
					if reverseItem[toAttr], err = attributevalue.Marshal(fromId); err != nil {
						return fmt.Errorf("add m2m edge: %w", err)
					}
					if reverseItem[fromAttr], err = attributevalue.Marshal(toId); err != nil {
						return fmt.Errorf("add m2m edge: %w", err)
					}
					batchWrite.Append(m2mTable, g.PutItem(m2mTable).SetItem(reverseItem))
				}
			}
		}
	}
	return nil
}

func (g *graph) clearM2MEdges(ctx context.Context, ids []interface{}, edges []*EdgeSpec, batchWrite *dynamodb.BatchWriteItemBuilder) (err error) {
	if len(edges) == 0 {
		return nil
	}
	for _, e := range edges {
		m2mTable := e.Table
		id := ids[0]
		partitionKey, sortKey := e.Attributes[0], e.Attributes[1]
		var joinTableQuery *dynamodb.Selector
		if e.Bidi {
			joinTableQuery = dynamodb.Select().From(e.Table).Where(dynamodb.Or(dynamodb.EQ(partitionKey, id), dynamodb.EQ(sortKey, id))).BuildExpressions()
		} else {
			filterAttr := partitionKey
			if e.Inverse {
				filterAttr = sortKey
			}
			joinTableQuery = dynamodb.Select().From(e.Table).Where(dynamodb.EQ(filterAttr, id)).BuildExpressions()
		}
		op, input := joinTableQuery.Op()
		var output sdk.ScanOutput
		if err := g.tx.Exec(ctx, op, input, &output); err != nil {
			return fmt.Errorf("remove %s edge for table %s: %w", e.Rel, e.Table, err)
		}
		for _, item := range output.Items {
			batchWrite.Append(m2mTable, g.DeleteItem(m2mTable).
				WithKey(partitionKey, item[partitionKey]).
				WithKey(sortKey, item[sortKey]))
		}
	}
	return nil
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx dialect.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%s: %v", err.Error(), rerr)
	}
	return err
}

// GroupRel groups edges by their relation type.
func (es EdgeSpecs) GroupRel() map[Rel][]*EdgeSpec {
	edges := make(map[Rel][]*EdgeSpec)
	for _, edge := range es {
		edges[edge.Rel] = append(edges[edge.Rel], edge)
	}
	return edges
}

// NewStep gets list of options and returns a configured step.
//
//	NewStep(
//		From("table", "id", V),
//		To("table", "id"),
//		Edge("name", O2M, "ref_id"),
//	)
func NewStep(opts ...StepOption) *Step {
	s := &Step{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// StepOption allows configuring Steps using functional options.
type StepOption func(*Step)

// A Step provides a path-step information to the traversal functions.
type Step struct {
	// From is the source of the step.
	From struct {
		// V can be either one vertex or set of vertices.
		// It can be a pre-processed step (dynamodb.Selector) or a simple Go type (integer or string).
		V interface{}
		// Table holds the collection name of V (from).
		Table string
		// Attribute to join with. Usually the "id" key.
		Attribute string
	}
	// Edge holds the edge information for getting the neighbors.
	Edge struct {
		// Rel of the edge.
		Rel Rel
		// Table name of where this edge keys reside.
		Table string
		// Attributes of the edge.
		// In O2O and M2O, it holds the foreign-key field. Hence, len == 1.
		// In M2M, it holds the primary-key keys of the join collection. Hence, len == 2.
		Attributes []string
		// Inverse indicates if the edge is an inverse edge.
		Inverse bool
		// Bidi indicates if this edge is a bidirectional edge. A self-reference
		// to the same type with the same name (symmetric relation). For example,
		// a User type have one of following edges:
		//
		//	edge.To("friends", User.Type)           // many 2 many.
		//	edge.To("spouse", User.Type).Unique()   // one 2 one.
		//
		Bidi bool
	}
	// To is the dest of the path (the neighbors).
	To struct {
		// Table holds the table name of the neighbors (to).
		Table string
		// Attribute to join with. Usually the "id" key.
		Attribute string
		// CollectionKeys holds all keys of the collection.
		// Note: Only use in HasNeighbors/HasNeighborsWith.
		CollectionKeys []string
	}
}

// From sets the source of the step.
func From(table, attr string, v ...interface{}) StepOption {
	return func(s *Step) {
		s.From.Table = table
		s.From.Attribute = attr
		if len(v) > 0 {
			s.From.V = v[0]
		}
	}
}

// To sets the destination of the step.
func To(table, attr string, collectionKeys []string) StepOption {
	return func(s *Step) {
		s.To.Table = table
		s.To.Attribute = attr
		s.To.CollectionKeys = collectionKeys
	}
}

// Edge sets the edge info for getting the neighbors.
func Edge(rel Rel, inverse bool, bidi bool, table string, attrs ...string) StepOption {
	return func(s *Step) {
		s.Edge.Rel = rel
		s.Edge.Table = table
		s.Edge.Attributes = attrs
		s.Edge.Inverse = inverse
		s.Edge.Bidi = bidi
	}
}

// Neighbors returns a Selector for evaluating the path-step
// and getting the neighbors of one vertex.
func Neighbors(s *Step, drv dialect.Driver) (q *dynamodb.Selector) {
	ctx := context.TODO()
	switch r := s.Edge.Rel; {
	case r == M2M && s.Edge.Bidi:
		pred := dynamodb.EQ(s.Edge.Attributes[0], s.From.V)
		joinTableQuery := dynamodb.Select(s.Edge.Attributes...).
			From(s.Edge.Table).
			Where(pred).
			BuildExpressions()
		op, args := joinTableQuery.Op()
		var output sdk.ScanOutput
		if err := drv.Query(ctx, op, args, &output); err != nil {
			q.AddError(err)
			return q
		}
		var ids []interface{}
		for _, i := range output.Items {
			ids = append(ids, i[s.Edge.Attributes[1]])
		}
		q = dynamodb.Select().
			From(s.To.Table).
			Where(dynamodb.In(s.To.Attribute, ids...))

	case r == M2M && s.Edge.Inverse:
		joinTableQuery := dynamodb.Select(s.Edge.Attributes[0]).
			From(s.Edge.Table).
			Where(dynamodb.EQ(s.Edge.Attributes[1], s.From.V)).
			BuildExpressions()
		op, args := joinTableQuery.Op()
		var output sdk.ScanOutput
		if err := drv.Query(ctx, op, args, &output); err != nil {
			q.AddError(err)
			return q
		}
		var ids []interface{}
		for _, i := range output.Items {
			ids = append(ids, i[s.Edge.Attributes[0]])
		}
		q = dynamodb.Select().
			From(s.To.Table).
			Where(dynamodb.In(s.To.Attribute, ids...))

	case r == M2M && !s.Edge.Inverse:
		joinTableQuery := dynamodb.Select(s.Edge.Attributes[1]).
			From(s.Edge.Table).
			Where(dynamodb.EQ(s.Edge.Attributes[0], s.From.V)).
			BuildExpressions()
		op, args := joinTableQuery.Op()
		var output sdk.ScanOutput
		if err := drv.Query(ctx, op, args, &output); err != nil {
			q.AddError(err)
			return q
		}
		var ids []interface{}
		for _, i := range output.Items {
			ids = append(ids, i[s.Edge.Attributes[1]])
		}
		q = dynamodb.Select().
			From(s.To.Table).
			Where(dynamodb.In(s.To.Attribute, ids...))

	case r == M2O || (r == O2O && s.Edge.Inverse):
		q = dynamodb.Select().
			From(s.To.Table)
		iq := dynamodb.Select(s.Edge.Attributes[0]).
			From(s.Edge.Table).
			Where(dynamodb.EQ(s.From.Attribute, s.From.V)).
			BuildExpressions()
		op, args := iq.Op()
		var output sdk.ScanOutput
		if err := drv.Query(ctx, op, args, &output); err != nil {
			q.AddError(err)
			return q
		}
		q.Where(dynamodb.EQ(s.To.Attribute, output.Items[0][s.Edge.Attributes[0]]))

	case r == O2M || (r == O2O && !s.Edge.Inverse):
		q = dynamodb.Select().
			From(s.Edge.Table).
			Where(dynamodb.EQ(s.Edge.Attributes[0], s.From.V))
	}
	return q
}

// QueryNodes queries the nodes in the graph query and scans them to the given values.
func QueryNodes(ctx context.Context, drv dialect.Driver, spec *QuerySpec) error {
	qr := &query{graph: graph{}, QuerySpec: spec}
	return qr.nodes(ctx, drv)
}

type query struct {
	graph
	*QuerySpec
}

// QuerySpec holds the information for querying
// nodes in the graph.
type QuerySpec struct {
	Node *NodeSpec          // Nodes info.
	From *dynamodb.Selector // Optional query source (from path).

	Limit     int
	Offset    int
	Order     func(*dynamodb.Selector)
	Predicate func(*dynamodb.Selector)

	Item   func() interface{}
	Assign func([]map[string]types.AttributeValue) error
}

func (q *query) nodes(ctx context.Context, drv dialect.Driver) error {
	selector, err := q.selector()
	if err != nil {
		return err
	}
	op, args := selector.Op()
	var output sdk.ScanOutput
	if err := drv.Query(ctx, op, args, &output); err != nil {
		return err
	}
	return q.Assign(output.Items)
}

func (q *query) selector() (*dynamodb.Selector, error) {
	selector := dynamodb.Select().From(q.Node.Table)
	if q.From != nil {
		selector = q.From
	}
	if pred := q.Predicate; pred != nil {
		pred(selector)
	}
	if order := q.Order; order != nil {
		order(selector)
	}
	selector.BuildExpressions()
	return selector, selector.Err()
}

// CountNodes counts the nodes in the given graph query.
func CountNodes(ctx context.Context, drv dialect.Driver, spec *QuerySpec) (int, error) {
	qr := &query{graph: graph{}, QuerySpec: spec}
	return qr.count(ctx, drv)
}

func (q *query) count(ctx context.Context, drv dialect.Driver) (int, error) {
	selector, err := q.selector()
	if err != nil {
		return 0, err
	}
	op, args := selector.Op()
	var output sdk.ScanOutput
	if err := drv.Query(ctx, op, args, &output); err != nil {
		return 0, err
	}
	return int(output.Count), nil
}

type (
	// EdgeMut defines edge mutations.
	EdgeMut struct {
		Add   []*EdgeSpec
		Clear []*EdgeSpec
	}

	// FieldMut defines field mutations.
	FieldMut struct {
		Set   []*FieldSpec // field = ?
		Add   []*FieldSpec // field = field + ?
		Clear []*FieldSpec // field = NULL
	}

	// UpdateSpec holds the information for updating one
	// or more nodes in the graph.
	UpdateSpec struct {
		Node      *NodeSpec
		Edges     EdgeMut
		Fields    FieldMut
		Predicate func(*dynamodb.Selector)

		Item   func() interface{}
		Assign func(interface{}) error
	}
)

// UpdateNode applies the UpdateSpec on one node in the graph.
func UpdateNode(ctx context.Context, drv dialect.Driver, spec *UpdateSpec) error {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return err
	}
	gr := graph{tx: tx}
	cr := &updater{UpdateSpec: spec, graph: gr}
	if err := cr.node(ctx, tx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// UpdateNodes applies the UpdateSpec on a set of nodes in the graph.
func UpdateNodes(ctx context.Context, drv dialect.Driver, spec *UpdateSpec) (int, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return 0, err
	}
	gr := graph{tx: tx}
	cr := &updater{UpdateSpec: spec, graph: gr}
	affected, err := cr.nodes(ctx, tx)
	if err != nil {
		return 0, rollback(tx, err)
	}
	return affected, tx.Commit()
}

// NotFoundError returns when trying to update an
// entity and it was not found in the database.
type NotFoundError struct {
	table string
	id    interface{}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("record with id %v not found in table %s", e.id, e.table)
}

type updater struct {
	graph
	*UpdateSpec
}

func (u *updater) node(ctx context.Context, tx dialect.ExecQuerier) (err error) {
	updateItemBuilder, batchWrite := u.Update(u.Node.Table), dynamodb.BatchWriteItem()
	addEdges, clearEdges := EdgeSpecs(u.Edges.Add).GroupRel(), EdgeSpecs(u.Edges.Clear).GroupRel()
	idValue, err := attributevalue.Marshal(u.Node.ID.Value)
	if err != nil {
		return fmt.Errorf("update node failed: %w", err)
	}
	updateItemBuilder.WithKey(u.Node.ID.Key, idValue)
	if err = u.update(ctx, u.Node.ID.Value, addEdges, clearEdges, updateItemBuilder, batchWrite); err != nil {
		return fmt.Errorf("update node failed: %w", err)
	}
	updateItemBuilder, err = updateItemBuilder.BuildExpression(types.ReturnValueAllNew)
	if err != nil {
		return fmt.Errorf("update node failed: %w", err)
	}
	op, args := updateItemBuilder.Op()
	var output sdk.UpdateItemOutput
	err = tx.Exec(ctx, op, args, &output)
	if err != nil {
		return fmt.Errorf("update node failed: %w", err)
	}
	if !batchWrite.IsEmpty {
		op, args := batchWrite.Op()
		if err := tx.Exec(ctx, op, args, &sdk.BatchWriteItemOutput{}); err != nil {
			return fmt.Errorf("update node failed: %w", err)
		}
	}
	return u.Assign(output.Attributes)
}

func (u *updater) update(
	ctx context.Context,
	id interface{},
	addEdges map[Rel][]*EdgeSpec,
	clearEdges map[Rel][]*EdgeSpec,
	updateBuilder *dynamodb.UpdateItemBuilder,
	batchWrite *dynamodb.BatchWriteItemBuilder) (err error) {
	u.setAddAttributesAndEdges(ctx, u.Fields.Set, addEdges, updateBuilder)
	u.setClearAttributesAndEdges(ctx, u.Fields.Clear, clearEdges, updateBuilder)
	if err := u.graph.clearFKEdges(ctx, []interface{}{id}, append(clearEdges[O2M], clearEdges[O2O]...)); err != nil {
		return fmt.Errorf("update node failed: %w", err)
	}
	if err := u.graph.addFKEdges(ctx, []interface{}{id}, append(addEdges[O2M], addEdges[O2O]...)); err != nil {
		return fmt.Errorf("update node failed: %w", err)
	}
	if err := u.graph.clearM2MEdges(ctx, []interface{}{id}, clearEdges[M2M], batchWrite); err != nil {
		return fmt.Errorf("update node failed: %w", err)
	}
	if err := u.graph.addM2MEdges(ctx, []interface{}{id}, addEdges[M2M], batchWrite); err != nil {
		return fmt.Errorf("update node failed: %w", err)
	}
	return nil
}

func (u *updater) nodes(ctx context.Context, tx dialect.ExecQuerier) (int, error) {
	selector := dynamodb.Select(u.Node.ID.Key).
		From(u.Node.Table)
	if pred := u.Predicate; pred != nil {
		pred(selector)
	}
	op, args := selector.BuildExpressions().Op()
	var scanOutput sdk.ScanOutput
	if err := tx.Exec(ctx, op, args, &scanOutput); err != nil {
		return 0, fmt.Errorf("update nodes failed: %w", err)
	}
	if len(scanOutput.Items) == 0 {
		return 0, nil
	}
	batchWrite := dynamodb.BatchWriteItem()
	for _, item := range scanOutput.Items {
		addEdges, clearEdges := EdgeSpecs(u.Edges.Add).GroupRel(), EdgeSpecs(u.Edges.Clear).GroupRel()
		updateItemBuilder := u.Update(u.Node.Table).WithKey(u.Node.ID.Key, item[u.Node.ID.Key])
		var id interface{}
		err := attributevalue.Unmarshal(item[u.Node.ID.Key], &id)
		if err != nil {
			return 0, fmt.Errorf("update nodes failed: %w", err)
		}
		if err := u.update(ctx, id, addEdges, clearEdges, updateItemBuilder, batchWrite); err != nil {
			return 0, fmt.Errorf("update nodes failed: %w", err)
		}
		updateItemBuilder, err = updateItemBuilder.BuildExpression(types.ReturnValueAllNew)
		if err != nil {
			return 0, fmt.Errorf("update nodes failed: %w", err)
		}
		op, args := updateItemBuilder.Op()
		err = tx.Exec(ctx, op, args, &sdk.UpdateItemOutput{})
		if err != nil {
			return 0, fmt.Errorf("update nodes failed: %w", err)
		}
	}
	if batchWrite.IsEmpty {
		return 0, nil
	}
	op, args = batchWrite.Op()
	if err := u.tx.Exec(ctx, op, args, &sdk.BatchWriteItemOutput{}); err != nil {
		return 0, fmt.Errorf("update nodes failed: %w", err)
	}
	return len(scanOutput.Items), nil
}

func (u *updater) setAddAttributesAndEdges(ctx context.Context, fields []*FieldSpec, addEdges map[Rel][]*EdgeSpec, update *dynamodb.UpdateItemBuilder) {
	for _, f := range fields {
		update.Set(f.Key, f.Value)
	}
	for _, e := range addEdges[M2O] {
		update.Set(e.Attributes[0], e.Target.Nodes[0])
	}
	for _, e := range addEdges[O2O] {
		if e.Inverse || e.Bidi {
			update.Set(e.Attributes[0], e.Target.Nodes[0])
		}
	}
}

func (u *updater) setClearAttributesAndEdges(ctx context.Context, fields []*FieldSpec, clearEdges map[Rel][]*EdgeSpec, update *dynamodb.UpdateItemBuilder) {
	for _, f := range fields {
		update.Remove(f.Key)
	}
	for _, e := range clearEdges[M2O] {
		update.Remove(e.Attributes[0])
	}
	for _, e := range clearEdges[O2O] {
		if e.Inverse || e.Bidi {
			update.Remove(e.Attributes[0])
		}
	}
}

// DeleteSpec holds the information for delete one
// or more nodes in the graph.
type DeleteSpec struct {
	Node       *NodeSpec
	Predicate  func(*dynamodb.Selector)
	ClearEdges []*EdgeSpec
}

type deleter struct {
	graph
	*DeleteSpec
}

// DeleteNodes applies the DeleteSpec on the graph.
func DeleteNodes(ctx context.Context, drv dialect.Driver, spec *DeleteSpec) (int, error) {
	tx, err := drv.Tx(ctx)
	gr := graph{tx: tx}
	dl := deleter{
		graph:      gr,
		DeleteSpec: spec,
	}
	if err != nil {
		return 0, err
	}
	id := spec.Node.ID.Value
	batchWrite := dynamodb.BatchWriteItem()
	selector := dynamodb.Select().From(spec.Node.Table)
	if pred := spec.Predicate; pred != nil {
		pred(selector)
	}
	if id != nil {
		return dl.deleteNode(ctx, id, selector, batchWrite)
	}
	return dl.deleteNodes(ctx, selector, batchWrite)
}

func (dl *deleter) deleteNode(ctx context.Context, id interface{}, selector *dynamodb.Selector, batchWrite *dynamodb.BatchWriteItemBuilder) (int, error) {
	keyVal, err := attributevalue.Marshal(id)
	if err != nil {
		return 0, fmt.Errorf("delete node failed: %w", err)
	}
	deleteItem := dl.DeleteItem(dl.DeleteSpec.Node.Table).Where(selector.P()).WithKey(dl.DeleteSpec.Node.ID.Key, keyVal)
	batchWrite.Append(dl.DeleteSpec.Node.Table, deleteItem)
	clearEdges := EdgeSpecs(dl.DeleteSpec.ClearEdges).GroupRel()
	if err := dl.graph.clearM2MEdges(ctx, []interface{}{id}, clearEdges[M2M], batchWrite); err != nil {
		return 0, fmt.Errorf("delete node failed: %w", err)
	}
	if batchWrite.IsEmpty {
		return 0, fmt.Errorf("delete node failed: %w", err)
	}
	op, args := batchWrite.Op()
	if err := dl.tx.Exec(ctx, op, args, &sdk.DeleteItemOutput{}); err != nil {
		return 0, fmt.Errorf("delete node failed: %w", err)
	}
	return 1, nil
}

func (dl *deleter) deleteNodes(ctx context.Context, selector *dynamodb.Selector, batchWrite *dynamodb.BatchWriteItemBuilder) (int, error) {
	op, args := dl.Select().From(dl.DeleteSpec.Node.Table).Where(selector.P()).BuildExpressions().Op()
	var scanOutput sdk.ScanOutput
	if err := dl.tx.Exec(ctx, op, args, &scanOutput); err != nil {
		return 0, fmt.Errorf("delete nodes failed: %w", err)
	}
	if len(scanOutput.Items) == 0 {
		return 0, nil
	}
	for _, item := range scanOutput.Items {
		deleteItem := dl.DeleteItem(dl.DeleteSpec.Node.Table).WithKey(dl.DeleteSpec.Node.ID.Key, item[dl.DeleteSpec.Node.ID.Key])
		batchWrite.Append(dl.DeleteSpec.Node.Table, deleteItem)
	}
	if batchWrite.IsEmpty {
		return 0, nil
	}
	op, args = batchWrite.Op()
	if err := dl.tx.Exec(ctx, op, args, &sdk.BatchWriteItemOutput{}); err != nil {
		return 0, fmt.Errorf("delete nodes failed: %w", err)
	}
	return len(scanOutput.Items), nil
}
