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
	"entgo.io/ent/examples/dynamodb/o2m2types/ent/pet"
	"entgo.io/ent/examples/dynamodb/o2m2types/ent/predicate"
	"entgo.io/ent/examples/dynamodb/o2m2types/ent/user"
	"entgo.io/ent/schema/field"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	hooks    []Hook
	mutation *PetMutation
}

// Where appends a list predicates to the PetUpdate builder.
func (pu *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PetUpdate) SetName(s string) *PetUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pu *PetUpdate) SetOwnerID(id int) *PetUpdate {
	pu.mutation.SetOwnerID(id)
	return pu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerID(id *int) *PetUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the "owner" edge to the User entity.
func (pu *PetUpdate) SetOwner(u *User) *PetUpdate {
	return pu.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (pu *PetUpdate) Mutation() *PetMutation {
	return pu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (pu *PetUpdate) ClearOwner() *PetUpdate {
	pu.mutation.ClearOwner()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.dynamodbSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.dynamodbSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PetUpdate) dynamodbSave(ctx context.Context) (n int, err error) {
	_spec := &dynamodbgraph.UpdateSpec{
		Node: &dynamodbgraph.NodeSpec{
			Table: pet.Table,
			Keys:  pet.Keys,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeInt,
				Key:  pet.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *dynamodb.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &dynamodbgraph.FieldSpec{
			Type:  field.TypeString,
			Value: value,
			Key:   pet.FieldName,
		})
	}
	if pu.mutation.OwnerCleared() {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.M2O,
			Inverse:    true,
			Table:      pet.OwnerTable,
			Attributes: []string{pet.OwnerAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeInt,
					Key:  user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.M2O,
			Inverse:    true,
			Table:      pet.OwnerTable,
			Attributes: []string{pet.OwnerAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeInt,
					Key:  user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = dynamodbgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		return 0, err
	}
	return n, nil
}

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PetMutation
}

// SetName sets the "name" field.
func (puo *PetUpdateOne) SetName(s string) *PetUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (puo *PetUpdateOne) SetOwnerID(id int) *PetUpdateOne {
	puo.mutation.SetOwnerID(id)
	return puo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerID(id *int) *PetUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the "owner" edge to the User entity.
func (puo *PetUpdateOne) SetOwner(u *User) *PetUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (puo *PetUpdateOne) Mutation() *PetMutation {
	return puo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (puo *PetUpdateOne) ClearOwner() *PetUpdateOne {
	puo.mutation.ClearOwner()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PetUpdateOne) Select(field string, fields ...string) *PetUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pet entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	var (
		err  error
		node *Pet
	)
	if len(puo.hooks) == 0 {
		node, err = puo.dynamodbSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.dynamodbSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Pet)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PetUpdateOne) dynamodbSave(ctx context.Context) (_node *Pet, err error) {
	_spec := &dynamodbgraph.UpdateSpec{
		Node: &dynamodbgraph.NodeSpec{
			Table: pet.Table,
			Keys:  pet.Keys,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeInt,
				Key:  pet.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Pet.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Keys = make([]string, 0, len(fields))
		_spec.Node.Keys = append(_spec.Node.Keys, pet.FieldID)
		for _, f := range fields {
			if !pet.ValidKey(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pet.FieldID {
				_spec.Node.Keys = append(_spec.Node.Keys, f)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &dynamodbgraph.FieldSpec{
			Type:  field.TypeString,
			Value: value,
			Key:   pet.FieldName,
		})
	}
	if puo.mutation.OwnerCleared() {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.M2O,
			Inverse:    true,
			Table:      pet.OwnerTable,
			Attributes: []string{pet.OwnerAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeInt,
					Key:  user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.M2O,
			Inverse:    true,
			Table:      pet.OwnerTable,
			Attributes: []string{pet.OwnerAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeInt,
					Key:  user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Item = _node.item
	_spec.Assign = func(item interface{}) error {
		_node = &Pet{}
		if err := _node.FromItem(item); err != nil {
			return err
		}
		_node.config = puo.config

		return nil
	}
	if err = dynamodbgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		return nil, err
	}
	return _node, nil
}
