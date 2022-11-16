// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeSpouse holds the string denoting the spouse edge name in mutations.
	EdgeSpouse = "spouse"
	// Table holds the table name of the user in the database.
	Table = "users"
	// SpouseTable is the table that holds the spouse relation/edge.
	SpouseTable = "users"
	// SpouseAttribute is the table column denoting the spouse relation/edge.
	SpouseAttribute = "user_spouse"
)

// Keys holds all DynamoDB keys for user fields.
var Keys = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAge,
	FieldName,
	"user_spouse",
}

// ForeignKeys holds the Mongo foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"user_spouse",
}

// ValidKey reports if the key is valid (one of keys of collection's fields).
func ValidKey(key string) bool {
	for i := range Keys {
		if key == Keys[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if key == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
