// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dynamodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	uuid "github.com/satori/go.uuid"

	"entgo.io/ent/dialect"
)

// Driver is a dialect.Driver implementation for DynamoDB based databases.
type Driver struct {
	Client
	dialect string
}

// NewDriver creates a new Driver with the given Conn and dialect.
func NewDriver(dialect string, c Client) *Driver {
	return &Driver{dialect: dialect, Client: c}
}

// Open returns a dialect.Driver that implements the ent/dialect.Driver interface.
func Open(dialect, source string) (*Driver, error) {
	var (
		awsCfg                aws.Config
		err                   error
		endpointUrlOptionFunc config.LoadOptionsFunc
	)
	if source != "" {
		endpointUrlOptionFunc = config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: source}, nil
				}),
		)
		awsCfg, err = config.LoadDefaultConfig(context.TODO(), endpointUrlOptionFunc)
	} else {
		awsCfg, err = config.LoadDefaultConfig(context.TODO())
	}
	if err != nil {
		return nil, err
	}
	// Using the Config value, create the DynamoDB client
	dynamoDBClient := dynamodb.NewFromConfig(awsCfg)
	return NewDriver(dialect, Client{dynamoDBClient}), nil
}

// Dialect implements the dialect.Dialect method.
func (d Driver) Dialect() string {
	return d.dialect
}

// Tx starts and returns a transaction.
func (d *Driver) Tx(ctx context.Context) (dialect.Tx, error) {
	return dialect.NopTx(d), nil
}

// Close is a noop operation when the dialect is dynamodb
func (d *Driver) Close() error {
	return nil
}

// Exec implements the dialect.Exec method.
func (c Client) Exec(ctx context.Context, op string, args, v any) error {
	return c.run(ctx, op, args, v)
}

// Query implements the dialect.Query method.
func (c Client) Query(ctx context.Context, op string, args, v any) error {
	return c.run(ctx, op, args, v)
}

// run decides which SDK command to call
func (c Client) run(ctx context.Context, op string, args, v interface{}) error {
	switch op {
	case CreateTableOperation:
		return c.createTable(ctx, args, v)
	case PutItemOperation:
		return c.putItem(ctx, args, v)
	case GetItemOperation:
		return c.getItem(ctx, args, v)
	case UpdateItemOperation:
		return c.updateItem(ctx, args, v)
	case BatchWriteOperation:
		return c.batchWrite(ctx, args, v)
	case TransactWriteOperation:
		return c.transactWrite(ctx, args, v)
	case ScanOperation:
		return c.scan(ctx, args, v)
	case DeleteItemOperation:
		return c.deleteItem(ctx, args, v)
	default:
		return fmt.Errorf("%s operation is unsupported", op)
	}
}

func (c Client) createTable(ctx context.Context, args, v interface{}) (err error) {
	createTableArgs := args.(*dynamodb.CreateTableInput)
	output := v.(*dynamodb.CreateTableOutput)
	res, err := c.CreateTable(ctx, createTableArgs)
	if err != nil {
		var inUseEx *types.ResourceInUseException
		if !errors.As(err, &inUseEx) {
			return fmt.Errorf("DynamoDB CreateTable operation: %v", err)
		}
	}
	waiter := dynamodb.NewTableExistsWaiter(c)
	err = waiter.Wait(ctx, &dynamodb.DescribeTableInput{
		TableName: createTableArgs.TableName,
	}, 30*time.Second)
	if err != nil {
		return fmt.Errorf("DynamoDB CreateTable operation: %v", err)
	}
	// in case err is a ResourceInUseException
	if res != nil {
		*output = *res
	}
	return nil
}

func (c Client) putItem(ctx context.Context, args, v interface{}) (err error) {
	output := v.(*dynamodb.PutItemOutput)
	putItemArgs := args.(*dynamodb.PutItemInput)
	res, err := c.PutItem(ctx, putItemArgs)
	if err != nil {
		return fmt.Errorf("DynamoDB PutItem operation: %v", err)
	}
	*output = *res
	return nil
}

func (c Client) getItem(ctx context.Context, args, v interface{}) (err error) {
	output := v.(*dynamodb.GetItemOutput)
	input := args.(*dynamodb.GetItemInput)
	res, err := c.GetItem(ctx, input)
	if err != nil {
		return fmt.Errorf("DynamoDB GetItem operation: %v", err)
	}
	*output = *res
	return nil
}

func (c Client) updateItem(ctx context.Context, args, v interface{}) (err error) {
	output := v.(*dynamodb.UpdateItemOutput)
	input := args.(*dynamodb.UpdateItemInput)
	res, err := c.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("DynamoDB UpdateItem operation: %v", err)
	}
	*output = *res
	return nil
}

