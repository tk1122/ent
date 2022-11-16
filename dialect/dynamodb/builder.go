// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dynamodb

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Oper wraps the basic stage implementation.
type Oper interface {
	// Op returns the database's operation representation.
	Op() (string, interface{})
}

// RootBuilder is the constructor for all DynamoDB operation builders.
type RootBuilder struct{}

// CreateTable returns a builder for CreateTable operation.
func (d RootBuilder) CreateTable(name string) *CreateTableBuilder {
	return &CreateTableBuilder{
		tableName: name,
	}
}

// PutItem returns a builder for PutItem operation.
func (d RootBuilder) PutItem(tableName string) *PutItemBuilder {
	return &PutItemBuilder{
		tableName: tableName,
	}
}

// DeleteItem returns a builder for DeleteItem operation.
func (d RootBuilder) DeleteItem(tableName string) *DeleteItemBuilder {
	return &DeleteItemBuilder{
		tableName:  tableName,
		key:        make(map[string]types.AttributeValue),
		expBuilder: expression.NewBuilder(),
	}
}

// Update returns a builder for UpdateItem operation.
func (d RootBuilder) Update(tableName string) *UpdateItemBuilder {
	return &UpdateItemBuilder{
		tableName:          tableName,
		expBuilder:         expression.NewBuilder(),
		key:                make(map[string]types.AttributeValue),
		updateAttributeMap: make(map[string]interface{}),
	}
}

// Select returns a builder for Select operation.
func (d RootBuilder) Select(keys ...string) *Selector {
	return (&Selector{
		expBuilder:     expression.NewBuilder(),
		isBuilderEmpty: true,
	}).Select(keys...)
}

const (
	CreateTableOperation = "CreateTable"
	PutItemOperation     = "PutItem"
	UpdateItemOperation  = "UpdateItem"
	BatchWriteOperation  = "BatchWrite"
	ScanOperation        = "ScanOperation"
	DeletItemOperation   = "DeleteItemOperation"
)

type (
	// CreateTableArgs contains input for CreateTable operation.
	CreateTableArgs struct {
		Name string
		Opts *dynamodb.CreateTableInput
	}

	// CreateTableBuilder is the builder for CreateTableArgs.
	CreateTableBuilder struct {
		attributeDefinitions []types.AttributeDefinition
		keySchema            []types.KeySchemaElement
		provisiondThroughput *types.ProvisionedThroughput
		tableName            string
	}
)

// AddAttribute adds one attribute to the table.
func (c *CreateTableBuilder) AddAttribute(attributeName string, attributeType types.ScalarAttributeType) *CreateTableBuilder {
	c.attributeDefinitions = append(c.attributeDefinitions, types.AttributeDefinition{
		AttributeName: aws.String(attributeName),
		AttributeType: attributeType,
	})
	return c
}

// AddKeySchemaElement adds one element to the key schema of the table.
func (c *CreateTableBuilder) AddKeySchemaElement(attributeName string, keyType types.KeyType) *CreateTableBuilder {
	c.keySchema = append(c.keySchema, types.KeySchemaElement{
		AttributeName: aws.String(attributeName),
		KeyType:       keyType,
	})
	return c
}

// SetProvisionedThroughput sets the provisioned throughput of the table.
func (c *CreateTableBuilder) SetProvisionedThroughput(readCap, writeCap int) *CreateTableBuilder {
	c.provisiondThroughput = &types.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(int64(readCap)),
		WriteCapacityUnits: aws.Int64(int64(writeCap)),
	}
	return c
}

// Op returns name and input for CreateTable operation.
func (c *CreateTableBuilder) Op() (string, interface{}) {
	return CreateTableOperation, &CreateTableArgs{
		Name: c.tableName,
		Opts: &dynamodb.CreateTableInput{
			TableName:             aws.String(c.tableName),
			AttributeDefinitions:  c.attributeDefinitions,
			KeySchema:             c.keySchema,
			ProvisionedThroughput: c.provisiondThroughput,
		},
	}
}

type (
	// PutItemArgs contains input for PutItem operation.
	PutItemArgs struct {
		Opts *dynamodb.PutItemInput
	}

	// PutItemBuilder is the builder for PutItemArgs.
	PutItemBuilder struct {
		item      map[string]types.AttributeValue
		tableName string
	}
)

