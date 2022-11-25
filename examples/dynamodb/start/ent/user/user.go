// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	uuid "github.com/satori/go.uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCars holds the string denoting the cars edge name in mutations.
	EdgeCars = "cars"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CarsTable is the table that holds the cars relation/edge.
	CarsTable = "cars"
	// CarsInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarsInverseTable = "cars"
	// CarsAttribute is the table column denoting the cars relation/edge.
	CarsAttribute = "user_cars"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_users"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
)

// Keys holds all DynamoDB keys for user fields.
var Keys = []string{
	FieldID,
	FieldAge,
	FieldName,
	"group_id",
}

var (
	// GroupsAttribute and GroupsAttribute2 are the collection keys denoting the
	// primary key for the groups relation (M2M).
	GroupsAttributes = []string{"group_id", "user_id"}
)

// ValidKey reports if the key is valid (one of keys of collection's fields).
func ValidKey(key string) bool {
	for i := range Keys {
		if key == Keys[i] {
			return true
		}
	}
	return false
}

var (
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
