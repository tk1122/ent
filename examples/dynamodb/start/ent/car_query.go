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
	"entgo.io/ent/examples/dynamodb/start/ent/car"
	"entgo.io/ent/examples/dynamodb/start/ent/predicate"
	"entgo.io/ent/examples/dynamodb/start/ent/user"
	"entgo.io/ent/schema/field"
	sdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	uuid "github.com/satori/go.uuid"
)

// CarQuery is the builder for querying Car entities.
type CarQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	fields     []string
	predicates []predicate.Car
	withOwner  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
	path     func(context.Context) (*dynamodb.Selector, error)
}

// Where adds a new predicate for the CarQuery builder.
func (cq *CarQuery) Where(ps ...predicate.Car) *CarQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CarQuery) Limit(limit int) *CarQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CarQuery) Offset(offset int) *CarQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CarQuery) Unique(unique bool) *CarQuery {
	cq.unique = &unique
	return cq
}

// First returns the first Car entity from the query.
// Returns a *NotFoundError when no Car was found.
func (cq *CarQuery) First(ctx context.Context) (*Car, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{car.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CarQuery) FirstX(ctx context.Context) *Car {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Car ID from the query.
// Returns a *NotFoundError when no Car ID was found.
func (cq *CarQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{car.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CarQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Car entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Car entity is found.
// Returns a *NotFoundError when no Car entities are found.
func (cq *CarQuery) Only(ctx context.Context) (*Car, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{car.Label}
	default:
		return nil, &NotSingularError{car.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CarQuery) OnlyX(ctx context.Context) *Car {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Car ID in the query.
// Returns a *NotSingularError when more than one Car ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CarQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{car.Label}
	default:
		err = &NotSingularError{car.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CarQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cars.
func (cq *CarQuery) All(ctx context.Context) ([]*Car, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.dynamodbAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CarQuery) AllX(ctx context.Context) []*Car {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Car IDs.
func (cq *CarQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := cq.Select(car.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CarQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CarQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.dynamodbCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CarQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CarQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.dynamodbExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CarQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CarQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CarQuery) Clone() *CarQuery {
	if cq == nil {
		return nil
	}
	return &CarQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		predicates: append([]predicate.Car{}, cq.predicates...),
		withOwner:  cq.withOwner.Clone(),
		// clone intermediate query.
		dynamodb: cq.dynamodb.Clone(),
		path:     cq.path,
		unique:   cq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CarQuery) WithOwner(opts ...func(*UserQuery)) *CarQuery {
	query := &UserQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withOwner = query
	return cq
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Model string `json:"model,omitempty"`
//	}
//
//	client.Car.Query().
//		Select(car.FieldModel).
//		Scan(ctx, &v)
func (cq *CarQuery) Select(fields ...string) *CarSelect {
	cq.fields = append(cq.fields, fields...)
	return &CarSelect{CarQuery: cq}
}

func (cq *CarQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !car.ValidKey(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.dynamodb = prev
	}
	return nil
}

func (cq *CarQuery) dynamodbAll(ctx context.Context) ([]*Car, error) {
	var (
		nodes       = []*Car{}
		_node       *Car
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withOwner != nil,
		}
	)
	if cq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Keys = append(_spec.Node.Keys, car.ForeignKeys...)
	}

	_spec.Item = _node.item
	_spec.Assign = func(items []map[string]types.AttributeValue) error {
		for _, item := range items {
			node := &Car{}
			if err := node.FromItem(item); err != nil {
				return err
			}
			node.Edges.loadedTypes = loadedTypes
			node.config = cq.config
			nodes = append(nodes, node)
		}

		return nil
	}
	if err := dynamodbgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cq.withOwner; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Car)
		for i := range nodes {
			if nodes[i].user_cars == nil {
				continue
			}
			fk := *nodes[i].user_cars
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_cars" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	return nodes, nil
}

func (cq *CarQuery) dynamodbCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	return dynamodbgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CarQuery) dynamodbExist(ctx context.Context) (bool, error) {
	n, err := cq.dynamodbCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cq *CarQuery) querySpec() *dynamodbgraph.QuerySpec {
	_spec := &dynamodbgraph.QuerySpec{
		Node: &dynamodbgraph.NodeSpec{
			Table: car.Table,
			Keys:  car.Keys,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeUUID,
				Key:  car.FieldID,
			},
		},
		From: cq.dynamodb,
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *dynamodb.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	return _spec
}

func (cq *CarQuery) dynamodbQuery(ctx context.Context) *dynamodb.Selector {
	c1 := car.Table
	selector := dynamodb.Select(car.Keys...).From(c1)
	if cq.dynamodb != nil {
		selector = cq.dynamodb
		selector.Select(car.Keys...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	return selector
}

// CarSelect is the builder for selecting fields of Car entities.
type CarSelect struct {
	*CarQuery
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CarSelect) Scan(ctx context.Context, v any) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.dynamodb = cs.CarQuery.dynamodbQuery(ctx)
	return cs.dynamodbScan(ctx, v)
}

func (cs *CarSelect) dynamodbScan(ctx context.Context, v interface{}) error {
	selector := cs.dynamodbQuery()
	op, args := selector.BuildExpressions().Op()
	var scanOutput sdk.ScanOutput
	if err := cs.driver.Exec(ctx, op, args, &scanOutput); err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	key := cs.fields[0]
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

func (cs *CarSelect) dynamodbQuery() *dynamodb.Selector {
	selector := cs.dynamodb
	selector.Select(cs.fields...)
	return selector
}
