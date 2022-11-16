// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/satori/go.uuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Blob holds the schema definition for the Blob entity.
type Blob struct {
	ent.Schema
}

// Fields of the Blob.
func (Blob) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.NewV4),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.NewV4).
			Unique(),
	}
}

// Edges of the Blob.
func (Blob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Blob.Type).
			Unique(),
		edge.To("links", Blob.Type),
	}
}

// Indexes of the Blob.
func (Blob) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
