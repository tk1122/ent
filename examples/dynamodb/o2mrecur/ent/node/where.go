// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package node

import (
	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/examples/dynamodb/o2mrecur/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldID, id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldID, id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.In(FieldID, v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(dynamodb.NotIn(FieldID, v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldID, id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldID, id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldID, id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldID, id))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldValue, v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.EQ(FieldValue, v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NEQ(FieldValue, v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int) predicate.Node {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.In(FieldValue, v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int) predicate.Node {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.NotIn(FieldValue, v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GT(FieldValue, v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.GTE(FieldValue, v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LT(FieldValue, v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s.Where(dynamodb.LTE(FieldValue, v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(ParentAttribute, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2O, true, false, ParentAttribute, ParentAttribute),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
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
func HasChildren() predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(Table, FieldID),
			dynamodbgraph.To(ChildrenAttribute, FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, ChildrenAttribute, ChildrenAttribute),
		)
		dynamodbgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		s1 := s.Clone()
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
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
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(func(s *dynamodb.Selector) {
		p(s.Not())
	})
}
