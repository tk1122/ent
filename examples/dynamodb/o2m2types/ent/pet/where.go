// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package pet

import (
	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/examples/dynamodb/o2m2types/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldID, id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.In(FieldID, v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.NotIn(FieldID, v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldID, id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldID, id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldID, id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldID, id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldName, v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldName, v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldName, v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Pet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.In(FieldName, v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Pet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NotIn(FieldName, v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldName, v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldName, v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldName, v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldName, v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.Contains(FieldName, v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s.Where(dynamodb.HasPrefix(FieldName, v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(OwnerTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2O, true, false, OwnerTable, OwnerAttribute),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
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
func And(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		s1 := s.Clone()
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
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
func Not(p predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *dynamodb.Selector) {
		p(s.Not())
	})
}
