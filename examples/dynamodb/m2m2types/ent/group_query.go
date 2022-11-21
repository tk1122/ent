// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"strconv"

	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/examples/dynamodb/m2m2types/ent/group"
	"entgo.io/ent/examples/dynamodb/m2m2types/ent/predicate"
	"entgo.io/ent/examples/dynamodb/m2m2types/ent/user"
	"entgo.io/ent/schema/field"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// GroupQuery is the builder for querying Group entities.
type GroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	fields     []string
	predicates []predicate.Group
	withUsers  *UserQuery
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
	path     func(context.Context) (*dynamodb.Selector, error)
}

// Where adds a new predicate for the GroupQuery builder.
func (gq *GroupQuery) Where(ps ...predicate.Group) *GroupQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GroupQuery) Limit(limit int) *GroupQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GroupQuery) Offset(offset int) *GroupQuery {
	gq.offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GroupQuery) Unique(unique bool) *GroupQuery {
	gq.unique = &unique
	return gq
}

// First returns the first Group entity from the query.
// Returns a *NotFoundError when no Group was found.
func (gq *GroupQuery) First(ctx context.Context) (*Group, error) {
	nodes, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{group.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GroupQuery) FirstX(ctx context.Context) *Group {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Group ID from the query.
// Returns a *NotFoundError when no Group ID was found.
func (gq *GroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{group.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GroupQuery) FirstIDX(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Group entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Group entity is found.
// Returns a *NotFoundError when no Group entities are found.
func (gq *GroupQuery) Only(ctx context.Context) (*Group, error) {
	nodes, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{group.Label}
	default:
		return nil, &NotSingularError{group.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GroupQuery) OnlyX(ctx context.Context) *Group {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Group ID in the query.
// Returns a *NotSingularError when more than one Group ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = &NotSingularError{group.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Groups.
func (gq *GroupQuery) All(ctx context.Context) ([]*Group, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.dynamodbAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GroupQuery) AllX(ctx context.Context) []*Group {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Group IDs.
func (gq *GroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gq.Select(group.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GroupQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GroupQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.dynamodbCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GroupQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.dynamodbExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GroupQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GroupQuery) Clone() *GroupQuery {
	if gq == nil {
		return nil
	}
	return &GroupQuery{
		config:     gq.config,
		limit:      gq.limit,
		offset:     gq.offset,
		predicates: append([]predicate.Group{}, gq.predicates...),
		withUsers:  gq.withUsers.Clone(),
		// clone intermediate query.
		dynamodb: gq.dynamodb.Clone(),
		path:     gq.path,
		unique:   gq.unique,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithUsers(opts ...func(*UserQuery)) *GroupQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withUsers = query
	return gq
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Group.Query().
//		Select(group.FieldName).
//		Scan(ctx, &v)
func (gq *GroupQuery) Select(fields ...string) *GroupSelect {
	gq.fields = append(gq.fields, fields...)
	return &GroupSelect{GroupQuery: gq}
}

func (gq *GroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gq.fields {
		if !group.ValidKey(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.dynamodb = prev
	}
	return nil
}

func (gq *GroupQuery) dynamodbAll(ctx context.Context) ([]*Group, error) {
	var (
		nodes       = []*Group{}
		_node       *Group
		_spec       = gq.querySpec()
		loadedTypes = [1]bool{
			gq.withUsers != nil,
		}
	)

	_spec.Item = _node.item
	_spec.Assign = func(items []map[string]types.AttributeValue) error {
		for _, item := range items {
			node := &Group{}
			if err := node.FromItem(item); err != nil {
				return err
			}
			node.Edges.loadedTypes = loadedTypes
			node.config = gq.config
			nodes = append(nodes, node)
		}

		return nil
	}
	if err := dynamodbgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gq.withUsers; query != nil {
		var (
			edgeids []int
			edges   = make(map[int][]*Group)
		)

		for _, node := range nodes {
			node.Edges.Users = []*User{}
			edgeids = append(edgeids, node.user_id...)
			for _, id := range node.user_id {
				edges[id] = append(edges[id], node)
			}
		}

		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Users = append(nodes[i].Edges.Users, n)
			}
		}
	}

	return nodes, nil
}

func (gq *GroupQuery) dynamodbCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	return dynamodbgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GroupQuery) dynamodbExist(ctx context.Context) (bool, error) {
	n, err := gq.dynamodbCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (gq *GroupQuery) querySpec() *dynamodbgraph.QuerySpec {
	_spec := &dynamodbgraph.QuerySpec{
		Node: &dynamodbgraph.NodeSpec{
			Table: group.Table,
			Keys:  group.Keys,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeInt,
				Key:  group.FieldID,
			},
		},
		From: gq.dynamodb,
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *dynamodb.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.offset; offset != nil {
		_spec.Offset = *offset
	}
	return _spec
}

func (gq *GroupQuery) dynamodbQuery(ctx context.Context) *dynamodb.Selector {
	c1 := group.Table
	selector := dynamodb.Select(group.Keys...).From(c1)
	if gq.dynamodb != nil {
		selector = gq.dynamodb
		selector.Select(group.Keys...)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	return selector
}

// GroupSelect is the builder for selecting fields of Group entities.
type GroupSelect struct {
	*GroupQuery
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GroupSelect) Scan(ctx context.Context, v any) error {
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	gs.dynamodb = gs.GroupQuery.dynamodbQuery(ctx)
	return gs.dynamodbScan(ctx, v)
}

func (gs *GroupSelect) dynamodbScan(ctx context.Context, v interface{}) error {
	selector := gs.dynamodbQuery()
	op, args := selector.BuildExpressions().Op()
	var scanOutput sdk.ScanOutput
	if err := gs.driver.Exec(ctx, op, args, &scanOutput); err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	key := gs.fields[0]
	switch vv := v.(type) {
	case *[]string:
		for _, item := range scanOutput.Items {
			if i, ok := item[key]; ok {
				if v, ok := i.(*types.AttributeValueMemberS); ok {
					*vv = append(*vv, v.Value)
				}
			}
		}
	case *[]int:
		for _, item := range scanOutput.Items {
			if i, ok := item[key]; ok {
				if v, ok := i.(*types.AttributeValueMemberN); ok {
					num, err := strconv.Atoi(v.Value)
					if err == nil {
						*vv = append(*vv, num)
					}
				}
			}
		}
	case *[]float64:
		for _, item := range scanOutput.Items {
			if i, ok := item[key]; ok {
				if v, ok := i.(*types.AttributeValueMemberN); ok {
					num, err := strconv.ParseFloat(v.Value, 64)
					if err == nil {
						*vv = append(*vv, num)
					}
				}
			}
		}
	case *[]bool:
		for _, item := range scanOutput.Items {
			if i, ok := item[key]; ok {
				if v, ok := i.(*types.AttributeValueMemberBOOL); ok {
					*vv = append(*vv, v.Value)
				}
			}
		}
	}
	return nil
}

func (gs *GroupSelect) dynamodbQuery() *dynamodb.Selector {
	selector := gs.dynamodb
	selector.Select(gs.fields...)
	return selector
}