func (c Client) batchWrite(ctx context.Context, args, v interface{}) (err error) {
	batchWriteArgs := args.(*BatchWriteItemArgs)
	requestItems := make(map[string][]types.WriteRequest)
	for tableName, ops := range batchWriteArgs.operationMap {
		for _, op := range ops {
			opName, opArgs := op.Op()
			switch opName {
			case PutItemOperation:
				input := opArgs.(*dynamodb.PutItemInput)
				requestItems[tableName] = append(requestItems[tableName], types.WriteRequest{
					PutRequest: &types.PutRequest{Item: input.Item},
				})
			case DeleteItemOperation:
				input := opArgs.(*dynamodb.DeleteItemInput)
				requestItems[tableName] = append(requestItems[tableName], types.WriteRequest{
					DeleteRequest: &types.DeleteRequest{Key: input.Key},
				})
			}
		}
	}
	_, err = c.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: requestItems,
	})
	if err != nil {
		return fmt.Errorf("DynamoDB BatchWriteItem operation: %v", err)
	}
	return nil
}

func (c Client) transactWrite(ctx context.Context, args, v interface{}) (err error) {
	transactWriteArgs := args.(*TransactWriteItemArgs)
	var requestItems []types.TransactWriteItem
	for _, ops := range transactWriteArgs.operationMap {
		for _, op := range ops {
			opName, opArgs := op.Op()
			switch opName {
			case PutItemOperation:
				input := opArgs.(*dynamodb.PutItemInput)
				requestItems = append(requestItems, types.TransactWriteItem{
					Put: &types.Put{
						TableName: input.TableName,
						Item:      input.Item,
					},
				})
			case DeleteItemOperation:
				input := opArgs.(*dynamodb.DeleteItemInput)
				requestItems = append(requestItems, types.TransactWriteItem{
					Delete: &types.Delete{
						TableName: input.TableName,
						Key:       input.Key,
					},
				})
			case UpdateItemOperation:
				input := opArgs.(*dynamodb.UpdateItemInput)
				requestItems = append(requestItems, types.TransactWriteItem{
					Update: &types.Update{
						TableName:                 input.TableName,
						Key:                       input.Key,
						ConditionExpression:       input.ConditionExpression,
						UpdateExpression:          input.UpdateExpression,
						ExpressionAttributeNames:  input.ExpressionAttributeNames,
						ExpressionAttributeValues: input.ExpressionAttributeValues,
					},
				})
			}
		}
	}
	_, err = c.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems:      requestItems,
		ClientRequestToken: aws.String(uuid.NewV4().String()),
	})
	if err != nil {
		return fmt.Errorf("DynamoDB TransactWriteItem operation: %v", err)
	}
	return nil
}

func (c Client) scan(ctx context.Context, args, v interface{}) (err error) {
	output := v.(*dynamodb.ScanOutput)
	input := args.(*dynamodb.ScanInput)
	accummulatedScanOutput := dynamodb.ScanOutput{}
	scanOutput, err := c.Scan(ctx, input)
	if err != nil {
		return fmt.Errorf("DynamoDB Scan operation: %v", err)
	}
	accummulatedScanOutput.Count += scanOutput.Count
	accummulatedScanOutput.Items = append(accummulatedScanOutput.Items, scanOutput.Items...)
	for len(scanOutput.LastEvaluatedKey) != 0 {
		input.ExclusiveStartKey = scanOutput.LastEvaluatedKey
		scanOutput, err = c.Scan(ctx, input)
		if err != nil {
			return fmt.Errorf("DynamoDB Scan operation: %v", err)
		}
		accummulatedScanOutput.Count += scanOutput.Count
		accummulatedScanOutput.Items = append(accummulatedScanOutput.Items, scanOutput.Items...)
	}
	*output = accummulatedScanOutput
	return nil
}

func (c Client) deleteItem(ctx context.Context, args, v interface{}) (err error) {
	output := v.(*dynamodb.DeleteItemOutput)
	input := args.(*dynamodb.DeleteItemInput)
	deleteOutput, err := c.DeleteItem(ctx, input)
	if err != nil {
		return fmt.Errorf("DynamoDB DeleteItem operation: %v", err)
	}
	*output = *deleteOutput
	return nil
}

// Client wrap a DynamoDB client from AWS SDK to implement dialect.ExecQuerier
type Client struct {
	*dynamodb.Client
}

var (
	_ dialect.Driver = (*Driver)(nil)
)
