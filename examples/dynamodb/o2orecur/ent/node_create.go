// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/examples/dynamodb/o2orecur/ent/node"
	"entgo.io/ent/schema/field"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	mutation *NodeMutation
	hooks    []Hook
}

// SetValue sets the "value" field.
func (nc *NodeCreate) SetValue(i int) *NodeCreate {
	nc.mutation.SetValue(i)
	return nc
}

// SetID sets the "id" field.
func (nc *NodeCreate) SetID(i int) *NodeCreate {
	nc.mutation.SetID(i)
	return nc
}

// SetPrevID sets the "prev" edge to the Node entity by ID.
func (nc *NodeCreate) SetPrevID(id int) *NodeCreate {
	nc.mutation.SetPrevID(id)
	return nc
}

// SetNillablePrevID sets the "prev" edge to the Node entity by ID if the given value is not nil.
func (nc *NodeCreate) SetNillablePrevID(id *int) *NodeCreate {
	if id != nil {
		nc = nc.SetPrevID(*id)
	}
	return nc
}

// SetPrev sets the "prev" edge to the Node entity.
func (nc *NodeCreate) SetPrev(n *Node) *NodeCreate {
	return nc.SetPrevID(n.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (nc *NodeCreate) SetNextID(id int) *NodeCreate {
	nc.mutation.SetNextID(id)
	return nc
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (nc *NodeCreate) SetNillableNextID(id *int) *NodeCreate {
	if id != nil {
		nc = nc.SetNextID(*id)
	}
	return nc
}

// SetNext sets the "next" edge to the Node entity.
func (nc *NodeCreate) SetNext(n *Node) *NodeCreate {
	return nc.SetNextID(n.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (nc *NodeCreate) Mutation() *NodeMutation {
	return nc.mutation
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	var (
		err  error
		node *Node
	)
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.dynamodbSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.dynamodbSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Node)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NodeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NodeCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NodeCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NodeCreate) check() error {
	if _, ok := nc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Node.value"`)}
	}
	return nil
}

func (nc *NodeCreate) dynamodbSave(ctx context.Context) (*Node, error) {
	_node, _spec := nc.createSpec()
	if err := dynamodbgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		return nil, err
	}
	return _node, nil
}

func (nc *NodeCreate) createSpec() (*Node, *dynamodbgraph.CreateSpec) {
	var (
		_node = &Node{config: nc.config}
		_spec = &dynamodbgraph.CreateSpec{
			Table: node.Table,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeInt,
				Key:  node.FieldID,
			},
		}
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &dynamodbgraph.FieldSpec{
			Type:  field.TypeInt,
			Value: value,
			Key:   node.FieldValue,
		})
		_node.Value = value
	}
	if nodes := nc.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.O2O,
			Inverse:    true,
			Table:      node.PrevTable,
			Attributes: []string{node.PrevAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeInt,
					Key:  node.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.node_next = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NextIDs(); len(nodes) > 0 {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.O2O,
			Inverse:    false,
			Table:      node.NextTable,
			Attributes: []string{node.NextAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeInt,
					Key:  node.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NodeCreateBulk is the builder for creating many Node entities in bulk.
type NodeCreateBulk struct {
	config
	builders []*NodeCreate
}

// Save creates the Node entities in the database.
func (ncb *NodeCreateBulk) Save(ctx context.Context) ([]*Node, error) {
	specs := make([]*dynamodbgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Node, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					err = dynamodbgraph.BatchCreate(ctx, ncb.driver, &dynamodbgraph.BatchCreateSpec{Nodes: specs})
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ncb *NodeCreateBulk) SaveX(ctx context.Context) []*Node {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
