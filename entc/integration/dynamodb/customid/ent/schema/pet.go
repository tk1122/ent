// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(25).
			NotEmpty().
			Unique().
			Immutable(),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
		edge.To("cars", Car.Type),
		edge.To("friends", Pet.Type),
		edge.To("best_friend", Pet.Type).
			Unique(),
	}
}

// Indexes of the Pet.
func (Pet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
