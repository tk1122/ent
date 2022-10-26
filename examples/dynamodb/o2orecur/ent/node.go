// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

// Node is the model entity for the Node schema.
type Node struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value int `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NodeQuery when eager-loading is set.
	Edges     NodeEdges `json:"edges"`
	node_next *int
}

// NodeEdges holds the relations/edges for other nodes in the graph.
type NodeEdges struct {
	// Prev holds the value of the prev edge.
	Prev *Node `json:"prev,omitempty"`
	// Next holds the value of the next edge.
	Next *Node `json:"next,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}
