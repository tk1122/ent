package gen

import (
	"fmt"

	dyschema "entgo.io/ent/dialect/dynamodb/schema"
)

// DyAttribute converts template Field to DynamoDB attribute.
func (f Field) DyAttribute() *dyschema.Attribute {
	return &dyschema.Attribute{
		Name: f.Name,
		Type: f.Type.Type,
	}
}

// AttributeConstant returns the attribute name of the relation attribute.
func (e Edge) AttributeConstant() string { return pascal(e.Name) + "Attribute" }

// AttributesConstant returns the attribute name of the relation attributes. Used for M2M edges.
func (e Edge) AttributesConstant() string { return pascal(e.Name) + "Attributes" }

// Key returns the relation key of "from" table. Used for M2M edges.
func (e Edge) Key() string {
	if e.IsInverse() {
		return e.Rel.Keys()[0]
	}

	return e.Rel.Keys()[1]
}

// Keys returns the relation key in the relation tables.
func (r Relation) Keys() []string {
	if len(r.Columns) == 0 {
		panic(fmt.Sprintf("missing keys for Relation.Table: %s", r.Table))
	}

	return r.Columns
}
