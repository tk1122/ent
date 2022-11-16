// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldID, id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.In(FieldID, v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.NotIn(FieldID, v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldID, id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldID, id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldID, id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldID, id))
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(GroupsTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, true, false, GroupsTable, GroupsAttributes...),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(GroupsInverseTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, true, false, GroupsTable, GroupsAttributes...),
		)
		dynamodbgraph.HasNeighborsWith(s, step, func(s *dynamodb.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(ParentTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2O, true, false, ParentTable, ParentAttribute),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(Table, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2O, true, false, ParentTable, ParentAttribute),
		)
		dynamodbgraph.HasNeighborsWith(s, step, func(s *dynamodb.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(ChildrenTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, ChildrenTable, ChildrenAttribute),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(Table, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, ChildrenTable, ChildrenAttribute),
		)
		dynamodbgraph.HasNeighborsWith(s, step, func(s *dynamodb.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPets applies the HasEdge predicate on the "pets" edge.
func HasPets() predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(PetsTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, PetsTable, PetsAttribute),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasPetsWith applies the HasEdge predicate on the "pets" edge with a given conditions (other predicates).
func HasPetsWith(preds ...predicate.Pet) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(PetsInverseTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, PetsTable, PetsAttribute),
		)
		dynamodbgraph.HasNeighborsWith(s, step, func(s *dynamodb.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s1 := s.Clone()
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		p(s.Not())
	})
}
