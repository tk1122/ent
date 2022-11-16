// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package car

import (
	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldID, id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.In(FieldID, v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.NotIn(FieldID, v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldID, id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldID, id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldID, id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldID, id))
	})
}

// BeforeID applies equality check predicate on the "before_id" field. It's identical to BeforeIDEQ.
func BeforeID(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldBeforeID, v))
	})
}

// AfterID applies equality check predicate on the "after_id" field. It's identical to AfterIDEQ.
func AfterID(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldAfterID, v))
	})
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldModel, v))
	})
}

// BeforeIDEQ applies the EQ predicate on the "before_id" field.
func BeforeIDEQ(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldBeforeID, v))
	})
}

// BeforeIDNEQ applies the NEQ predicate on the "before_id" field.
func BeforeIDNEQ(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldBeforeID, v))
	})
}

// BeforeIDIn applies the In predicate on the "before_id" field.
func BeforeIDIn(vs ...float64) predicate.Car {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.In(FieldBeforeID, v...))
	})
}

// BeforeIDNotIn applies the NotIn predicate on the "before_id" field.
func BeforeIDNotIn(vs ...float64) predicate.Car {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NotIn(FieldBeforeID, v...))
	})
}

// BeforeIDGT applies the GT predicate on the "before_id" field.
func BeforeIDGT(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldBeforeID, v))
	})
}

// BeforeIDGTE applies the GTE predicate on the "before_id" field.
func BeforeIDGTE(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldBeforeID, v))
	})
}

// BeforeIDLT applies the LT predicate on the "before_id" field.
func BeforeIDLT(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldBeforeID, v))
	})
}

// BeforeIDLTE applies the LTE predicate on the "before_id" field.
func BeforeIDLTE(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldBeforeID, v))
	})
}

// AfterIDEQ applies the EQ predicate on the "after_id" field.
func AfterIDEQ(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldAfterID, v))
	})
}

// AfterIDNEQ applies the NEQ predicate on the "after_id" field.
func AfterIDNEQ(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldAfterID, v))
	})
}

// AfterIDIn applies the In predicate on the "after_id" field.
func AfterIDIn(vs ...float64) predicate.Car {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.In(FieldAfterID, v...))
	})
}

// AfterIDNotIn applies the NotIn predicate on the "after_id" field.
func AfterIDNotIn(vs ...float64) predicate.Car {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NotIn(FieldAfterID, v...))
	})
}

// AfterIDGT applies the GT predicate on the "after_id" field.
func AfterIDGT(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldAfterID, v))
	})
}

// AfterIDGTE applies the GTE predicate on the "after_id" field.
func AfterIDGTE(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldAfterID, v))
	})
}

// AfterIDLT applies the LT predicate on the "after_id" field.
func AfterIDLT(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldAfterID, v))
	})
}

// AfterIDLTE applies the LTE predicate on the "after_id" field.
func AfterIDLTE(v float64) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldAfterID, v))
	})
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldModel, v))
	})
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldModel, v))
	})
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Car {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.In(FieldModel, v...))
	})
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Car {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NotIn(FieldModel, v...))
	})
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldModel, v))
	})
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldModel, v))
	})
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldModel, v))
	})
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldModel, v))
	})
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.Contains(FieldModel, v))
	})
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s.Where(dynamodb.HasPrefix(FieldModel, v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(OwnerTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2O, true, false, OwnerTable, OwnerAttribute),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Pet) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(OwnerInverseTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2O, true, false, OwnerTable, OwnerAttribute),
		)
		dynamodbgraph.HasNeighborsWith(s, step, func(s *dynamodb.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s1 := s.Clone()
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		s1 := s.Clone()
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Car) predicate.Car {
	return predicate.Car(func(s *dynamodb.Selector) {
		p(s.Not())
	})
}
