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
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/pet"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/predicate"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/user"
	"entgo.io/ent/schema/field"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// PetQuery is the builder for querying Pet entities.
type PetQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	fields         []string
	predicates     []predicate.Pet
	withOwner      *UserQuery
	withCars       *CarQuery
	withFriends    *PetQuery
	withBestFriend *PetQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
	path     func(context.Context) (*dynamodb.Selector, error)
}

// Where adds a new predicate for the PetQuery builder.
func (pq *PetQuery) Where(ps ...predicate.Pet) *PetQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PetQuery) Limit(limit int) *PetQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PetQuery) Offset(offset int) *PetQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PetQuery) Unique(unique bool) *PetQuery {
	pq.unique = &unique
	return pq
}

// First returns the first Pet entity from the query.
// Returns a *NotFoundError when no Pet was found.
func (pq *PetQuery) First(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PetQuery) FirstX(ctx context.Context) *Pet {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pet ID from the query.
// Returns a *NotFoundError when no Pet ID was found.
func (pq *PetQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PetQuery) FirstIDX(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pet entity is found.
// Returns a *NotFoundError when no Pet entities are found.
func (pq *PetQuery) Only(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pet.Label}
	default:
		return nil, &NotSingularError{pet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PetQuery) OnlyX(ctx context.Context) *Pet {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pet ID in the query.
// Returns a *NotSingularError when more than one Pet ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PetQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = &NotSingularError{pet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PetQuery) OnlyIDX(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pets.
func (pq *PetQuery) All(ctx context.Context) ([]*Pet, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.dynamodbAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PetQuery) AllX(ctx context.Context) []*Pet {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pet IDs.
func (pq *PetQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := pq.Select(pet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PetQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PetQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.dynamodbCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PetQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PetQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.dynamodbExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PetQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PetQuery) Clone() *PetQuery {
	if pq == nil {
		return nil
	}
	return &PetQuery{
		config:         pq.config,
		limit:          pq.limit,
		offset:         pq.offset,
		predicates:     append([]predicate.Pet{}, pq.predicates...),
		withOwner:      pq.withOwner.Clone(),
		withCars:       pq.withCars.Clone(),
		withFriends:    pq.withFriends.Clone(),
		withBestFriend: pq.withBestFriend.Clone(),
		// clone intermediate query.
		dynamodb: pq.dynamodb.Clone(),
		path:     pq.path,
		unique:   pq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithOwner(opts ...func(*UserQuery)) *PetQuery {
	query := &UserQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withOwner = query
	return pq
}

// WithCars tells the query-builder to eager-load the nodes that are connected to
// the "cars" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithCars(opts ...func(*CarQuery)) *PetQuery {
	query := &CarQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withCars = query
	return pq
}

// WithFriends tells the query-builder to eager-load the nodes that are connected to
// the "friends" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithFriends(opts ...func(*PetQuery)) *PetQuery {
	query := &PetQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withFriends = query
	return pq
}

// WithBestFriend tells the query-builder to eager-load the nodes that are connected to
// the "best_friend" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithBestFriend(opts ...func(*PetQuery)) *PetQuery {
	query := &PetQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withBestFriend = query
	return pq
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (pq *PetQuery) Select(fields ...string) *PetSelect {
	pq.fields = append(pq.fields, fields...)
	return &PetSelect{PetQuery: pq}
}

func (pq *PetQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !pet.ValidKey(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.dynamodb = prev
	}
	return nil
}

func (pq *PetQuery) dynamodbAll(ctx context.Context) ([]*Pet, error) {
	var (
		nodes       = []*Pet{}
		_node       *Pet
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [4]bool{
			pq.withOwner != nil,
			pq.withCars != nil,
			pq.withFriends != nil,
			pq.withBestFriend != nil,
		}
	)
	if pq.withOwner != nil || pq.withBestFriend != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Keys = append(_spec.Node.Keys, pet.ForeignKeys...)
	}

	_spec.Item = _node.item
	_spec.Assign = func(items []map[string]types.AttributeValue) error {
		for _, item := range items {
			node := &Pet{}
			if err := node.FromItem(item); err != nil {
				return err
			}
			node.Edges.loadedTypes = loadedTypes
			node.config = pq.config
			nodes = append(nodes, node)
		}

		return nil
	}
	if err := dynamodbgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Pet)
		for i := range nodes {
			if nodes[i].user_pets == nil {
				continue
			}
			fk := *nodes[i].user_pets
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_pets" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	if query := pq.withCars; query != nil {
		fks := make([]interface{}, 0, len(nodes))
		nodeids := make(map[string]*Pet)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Cars = []*Car{}
		}
		query.withFKs = true
		query.Where(predicate.Car(func(s *dynamodb.Selector) {
			s.Where(dynamodb.In(pet.CarsAttribute, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.pet_cars
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "pet_cars" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pet_cars" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Cars = append(node.Edges.Cars, n)
		}
	}

	if query := pq.withFriends; query != nil {
		var (
			edgeids []string
			edges   = make(map[string][]*Pet)
		)

		for _, node := range nodes {
			node.Edges.Friends = []*Pet{}
			edgeids = append(edgeids, node.friend_id...)
			for _, id := range node.friend_id {
				edges[id] = append(edges[id], node)
			}
		}

		query.Where(pet.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "friends" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Friends = append(nodes[i].Edges.Friends, n)
			}
		}
	}

	if query := pq.withBestFriend; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Pet)
		for i := range nodes {
			if nodes[i].pet_best_friend == nil {
				continue
			}
			fk := *nodes[i].pet_best_friend
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(pet.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pet_best_friend" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BestFriend = n
			}
		}
	}

	return nodes, nil
}

func (pq *PetQuery) dynamodbCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return dynamodbgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PetQuery) dynamodbExist(ctx context.Context) (bool, error) {
	n, err := pq.dynamodbCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PetQuery) querySpec() *dynamodbgraph.QuerySpec {
	_spec := &dynamodbgraph.QuerySpec{
		Node: &dynamodbgraph.NodeSpec{
			Table: pet.Table,
			Keys:  pet.Keys,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeString,
				Key:  pet.FieldID,
			},
		},
		From: pq.dynamodb,
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *dynamodb.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	return _spec
}

func (pq *PetQuery) dynamodbQuery(ctx context.Context) *dynamodb.Selector {
	c1 := pet.Table
	selector := dynamodb.Select(pet.Keys...).From(c1)
	if pq.dynamodb != nil {
		selector = pq.dynamodb
		selector.Select(pet.Keys...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	return selector
}

// PetSelect is the builder for selecting fields of Pet entities.
type PetSelect struct {
	*PetQuery
	// intermediate query (i.e. traversal path).
	dynamodb *dynamodb.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PetSelect) Scan(ctx context.Context, v any) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.dynamodb = ps.PetQuery.dynamodbQuery(ctx)
	return ps.dynamodbScan(ctx, v)
}

func (ps *PetSelect) dynamodbScan(ctx context.Context, v interface{}) error {
	selector := ps.dynamodbQuery()
	op, args := selector.BuildExpressions().Op()
	var scanOutput sdk.ScanOutput
	if err := ps.driver.Exec(ctx, op, args, &scanOutput); err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	key := ps.fields[0]
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

func (ps *PetSelect) dynamodbQuery() *dynamodb.Selector {
	selector := ps.dynamodb
	selector.Select(ps.fields...)
	return selector
}
