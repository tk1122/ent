// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/entc/integration/dynamodb/hooks/ent/predicate"
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

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldVersion, v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldName, v))
	})
}

// Worth applies equality check predicate on the "worth" field. It's identical to WorthEQ.
func Worth(v uint) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldWorth, v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldVersion, v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldVersion, v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.In(FieldVersion, v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NotIn(FieldVersion, v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldVersion, v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldVersion, v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldVersion, v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldVersion, v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldName, v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldName, v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.In(FieldName, v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NotIn(FieldName, v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldName, v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldName, v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldName, v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldName, v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.Contains(FieldName, v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.HasPrefix(FieldName, v))
	})
}

// WorthEQ applies the EQ predicate on the "worth" field.
func WorthEQ(v uint) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldWorth, v))
	})
}

// WorthNEQ applies the NEQ predicate on the "worth" field.
func WorthNEQ(v uint) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldWorth, v))
	})
}

// WorthIn applies the In predicate on the "worth" field.
func WorthIn(vs ...uint) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.In(FieldWorth, v...))
	})
}

// WorthNotIn applies the NotIn predicate on the "worth" field.
func WorthNotIn(vs ...uint) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NotIn(FieldWorth, v...))
	})
}

// WorthGT applies the GT predicate on the "worth" field.
func WorthGT(v uint) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldWorth, v))
	})
}

// WorthGTE applies the GTE predicate on the "worth" field.
func WorthGTE(v uint) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldWorth, v))
	})
}

// WorthLT applies the LT predicate on the "worth" field.
func WorthLT(v uint) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldWorth, v))
	})
}

// WorthLTE applies the LTE predicate on the "worth" field.
func WorthLTE(v uint) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldWorth, v))
	})
}

// HasCards applies the HasEdge predicate on the "cards" edge.
func HasCards() predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(CardsTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, CardsTable, CardsAttribute),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasCardsWith applies the HasEdge predicate on the "cards" edge with a given conditions (other predicates).
func HasCardsWith(preds ...predicate.Card) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(CardsInverseTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, CardsTable, CardsAttribute),
		)
		dynamodbgraph.HasNeighborsWith(s, step, func(s *dynamodb.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFriends applies the HasEdge predicate on the "friends" edge.
func HasFriends() predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(FriendsTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, false, true, FriendsTable, FriendsAttributes...),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasFriendsWith applies the HasEdge predicate on the "friends" edge with a given conditions (other predicates).
func HasFriendsWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(Table, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, false, true, FriendsTable, FriendsAttributes...),
		)
		dynamodbgraph.HasNeighborsWith(s, step, func(s *dynamodb.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBestFriend applies the HasEdge predicate on the "best_friend" edge.
func HasBestFriend() predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(BestFriendTable, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2O, false, true, BestFriendTable, BestFriendAttribute),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasBestFriendWith applies the HasEdge predicate on the "best_friend" edge with a given conditions (other predicates).
func HasBestFriendWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(Table, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2O, false, true, BestFriendTable, BestFriendAttribute),
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
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *dynamodb.Selector) {
		p(s.Not())
	})
}
