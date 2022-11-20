// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package group

import (
	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldID, id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.In(FieldID, v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.NotIn(FieldID, v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldID, id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldID, id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldID, id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldID, id))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(UsersTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, false, false, UsersTable, UsersAttributes...),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(UsersInverseTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, false, false, UsersTable, UsersAttributes...),
		)
		dynamodbgraph.HasNeighborsWith(s, step, func(s *dynamodb.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		var ps []*dynamodb.Predicate
		for _, p := range predicates {
			selector := dynamodb.Select()
			p(selector)
			ps = append(ps, selector.P())
		}
		s.Where(dynamodb.And(ps...))
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		var ps []*dynamodb.Predicate
		for _, p := range predicates {
			selector := dynamodb.Select()
			p(selector)
			ps = append(ps, selector.P())
		}
		s.Where(dynamodb.Or(ps...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(func(s *dynamodb.Selector) {
		p(s.Not())
	})
}
