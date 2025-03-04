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

// AppendOper wraps multiple operation builders like BatchWrite, TransactWrite.
type AppendOper interface {
	Oper
	Append(string, Oper) Oper
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
		tableName:         tableName,
		expBuilder:        expression.NewBuilder(),
		key:               make(map[string]types.AttributeValue),
		setAttributes:     make(map[string]interface{}),
		addedAttributes:   make(map[string]interface{}),
		removedAttributes: make(map[string]string),
		IsEmpty:           true,
	}
}

// GetItem returns a builder for GetItem operation.
func (d RootBuilder) GetItem() *GetItemBuilder {
	return &GetItemBuilder{
		key: make(map[string]types.AttributeValue),
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
	CreateTableOperation   = "CreateTable"
	GetItemOperation       = "GetItem"
	PutItemOperation       = "PutItem"
	UpdateItemOperation    = "UpdateItem"
	BatchWriteOperation    = "BatchWrite"
	TransactWriteOperation = "TransactWrite"
	ScanOperation          = "ScanOperation"
	DeleteItemOperation    = "DeleteItemOperation"
)

type (
	// CreateTableBuilder is the builder for CreateTable operation.
	CreateTableBuilder struct {
		attributeDefinitions  []types.AttributeDefinition
		keySchema             []types.KeySchemaElement
		provisionedThroughput *types.ProvisionedThroughput
		tableName             string
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
	c.provisionedThroughput = &types.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(int64(readCap)),
		WriteCapacityUnits: aws.Int64(int64(writeCap)),
	}
	return c
}

// Op returns name and input for CreateTable operation.
func (c *CreateTableBuilder) Op() (string, interface{}) {
	return CreateTableOperation, &dynamodb.CreateTableInput{
		TableName:             aws.String(c.tableName),
		AttributeDefinitions:  c.attributeDefinitions,
		KeySchema:             c.keySchema,
		ProvisionedThroughput: c.provisionedThroughput,
	}
}

type (
	// PutItemBuilder is the builder for PutItem operation.
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
	return PutItemOperation, &dynamodb.PutItemInput{
		TableName: aws.String(p.tableName),
		Item:      p.item,
	}
}

type (
	// UpdateItemBuilder is the builder for UpdateItem operation.
	UpdateItemBuilder struct {
		tableName         string
		key               map[string]types.AttributeValue
		setAttributes     map[string]interface{}
		addedAttributes   map[string]interface{}
		removedAttributes map[string]string
		expBuilder        expression.Builder
		exp               expression.Expression
		returnValues      types.ReturnValue
		IsEmpty           bool
	}
)

// WithKey selects which item to be updated for the UpdateItem operation.
func (u *UpdateItemBuilder) WithKey(k string, v types.AttributeValue) *UpdateItemBuilder {
	u.key[k] = v
	return u
}

// Set sets the attribute to be set.
func (u *UpdateItemBuilder) Set(k string, v interface{}) *UpdateItemBuilder {
	u.setAttributes[k] = v
	u.IsEmpty = false
	return u
}

// Add sets the attribute to be added.
func (u *UpdateItemBuilder) Add(k string, v interface{}) *UpdateItemBuilder {
	u.addedAttributes[k] = v
	u.IsEmpty = false
	return u
}

// Remove clears the attribute of the item.
func (u *UpdateItemBuilder) Remove(k string) *UpdateItemBuilder {
	u.removedAttributes[k] = k
	u.IsEmpty = false
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
	for attr, val := range u.setAttributes {
		if i == 0 {
			setExp = expression.Set(expression.Name(attr), expression.Value(val))
		} else {
			setExp = setExp.Set(expression.Name(attr), expression.Value(val))
			expression.Remove(expression.Name(attr))
		}
		i += 1
	}
	for _, attr := range u.removedAttributes {
		if _, ok := u.setAttributes[attr]; ok {
			continue
		}
		if i == 0 {
			setExp = expression.Remove(expression.Name(attr))
		} else {
			setExp = setExp.Remove(expression.Name(attr))
		}
		i += 1
	}
	for attr, val := range u.addedAttributes {
		if _, ok := u.setAttributes[attr]; ok {
			continue
		}
		if _, ok := u.removedAttributes[attr]; ok {
			continue
		}
		if i == 0 {
			setExp = expression.Add(expression.Name(attr), expression.Value(val))
		} else {
			setExp = setExp.Add(expression.Name(attr), expression.Value(val))
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
		IsEmpty    bool
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
		IsEmpty:    true,
	}
}

// Append appends a WriteRequest to BatchWriteItem.
func (b *BatchWriteItemBuilder) Append(tableName string, op Oper) Oper {
	b.requestMap[tableName] = append(b.requestMap[tableName], op)
	b.IsEmpty = false
	return b
}

// Op returns name and input for BatchWriteItem operation.
func (b *BatchWriteItemBuilder) Op() (string, interface{}) {
	return BatchWriteOperation, &BatchWriteItemArgs{
		operationMap: b.requestMap,
	}
}

type (
	// TransactWriteItemBuilder is the builder for TransactWriteItemBuilder operation.
	TransactWriteItemBuilder struct {
		requestMap map[string][]Oper
		IsEmpty    bool
	}

	// TransactWriteItemArgs contains input of TransactWriteItem operation.
	TransactWriteItemArgs struct {
		operationMap map[string][]Oper
	}
)

// TransactWriteItem returns a builder for TransactWriteItem operation.
func TransactWriteItem() *TransactWriteItemBuilder {
	return &TransactWriteItemBuilder{
		requestMap: make(map[string][]Oper),
		IsEmpty:    true,
	}
}

// Append appends a WriteRequest to TransactWriteItem.
func (b *TransactWriteItemBuilder) Append(tableName string, op Oper) Oper {
	b.requestMap[tableName] = append(b.requestMap[tableName], op)
	b.IsEmpty = false
	return b
}

// Op returns name and input for TransactWriteItem operation.
func (b *TransactWriteItemBuilder) Op() (string, interface{}) {
	return TransactWriteOperation, &TransactWriteItemArgs{
		operationMap: b.requestMap,
	}
}

type (
	// GetItemBuilder is the builder for GetItem operation.
	GetItemBuilder struct {
		tableName string
		key       map[string]types.AttributeValue
	}
)

// From selects which table for the GetItem operation.
func (u *GetItemBuilder) From(tableName string) *GetItemBuilder {
	u.tableName = tableName
	return u
}

// WithKey selects which item to be selected for the GetItem operation.
func (u *GetItemBuilder) WithKey(k string, v types.AttributeValue) *GetItemBuilder {
	u.key[k] = v
	return u
}

// Op returns name and input for GetItem operation.
func (u *GetItemBuilder) Op() (string, interface{}) {
	return GetItemOperation, &dynamodb.GetItemInput{
		TableName: aws.String(u.tableName),
		Key:       u.key,
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
	return DeleteItemOperation, &dynamodb.DeleteItemInput{
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
	orderDesc      bool
	errs           []error // errors that added during the selection construction.
}

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

// Op returns name and args of Scan operation.
func (s *Selector) Op() (string, interface{}) {
	return ScanOperation, &dynamodb.ScanInput{
		TableName:                 aws.String(s.table),
		ProjectionExpression:      s.exp.Projection(),
		FilterExpression:          s.exp.Filter(),
		ExpressionAttributeNames:  s.exp.Names(),
		ExpressionAttributeValues: s.exp.Values(),
	}
}
