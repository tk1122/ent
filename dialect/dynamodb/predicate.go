package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
)

// Predicate is a query predicate.
type Predicate struct {
	QueryBuilder
	fns []func() expression.ConditionBuilder
}

// QueryBuilder is the builder for DynamoDB condition expression.
type QueryBuilder struct {
	condBuilder expression.ConditionBuilder
}

// P creates a new predicate.
//
// P().EQ("name", "tk1122").And().EQ("age", 24)
func P() *Predicate {
	return &Predicate{}
}

// Query runs all appended build steps.
func (p *Predicate) Query() expression.ConditionBuilder {
	p.condBuilder = p.fns[0]()
	for _, f := range p.fns[1:] {
		p.condBuilder = expression.And(p.condBuilder, f())
	}
	return p.condBuilder
}

// Append appends match builders to predicate.
func (p *Predicate) Append(fs ...func() expression.ConditionBuilder) *Predicate {
	p.fns = append(p.fns, fs...)
	return p
}

// EQ returns `expression.Name(key).Equal(expression.Value(val))` predicate.
func EQ(key string, val interface{}) *Predicate {
	return P().EQ(key, val)
}

// EQ appends a "expression.Name(key).Equal(expression.Value(val))" predicate.
func (p *Predicate) EQ(key string, val interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).Equal(expression.Value(val))
	})
}

// In returns the `expression.Name(key).In(expression.Value(vals))` predicate.
func In(key string, vals ...interface{}) *Predicate {
	return P().In(key, vals...)
}

// In appends the `expression.Name(key).In(expression.Value(vals))` predicate.
func (p *Predicate) In(key string, vals ...interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		if len(vals) == 0 {
			return expression.Name(key).In(expression.Value(vals))
		}
		exp := expression.Name(key).In(expression.Value(vals[0]))
		if len(vals) == 1 {
			return exp
		}
		var exValues []expression.OperandBuilder
		for _, val := range vals[1:] {
			exValues = append(exValues, expression.Value(val))
		}
		exp = expression.Name(key).In(expression.Value(vals[0]), exValues...)
		return exp
	})
}

// And combines all given predicates with expression.And between them.
func And(preds ...*Predicate) *Predicate {
	return P().Append(func() expression.ConditionBuilder {
		cond := preds[0].Query()
		for _, pred := range preds[1:] {
			cond = expression.And(cond, pred.Query())
		}
		return cond
	})
}

// Or combines all given predicates with expression.Or between them.
func Or(preds ...*Predicate) *Predicate {
	return P().Append(func() expression.ConditionBuilder {
		cond := preds[0].Query()
		for _, pred := range preds[1:] {
			cond = expression.Or(cond, pred.Query())
		}
		return cond
	})
}

// NotExist returns the `expression.Not(expression.Name(key).AttributeExists())` predicate.
func NotExist(col string) *Predicate {
	return P().NotExist(col)
}

// NotExist appends the `expression.Not(expression.Name(key).AttributeExists())` predicate.
func (p *Predicate) NotExist(key string) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Not(expression.Name(key).AttributeExists())
	})
}

// Exist returns the `expression.Name(key).AttributeExists()` predicate.
func Exist(col string) *Predicate {
	return P().Exist(col)
}

// Exist appends the `expression.Name(key).AttributeExists()` predicate.
func (p *Predicate) Exist(key string) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).AttributeExists()
	})
}

// Not reverses the logic of the predicate.
//
//	Not(Or(EQ("name", "foo"), EQ("name", "bar")))
func Not(pred *Predicate) *Predicate {
	return P().Append(func() expression.ConditionBuilder {
		cond := pred.Query()
		return expression.Not(cond)
	})
}

// Clone creates a clone of a predicate.
func Clone(pred *Predicate) *Predicate {
	return P().Append(func() expression.ConditionBuilder {
		cond := pred.Query()
		return cond
	})
}

// NEQ returns a "expression.Not(expression.Name(key).Equal(expression.Value(val)))" predicate.
func NEQ(key string, val interface{}) *Predicate {
	return P().NEQ(key, val)
}

// NEQ appends a "expression.Not(expression.Name(key).Equal(expression.Value(val)))" predicate.
func (p *Predicate) NEQ(key string, val interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Not(expression.Name(key).Equal(expression.Value(val)))
	})
}

// NotIn returns the `expression.Not(expression.Name(key).In(expression.Value(vals)))` predicate.
func NotIn(key string, vals ...interface{}) *Predicate {
	return P().NotIn(key, vals...)
}

// NotIn appends the `expression.Not(expression.Name(key).In(expression.Value(vals)))` predicate.
func (p *Predicate) NotIn(key string, vals ...interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Not(expression.Name(key).In(expression.Value(vals)))
	})
}

// GT returns a `expression.Name(key).GreaterThan(expression.Value(val))` predicate.
func GT(key string, val interface{}) *Predicate {
	return P().GT(key, val)
}

// GT appends a `expression.Name(key).GreaterThan(expression.Value(val))` predicate.
func (p *Predicate) GT(key string, val interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).GreaterThan(expression.Value(val))
	})
}

// GTE returns a `expression.Name(key).GreaterThanEqual(expression.Value(val))` predicate.
func GTE(key string, val interface{}) *Predicate {
	return P().GTE(key, val)
}

// GTE appends a `expression.Name(key).GreaterThanEqual(expression.Value(val))` predicate.
func (p *Predicate) GTE(key string, val interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).GreaterThanEqual(expression.Value(val))
	})
}

// LT returns a `expression.Name(key).LessThan(expression.Value(val))` predicate.
func LT(key string, val interface{}) *Predicate {
	return P().LT(key, val)
}

// LT appends a `expression.Name(key).LessThan(expression.Value(val))` predicate.
func (p *Predicate) LT(key string, val interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).LessThan(expression.Value(val))
	})
}

// LTE returns a `expression.Name(key).LessThan(expression.Value(val))` predicate.
func LTE(key string, val interface{}) *Predicate {
	return P().LTE(key, val)
}

// LTE appends a `expression.Name(key).LessThanEqual(expression.Value(val))` predicate.
func (p *Predicate) LTE(key string, val interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).LessThanEqual(expression.Value(val))
	})
}

// Contains returns a `expression.Name(key).Contains(val)` predicate.
func Contains(key, val string) *Predicate {
	return P().Contains(key, val)
}

// Contains appends a `expression.Name(key).Contains(val)` predicate.
func (p *Predicate) Contains(key, val string) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).Contains(val)
	})
}

// HasPrefix returns a `expression.Name(key).BeginsWith(val)` predicate.
func HasPrefix(key, val string) *Predicate {
	return P().HasPrefix(key, val)
}

// HasPrefix appends a `expression.Name(key).BeginsWith(val)` predicate.
func (p *Predicate) HasPrefix(key, val string) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).BeginsWith(val)
	})
}

// clone returns a shallow clone of p.
func (p *Predicate) clone() *Predicate {
	if p == nil {
		return p
	}
	return &Predicate{
		fns: append(p.fns[:0:0], p.fns...),
	}
}
