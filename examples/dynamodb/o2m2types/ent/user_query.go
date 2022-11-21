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
	"entgo.io/ent/examples/dynamodb/o2m2types/ent/predicate"
	"entgo.io/ent/examples/dynamodb/o2m2types/ent/user"
	"entgo.io/ent/schema/field"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	fields     []string
	predicates []predicate.User
	withPets   *PetQuery
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
	path     func(context.Context) (*dynamodb.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.unique = &unique
	return uq
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uq.dynamodbAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uq.dynamodbCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uq.dynamodbExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:     uq.config,
		limit:      uq.limit,
		offset:     uq.offset,
		predicates: append([]predicate.User{}, uq.predicates...),
		withPets:   uq.withPets.Clone(),
		// clone intermediate query.
		dynamodb: uq.dynamodb.Clone(),
		path:     uq.path,
		unique:   uq.unique,
	}
}

// WithPets tells the query-builder to eager-load the nodes that are connected to
// the "pets" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithPets(opts ...func(*PetQuery)) *UserQuery {
	query := &PetQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withPets = query
	return uq
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Age int `json:"age,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldAge).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.fields = append(uq.fields, fields...)
	return &UserSelect{UserQuery: uq}
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uq.fields {
		if !user.ValidKey(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.dynamodb = prev
	}
	return nil
}

func (uq *UserQuery) dynamodbAll(ctx context.Context) ([]*User, error) {
	var (
		nodes       = []*User{}
		_node       *User
		_spec       = uq.querySpec()
		loadedTypes = [1]bool{
			uq.withPets != nil,
		}
	)

	_spec.Item = _node.item
	_spec.Assign = func(items []map[string]types.AttributeValue) error {
		for _, item := range items {
			node := &User{}
			if err := node.FromItem(item); err != nil {
				return err
			}
			node.Edges.loadedTypes = loadedTypes
			node.config = uq.config
			nodes = append(nodes, node)
		}

		return nil
	}
	if err := dynamodbgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uq.withPets; query != nil {
		fks := make([]interface{}, 0, len(nodes))
		nodeids := make(map[int]*User)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Pets = []*Pet{}
		}
		query.withFKs = true
		query.Where(predicate.Pet(func(s *dynamodb.Selector) {
			s.Where(dynamodb.In(user.PetsAttribute, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.user_pets
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "user_pets" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_pets" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Pets = append(node.Edges.Pets, n)
		}
	}

	return nodes, nil
}

func (uq *UserQuery) dynamodbCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	return dynamodbgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) dynamodbExist(ctx context.Context) (bool, error) {
	n, err := uq.dynamodbCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (uq *UserQuery) querySpec() *dynamodbgraph.QuerySpec {
	_spec := &dynamodbgraph.QuerySpec{
		Node: &dynamodbgraph.NodeSpec{
			Table: user.Table,
			Keys:  user.Keys,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeInt,
				Key:  user.FieldID,
			},
		},
		From: uq.dynamodb,
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *dynamodb.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	return _spec
}

func (uq *UserQuery) dynamodbQuery(ctx context.Context) *dynamodb.Selector {
	return nil
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	us.dynamodb = us.UserQuery.dynamodbQuery(ctx)
	return us.dynamodbScan(ctx, v)
}

func (us *UserSelect) dynamodbScan(ctx context.Context, v interface{}) error {
	return nil
}

func (us *UserSelect) dynamodbQuery() *dynamodb.Selector {
	return nil
}
