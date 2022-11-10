// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetQuery when eager-loading is set.
	Edges     PetEdges `json:"edges"`
	user_pets *int
}

// PetEdges holds the relations/edges for other nodes in the graph.
type PetEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PetItem represents item schema in MongoDB.
type PetItem struct {
	ID   int    `dynamodbav:"id"`
	Name string `dynamodbav:"name"`

	UserPets *int `dynamodbav:"user_pets"`
}

// item returns the object for receiving item from dynamodb.
func (*Pet) item() interface{} {
	return &PetItem{}
}

// FromItem scans the dynamodb response item into Pet.
func (pe *Pet) FromItem(item interface{}) error {
	var petItem PetItem
	err := attributevalue.UnmarshalMap(item.(map[string]types.AttributeValue), &petItem)
	if err != nil {
		return err
	}
	pe.ID = petItem.ID
	pe.Name = petItem.Name

	pe.user_pets = petItem.UserPets

	return nil
}

// QueryOwner queries the "owner" edge of the Pet entity.
func (pe *Pet) QueryOwner() *UserQuery {
	return (&PetClient{config: pe.config}).QueryOwner(pe)
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
