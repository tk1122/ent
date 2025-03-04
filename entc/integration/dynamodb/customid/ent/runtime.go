// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/blob"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/car"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/pet"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/schema"
	uuid "github.com/satori/go.uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	blobFields := schema.Blob{}.Fields()
	_ = blobFields
	// blobDescUUID is the schema descriptor for uuid field.
	blobDescUUID := blobFields[1].Descriptor()
	// blob.DefaultUUID holds the default value on creation for the uuid field.
	blob.DefaultUUID = blobDescUUID.Default.(func() uuid.UUID)
	// blobDescID is the schema descriptor for id field.
	blobDescID := blobFields[0].Descriptor()
	// blob.DefaultID holds the default value on creation for the id field.
	blob.DefaultID = blobDescID.Default.(func() uuid.UUID)
	carMixin := schema.Car{}.Mixin()
	carMixinFields0 := carMixin[0].Fields()
	_ = carMixinFields0
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescBeforeID is the schema descriptor for before_id field.
	carDescBeforeID := carMixinFields0[0].Descriptor()
	// car.BeforeIDValidator is a validator for the "before_id" field. It is called by the builders before save.
	car.BeforeIDValidator = carDescBeforeID.Validators[0].(func(float64) error)
	// carDescAfterID is the schema descriptor for after_id field.
	carDescAfterID := carMixinFields0[2].Descriptor()
	// car.AfterIDValidator is a validator for the "after_id" field. It is called by the builders before save.
	car.AfterIDValidator = carDescAfterID.Validators[0].(func(float64) error)
	// carDescID is the schema descriptor for id field.
	carDescID := carMixinFields0[1].Descriptor()
	// car.IDValidator is a validator for the "id" field. It is called by the builders before save.
	car.IDValidator = carDescID.Validators[0].(func(int) error)
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescID is the schema descriptor for id field.
	petDescID := petFields[0].Descriptor()
	// pet.IDValidator is a validator for the "id" field. It is called by the builders before save.
	pet.IDValidator = func() func(string) error {
		validators := petDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
