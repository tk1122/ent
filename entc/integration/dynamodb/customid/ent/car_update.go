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
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/car"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/pet"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/predicate"
	"entgo.io/ent/schema/field"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarUpdate builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetBeforeID sets the "before_id" field.
func (cu *CarUpdate) SetBeforeID(f float64) *CarUpdate {
	cu.mutation.ResetBeforeID()
	cu.mutation.SetBeforeID(f)
	return cu
}

// SetNillableBeforeID sets the "before_id" field if the given value is not nil.
func (cu *CarUpdate) SetNillableBeforeID(f *float64) *CarUpdate {
	if f != nil {
		cu.SetBeforeID(*f)
	}
	return cu
}

// AddBeforeID adds f to the "before_id" field.
func (cu *CarUpdate) AddBeforeID(f float64) *CarUpdate {
	cu.mutation.AddBeforeID(f)
	return cu
}

// ClearBeforeID clears the value of the "before_id" field.
func (cu *CarUpdate) ClearBeforeID() *CarUpdate {
	cu.mutation.ClearBeforeID()
	return cu
}

// SetAfterID sets the "after_id" field.
func (cu *CarUpdate) SetAfterID(f float64) *CarUpdate {
	cu.mutation.ResetAfterID()
	cu.mutation.SetAfterID(f)
	return cu
}

// SetNillableAfterID sets the "after_id" field if the given value is not nil.
func (cu *CarUpdate) SetNillableAfterID(f *float64) *CarUpdate {
	if f != nil {
		cu.SetAfterID(*f)
	}
	return cu
}

// AddAfterID adds f to the "after_id" field.
func (cu *CarUpdate) AddAfterID(f float64) *CarUpdate {
	cu.mutation.AddAfterID(f)
	return cu
}

// ClearAfterID clears the value of the "after_id" field.
func (cu *CarUpdate) ClearAfterID() *CarUpdate {
	cu.mutation.ClearAfterID()
	return cu
}

// SetModel sets the "model" field.
func (cu *CarUpdate) SetModel(s string) *CarUpdate {
	cu.mutation.SetModel(s)
	return cu
}

// SetOwnerID sets the "owner" edge to the Pet entity by ID.
func (cu *CarUpdate) SetOwnerID(id string) *CarUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the "owner" edge to the Pet entity by ID if the given value is not nil.
func (cu *CarUpdate) SetNillableOwnerID(id *string) *CarUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the "owner" edge to the Pet entity.
func (cu *CarUpdate) SetOwner(p *Pet) *CarUpdate {
	return cu.SetOwnerID(p.ID)
}

// Mutation returns the CarMutation object of the builder.
func (cu *CarUpdate) Mutation() *CarMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the Pet entity.
func (cu *CarUpdate) ClearOwner() *CarUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.dynamodbSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.dynamodbSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CarUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CarUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CarUpdate) check() error {
	if v, ok := cu.mutation.BeforeID(); ok {
		if err := car.BeforeIDValidator(v); err != nil {
			return &ValidationError{Name: "before_id", err: fmt.Errorf(`ent: validator failed for field "Car.before_id": %w`, err)}
		}
	}
	if v, ok := cu.mutation.AfterID(); ok {
		if err := car.AfterIDValidator(v); err != nil {
			return &ValidationError{Name: "after_id", err: fmt.Errorf(`ent: validator failed for field "Car.after_id": %w`, err)}
		}
	}
	return nil
}

func (cu *CarUpdate) dynamodbSave(ctx context.Context) (n int, err error) {
	_spec := &dynamodbgraph.UpdateSpec{
		Node: &dynamodbgraph.NodeSpec{
			Table: car.Table,
			Keys:  car.Keys,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeInt,
				Key:  car.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *dynamodb.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.BeforeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &dynamodbgraph.FieldSpec{
			Type:  field.TypeFloat64,
			Value: value,
			Key:   car.FieldBeforeID,
		})
	}
	if value, ok := cu.mutation.AddedBeforeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &dynamodbgraph.FieldSpec{
			Type:  field.TypeFloat64,
			Value: value,
			Key:   car.FieldBeforeID,
		})
	}
	if cu.mutation.BeforeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &dynamodbgraph.FieldSpec{
			Type: field.TypeFloat64,
			Key:  car.FieldBeforeID,
		})
	}
	if value, ok := cu.mutation.AfterID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &dynamodbgraph.FieldSpec{
			Type:  field.TypeFloat64,
			Value: value,
			Key:   car.FieldAfterID,
		})
	}
	if value, ok := cu.mutation.AddedAfterID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &dynamodbgraph.FieldSpec{
			Type:  field.TypeFloat64,
			Value: value,
			Key:   car.FieldAfterID,
		})
	}
	if cu.mutation.AfterIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &dynamodbgraph.FieldSpec{
			Type: field.TypeFloat64,
			Key:  car.FieldAfterID,
		})
	}
	if value, ok := cu.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &dynamodbgraph.FieldSpec{
			Type:  field.TypeString,
			Value: value,
			Key:   car.FieldModel,
		})
	}
	if cu.mutation.OwnerCleared() {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.M2O,
			Inverse:    true,
			Table:      car.OwnerTable,
			Attributes: []string{car.OwnerAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeString,
					Key:  pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.M2O,
			Inverse:    true,
			Table:      car.OwnerTable,
			Attributes: []string{car.OwnerAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeString,
					Key:  pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = dynamodbgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		return 0, err
	}
	return n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CarMutation
}

// SetBeforeID sets the "before_id" field.
func (cuo *CarUpdateOne) SetBeforeID(f float64) *CarUpdateOne {
	cuo.mutation.ResetBeforeID()
	cuo.mutation.SetBeforeID(f)
	return cuo
}

// SetNillableBeforeID sets the "before_id" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableBeforeID(f *float64) *CarUpdateOne {
	if f != nil {
		cuo.SetBeforeID(*f)
	}
	return cuo
}