// SetItem provides the data of the item to be put to that table.
func (p *PutItemBuilder) SetItem(i map[string]types.AttributeValue) *PutItemBuilder {
	p.item = i
	return p
}

// Op returns name and input for PutItem operation.
func (p *PutItemBuilder) Op() (string, interface{}) {
	return PutItemOperation, &PutItemArgs{
		Opts: &dynamodb.PutItemInput{
			TableName: aws.String(p.tableName),
			Item:      p.item,
		},
	}
}

type (
	// UpdateItemBuilder is the builder for UpdateItem operation.
	UpdateItemBuilder struct {
		tableName          string
		key                map[string]types.AttributeValue
		updateAttributeMap map[string]interface{}
		removeAttributes   []string
		expBuilder         expression.Builder
		exp                expression.Expression
		returnValues       types.ReturnValue
	}
)

// WithKey selects which item to be updated for the UpdateItem operation.
func (u *UpdateItemBuilder) WithKey(k string, v types.AttributeValue) *UpdateItemBuilder {
	u.key[k] = v
	return u
}

// Set sets the attribute to be updated.
func (u *UpdateItemBuilder) Set(k string, v interface{}) *UpdateItemBuilder {
	u.updateAttributeMap[k] = v
	return u
}

// Remove clears the attribute of the item.
func (u *UpdateItemBuilder) Remove(k string) *UpdateItemBuilder {
	u.removeAttributes = append(u.removeAttributes, k)
	return u
}

// Where sets the expression for the UpdateItem operation.
func (u *UpdateItemBuilder) Where(p *Predicate) *UpdateItemBuilder {
	u.expBuilder = u.expBuilder.WithCondition(p.Query())
	return u
}

// BuildExpression returns potential errors during UpdateItemBuilder's build process.
func (u *UpdateItemBuilder) BuildExpression(rv types.ReturnValue) (*UpdateItemBuilder, error) {
	var err error
	u.returnValues = rv
	var setExp expression.UpdateBuilder
	i := 0
	for attr, val := range u.updateAttributeMap {
		if i == 0 {
			setExp = expression.Set(expression.Name(attr), expression.Value(val))
		} else {
			setExp = setExp.Set(expression.Name(attr), expression.Value(val))
			expression.Remove(expression.Name(attr))
		}
		i += 1
	}
	for _, attr := range u.removeAttributes {
		if i == 0 {
			setExp = expression.Remove(expression.Name(attr))
		} else {
			setExp = setExp.Remove(expression.Name(attr))
		}
		i += 1
	}
	u.expBuilder = u.expBuilder.WithUpdate(setExp)
	u.exp, err = u.expBuilder.Build()
	return u, err
}

// Op returns name and input for UpdateItem operation.
func (u *UpdateItemBuilder) Op() (string, interface{}) {
	return UpdateItemOperation, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(u.tableName),
		Key:                       u.key,
		ConditionExpression:       u.exp.Condition(),
		UpdateExpression:          u.exp.Update(),
		ExpressionAttributeNames:  u.exp.Names(),
		ExpressionAttributeValues: u.exp.Values(),
		ReturnValues:              u.returnValues,
	}
}

type (
	// BatchWriteItemBuilder is the builder for BatchWriteItem operation.
	BatchWriteItemBuilder struct {
		requestMap map[string][]Oper
	}

	// BatchWriteItemArgs contains input of BatchWriteItem operation.
	BatchWriteItemArgs struct {
		operationMap map[string][]Oper
	}
)

// BatchWriteItem returns a builder for BatchWriteItem operation.
func BatchWriteItem() *BatchWriteItemBuilder {
	return &BatchWriteItemBuilder{
		requestMap: make(map[string][]Oper),
	}
}

// Append appends a WriteRequest to BatchWriteItem.
func (b *BatchWriteItemBuilder) Append(tableName string, op Oper) *BatchWriteItemBuilder {
	b.requestMap[tableName] = append(b.requestMap[tableName], op)
	return b
}

// Op returns name and input for BatchWriteItem operation.
func (b *BatchWriteItemBuilder) Op() (string, interface{}) {
	return BatchWriteOperation, &BatchWriteItemArgs{
		operationMap: b.requestMap,
	}
}

type (
	// DeleteItemBuilder is the builder for DeleteItem operation.
	DeleteItemBuilder struct {
		tableName  string
		expBuilder expression.Builder
		exp        expression.Expression
		key        map[string]types.AttributeValue
	}
)

