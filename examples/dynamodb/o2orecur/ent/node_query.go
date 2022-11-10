// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/examples/dynamodb/o2orecur/ent/node"
	"entgo.io/ent/examples/dynamodb/o2orecur/ent/predicate"
	"entgo.io/ent/schema/field"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// NodeQuery is the builder for querying Node entities.
type NodeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Node
	withPrev   *NodeQuery
	withNext   *NodeQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
	path     func(context.Context) (*dynamodb.Selector, error)
}

// Where adds a new predicate for the NodeQuery builder.
func (nq *NodeQuery) Where(ps ...predicate.Node) *NodeQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NodeQuery) Limit(limit int) *NodeQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NodeQuery) Offset(offset int) *NodeQuery {
	nq.offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NodeQuery) Unique(unique bool) *NodeQuery {
	nq.unique = &unique
	return nq
}

// Order adds an order step to the query.
func (nq *NodeQuery) Order(o ...OrderFunc) *NodeQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryPrev chains the current query on the "prev" edge.
func (nq *NodeQuery) QueryPrev() *NodeQuery {
	query := &NodeQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *dynamodb.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.dynamodbQuery(ctx)
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(node.Table, node.FieldID, selector),
			dynamodbgraph.To(node.Table, node.FieldID, node.Keys),
			dynamodbgraph.Edge(dynamodbgraph.O2O, true, false, node.PrevAttribute, node.PrevAttribute),
		)
		fromU = dynamodbgraph.SetNeighbors(step)
		return fromU, nil
	}
	return query
}

// QueryNext chains the current query on the "next" edge.
func (nq *NodeQuery) QueryNext() *NodeQuery {
	query := &NodeQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *dynamodb.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.dynamodbQuery(ctx)
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(node.Table, node.FieldID, selector),
			dynamodbgraph.To(node.Table, node.FieldID, node.Keys),
			dynamodbgraph.Edge(dynamodbgraph.O2O, false, false, node.NextAttribute, node.NextAttribute),
		)
		fromU = dynamodbgraph.SetNeighbors(step)
		return fromU, nil
	}
	return query
}

// First returns the first Node entity from the query.
// Returns a *NotFoundError when no Node was found.
func (nq *NodeQuery) First(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{node.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NodeQuery) FirstX(ctx context.Context) *Node {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Node ID from the query.
// Returns a *NotFoundError when no Node ID was found.
func (nq *NodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{node.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NodeQuery) FirstIDX(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Node entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Node entity is found.
// Returns a *NotFoundError when no Node entities are found.
func (nq *NodeQuery) Only(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{node.Label}
	default:
		return nil, &NotSingularError{node.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NodeQuery) OnlyX(ctx context.Context) *Node {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Node ID in the query.
// Returns a *NotSingularError when more than one Node ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{node.Label}
	default:
		err = &NotSingularError{node.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nodes.
func (nq *NodeQuery) All(ctx context.Context) ([]*Node, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nq.dynamodbAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NodeQuery) AllX(ctx context.Context) []*Node {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Node IDs.
func (nq *NodeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := nq.Select(node.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NodeQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NodeQuery) Count(ctx context.Context) (int, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nq.dynamodbCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NodeQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NodeQuery) Exist(ctx context.Context) (bool, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nq.dynamodbExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NodeQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NodeQuery) Clone() *NodeQuery {
	if nq == nil {
		return nil
	}
	return &NodeQuery{
		config:     nq.config,
		limit:      nq.limit,
		offset:     nq.offset,
		order:      append([]OrderFunc{}, nq.order...),
		predicates: append([]predicate.Node{}, nq.predicates...),
		withPrev:   nq.withPrev.Clone(),
		withNext:   nq.withNext.Clone(),
		// clone intermediate query.
		dynamodb: nq.dynamodb.Clone(),
		path:     nq.path,
		unique:   nq.unique,
	}
}

// WithPrev tells the query-builder to eager-load the nodes that are connected to
// the "prev" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithPrev(opts ...func(*NodeQuery)) *NodeQuery {
	query := &NodeQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withPrev = query
	return nq
}

// WithNext tells the query-builder to eager-load the nodes that are connected to
// the "next" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithNext(opts ...func(*NodeQuery)) *NodeQuery {
	query := &NodeQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withNext = query
	return nq
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value int `json:"value,omitempty"`
//	}
//
//	client.Node.Query().
//		Select(node.FieldValue).
//		Scan(ctx, &v)
func (nq *NodeQuery) Select(fields ...string) *NodeSelect {
	nq.fields = append(nq.fields, fields...)
	return &NodeSelect{NodeQuery: nq}
}

func (nq *NodeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nq.fields {
		if !node.ValidKey(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.dynamodb = prev
	}
	return nil
}

func (nq *NodeQuery) dynamodbAll(ctx context.Context) ([]*Node, error) {
	var (
		nodes       = []*Node{}
		_node       *Node
		withFKs     = nq.withFKs
		_spec       = nq.querySpec()
		loadedTypes = [2]bool{
			nq.withPrev != nil,
			nq.withNext != nil,
		}
	)
	if nq.withPrev != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Keys = append(_spec.Node.Keys, node.ForeignKeys...)
	}

	_spec.Item = _node.item
	_spec.Assign = func(items []map[string]types.AttributeValue) error {
		for _, item := range items {
			node := &Node{}
			if err := node.FromItem(item); err != nil {
				return err
			}
			node.Edges.loadedTypes = loadedTypes
			node.config = nq.config
			nodes = append(nodes, node)
		}

		return nil
	}
	if err := dynamodbgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := nq.withPrev; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Node)
		for i := range nodes {
			if nodes[i].node_next == nil {
				continue
			}
			fk := *nodes[i].node_next
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(node.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "node_next" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Prev = n
			}
		}
	}

	if query := nq.withNext; query != nil {
		fks := make([]interface{}, 0, len(nodes))
		nodeids := make(map[int]*Node)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Node(func(s *dynamodb.Selector) {
			s.Where(dynamodb.In(node.NextAttribute, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.node_next
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "node_next" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "node_next" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Next = n
		}
	}

	return nodes, nil
}

func (nq *NodeQuery) dynamodbCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	return dynamodbgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NodeQuery) dynamodbExist(ctx context.Context) (bool, error) {
	n, err := nq.dynamodbCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (nq *NodeQuery) querySpec() *dynamodbgraph.QuerySpec {
	_spec := &dynamodbgraph.QuerySpec{
		Node: &dynamodbgraph.NodeSpec{
			Table: node.Table,
			Keys:  node.Keys,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeInt,
				Key:  node.FieldID,
			},
		},
		From: nq.dynamodb,
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *dynamodb.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *dynamodb.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NodeQuery) dynamodbQuery(ctx context.Context) *dynamodb.Selector {
	return nil
}

// NodeSelect is the builder for selecting fields of Node entities.
type NodeSelect struct {
	*NodeQuery
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NodeSelect) Scan(ctx context.Context, v any) error {
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	ns.dynamodb = ns.NodeQuery.dynamodbQuery(ctx)
	return ns.dynamodbScan(ctx, v)
}

func (ns *NodeSelect) dynamodbScan(ctx context.Context, v interface{}) error {
	return nil
}

func (ns *NodeSelect) dynamodbQuery() *dynamodb.Selector {
	return nil
}