// AddBeforeID adds f to the "before_id" field.
func (cuo *CarUpdateOne) AddBeforeID(f float64) *CarUpdateOne {
	cuo.mutation.AddBeforeID(f)
	return cuo
}

// ClearBeforeID clears the value of the "before_id" field.
func (cuo *CarUpdateOne) ClearBeforeID() *CarUpdateOne {
	cuo.mutation.ClearBeforeID()
	return cuo
}

// SetAfterID sets the "after_id" field.
func (cuo *CarUpdateOne) SetAfterID(f float64) *CarUpdateOne {
	cuo.mutation.ResetAfterID()
	cuo.mutation.SetAfterID(f)
	return cuo
}

// SetNillableAfterID sets the "after_id" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableAfterID(f *float64) *CarUpdateOne {
	if f != nil {
		cuo.SetAfterID(*f)
	}
	return cuo
}

// AddAfterID adds f to the "after_id" field.
func (cuo *CarUpdateOne) AddAfterID(f float64) *CarUpdateOne {
	cuo.mutation.AddAfterID(f)
	return cuo
}

// ClearAfterID clears the value of the "after_id" field.
func (cuo *CarUpdateOne) ClearAfterID() *CarUpdateOne {
	cuo.mutation.ClearAfterID()
	return cuo
}

// SetModel sets the "model" field.
func (cuo *CarUpdateOne) SetModel(s string) *CarUpdateOne {
	cuo.mutation.SetModel(s)
	return cuo
}

// SetOwnerID sets the "owner" edge to the Pet entity by ID.
func (cuo *CarUpdateOne) SetOwnerID(id string) *CarUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the "owner" edge to the Pet entity by ID if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableOwnerID(id *string) *CarUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the "owner" edge to the Pet entity.
func (cuo *CarUpdateOne) SetOwner(p *Pet) *CarUpdateOne {
	return cuo.SetOwnerID(p.ID)
}

// Mutation returns the CarMutation object of the builder.
func (cuo *CarUpdateOne) Mutation() *CarMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the Pet entity.
func (cuo *CarUpdateOne) ClearOwner() *CarUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CarUpdateOne) Select(field string, fields ...string) *CarUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Car entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	var (
		err  error
		node *Car
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.dynamodbSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.dynamodbSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Car)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CarMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CarUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CarUpdateOne) check() error {
	if v, ok := cuo.mutation.BeforeID(); ok {
		if err := car.BeforeIDValidator(v); err != nil {
			return &ValidationError{Name: "before_id", err: fmt.Errorf(`ent: validator failed for field "Car.before_id": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.AfterID(); ok {
		if err := car.AfterIDValidator(v); err != nil {
			return &ValidationError{Name: "after_id", err: fmt.Errorf(`ent: validator failed for field "Car.after_id": %w`, err)}
		}
	}
	return nil
}

func (cuo *CarUpdateOne) dynamodbSave(ctx context.Context) (_node *Car, err error) {
	_spec := &dynamodbgraph.UpdateSpec{
		Node: &dynamodbgraph.NodeSpec{
			Table: car.Table,
			Keys:  car.Keys,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeInt,
				Key:  car.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Car.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Keys = make([]string, 0, len(fields))
		_spec.Node.Keys = append(_spec.Node.Keys, car.FieldID)
		for _, f := range fields {
			if !car.ValidKey(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != car.FieldID {
				_spec.Node.Keys = append(_spec.Node.Keys, f)
			}
		}
	}
	if value, ok := cuo.mutation.BeforeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &dynamodbgraph.FieldSpec{
			Type:  field.TypeFloat64,
			Value: value,
			Key:   car.FieldBeforeID,
		})
	}
	if value, ok := cuo.mutation.AddedBeforeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &dynamodbgraph.FieldSpec{
			Type:  field.TypeFloat64,
			Value: value,
			Key:   car.FieldBeforeID,
		})
	}
	if cuo.mutation.BeforeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &dynamodbgraph.FieldSpec{
			Type: field.TypeFloat64,
			Key:  car.FieldBeforeID,
		})
	}
	if value, ok := cuo.mutation.AfterID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &dynamodbgraph.FieldSpec{
			Type:  field.TypeFloat64,
			Value: value,
			Key:   car.FieldAfterID,
		})
	}
	if value, ok := cuo.mutation.AddedAfterID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &dynamodbgraph.FieldSpec{
			Type:  field.TypeFloat64,
			Value: value,
			Key:   car.FieldAfterID,
		})
	}
	if cuo.mutation.AfterIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &dynamodbgraph.FieldSpec{
			Type: field.TypeFloat64,
			Key:  car.FieldAfterID,
		})
	}
	if value, ok := cuo.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &dynamodbgraph.FieldSpec{
			Type:  field.TypeString,
			Value: value,
			Key:   car.FieldModel,
		})
	}
	if cuo.mutation.OwnerCleared() {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.M2O,
			Inverse:    true,
			Table:      car.OwnerTable,
			Attributes: []string{car.OwnerAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeString,
					Key:  pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.M2O,
			Inverse:    true,
			Table:      car.OwnerTable,
			Attributes: []string{car.OwnerAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeString,
					Key:  pet.FieldID,
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
		_node = &Car{}
		if err := _node.FromItem(item); err != nil {
			return err
		}
		_node.config = cuo.config

		return nil
	}
	if err = dynamodbgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		return nil, err
	}
	return _node, nil
}