// WithKey selects which item to be deleted for the DeleteItem operation.
func (u *DeleteItemBuilder) WithKey(k string, v types.AttributeValue) *DeleteItemBuilder {
	u.key[k] = v
	return u
}

// Where sets the expression for the DeleteItem operation.
func (u *DeleteItemBuilder) Where(p *Predicate) *DeleteItemBuilder {
	u.expBuilder = u.expBuilder.WithCondition(p.Query())
	return u
}

// Op returns name and input for DeleteItem operation.
func (u *DeleteItemBuilder) Op() (string, interface{}) {
	return DeletItemOperation, &dynamodb.DeleteItemInput{
		TableName: aws.String(u.tableName),
		Key:       u.key,
	}
}

// Selector is a builder for the `SELECT` statement.
type Selector struct {
	ctx            context.Context
	table          string
	query          *Predicate
	expBuilder     expression.Builder
	exp            expression.Expression
	isBuilderEmpty bool
	not            bool
	or             bool
	orderDesc      bool
	errs           []error // errors that added during the selection construction.
}

// ordering direction aliases.
const (
	Asc  = true
	Desc = false
)

func Select(keys ...string) *Selector {
	return (&Selector{
		expBuilder:     expression.NewBuilder(),
		isBuilderEmpty: true,
	}).Select(keys...)
}

// From sets collection name of the selector.
func (s *Selector) From(table string) *Selector {
	s.table = table
	return s
}

// Where sets or appends the given predicate to the selector.
func (s *Selector) Where(p *Predicate) *Selector {
	if p == nil {
		return s
	}
	if s.not {
		p = Not(p)
	}
	cond := Clone(p)
	switch {
	case s.query == nil:
		s.query = cond
	default:
		s.query = And(cond, s.query)
	}
	s.expBuilder = s.expBuilder.WithFilter(s.query.Query())
	s.isBuilderEmpty = false
	return s
}

// Clone returns a duplicate of the selector, including all associated steps. It can be
// used to prepare common SELECT statements and use them differently after the clone is made.
func (s *Selector) Clone() *Selector {
	return &Selector{
		table:          s.table,
		query:          s.query.clone(),
		errs:           append(s.errs[:0:0], s.errs...),
		isBuilderEmpty: s.isBuilderEmpty,
	}
}

// P returns the predicate of a selector.
func (s *Selector) P() *Predicate {
	return s.query
}

// Or sets the next coming predicate with OR operator (disjunction).
func (s *Selector) Or() *Selector {
	s.or = true
	return s
}

// Not sets the next coming predicate with not.
func (s *Selector) Not() *Selector {
	s.not = true
	return s
}

// OrderBy appends the orders to query statement.
func (s *Selector) OrderBy(asc bool) *Selector {
	s.orderDesc = !asc
	return s
}

// Select adds the keys selection of the selector.
func (s *Selector) Select(keys ...string) *Selector {
	if len(keys) == 0 {
		return s
	}
	proj := expression.NamesList(expression.Name(keys[0]))
	for _, k := range keys[1:] {
		proj = proj.AddNames(expression.Name(k))
	}
	s.expBuilder = s.expBuilder.WithProjection(proj)
	s.isBuilderEmpty = false
	return s
}

// BuildExpressions builds and validates the expression.
func (s *Selector) BuildExpressions() *Selector {
	if s.isBuilderEmpty {
		return s
	}
	var err error
	s.exp, err = s.expBuilder.Build()
	if err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

// Err returns a concatenated error of all errors encountered during
// the selection-building, or were added manually by calling AddError.
func (s *Selector) Err() error {
	if s.errs == nil {
		return nil
	}
	br := strings.Builder{}
	for i := range s.errs {
		if i > 0 {
			br.WriteString("; ")
		}
		br.WriteString(s.errs[i].Error())
	}
	return fmt.Errorf(br.String())
}

// AddError appends an error to the selector errors.
func (s *Selector) AddError(err error) *Selector {
	s.errs = append(s.errs, err)
	return s
}

// Op retutns name and args of Scan operation.
func (s *Selector) Op() (string, interface{}) {
	return ScanOperation, &dynamodb.ScanInput{
		TableName:                 aws.String(s.table),
		ProjectionExpression:      s.exp.Projection(),
		FilterExpression:          s.exp.Filter(),
		ExpressionAttributeNames:  s.exp.Names(),
		ExpressionAttributeValues: s.exp.Values(),
	}
}
