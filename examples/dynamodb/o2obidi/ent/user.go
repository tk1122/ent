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

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges       UserEdges `json:"edges"`
	user_spouse *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Spouse holds the value of the spouse edge.
	Spouse *User `json:"spouse,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserItem represents item schema in MongoDB.
type UserItem struct {
	ID   int    `dynamodbav:"id"`
	Age  int    `dynamodbav:"age"`
	Name string `dynamodbav:"name"`

	UserSpouse *int `dynamodbav:"user_spouse"`
}

// item returns the object for receiving item from dynamodb.
func (*User) item() interface{} {
	return &UserItem{}
}

// FromItem scans the dynamodb response item into User.
func (u *User) FromItem(item interface{}) error {
	var userItem UserItem
	err := attributevalue.UnmarshalMap(item.(map[string]types.AttributeValue), &userItem)
	if err != nil {
		return err
	}
	u.ID = userItem.ID
	u.Age = userItem.Age
	u.Name = userItem.Name

	u.user_spouse = userItem.UserSpouse

	return nil
}

// QuerySpouse queries the "spouse" edge of the User entity.
func (u *User) QuerySpouse() *UserQuery {
	return (&UserClient{config: u.config}).QuerySpouse(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
