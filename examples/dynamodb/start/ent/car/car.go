// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package car

import (
	uuid "github.com/satori/go.uuid"
)

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldRegisteredAt holds the string denoting the registered_at field in the database.
	FieldRegisteredAt = "registered_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the car in the database.
	Table = "cars"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "cars"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerAttribute is the table column denoting the owner relation/edge.
	OwnerAttribute = "user_cars"
)

// Keys holds all DynamoDB keys for car fields.
var Keys = []string{
	FieldID,
	FieldModel,
	FieldRegisteredAt,
	"user_cars",
}

// ForeignKeys holds the Mongo foreign-keys that are owned by the Car type.
var ForeignKeys = []string{
	"user_cars",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
