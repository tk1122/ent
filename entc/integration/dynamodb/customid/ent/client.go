// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/entc/integration/dynamodb/customid/ent/migrate"
	uuid "github.com/satori/go.uuid"

	"entgo.io/ent/entc/integration/dynamodb/customid/ent/blob"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/car"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/group"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/pet"
	"entgo.io/ent/entc/integration/dynamodb/customid/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Blob is the client for interacting with the Blob builders.
	Blob *BlobClient
	// Car is the client for interacting with the Car builders.
	Car *CarClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Pet is the client for interacting with the Pet builders.
	Pet *PetClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Blob = NewBlobClient(c.config)
	c.Car = NewCarClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Pet = NewPetClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.DynamoDB:
		drv, err := dynamodb.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}

		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Blob.Use(hooks...)
	c.Car.Use(hooks...)
	c.Group.Use(hooks...)
	c.Pet.Use(hooks...)
	c.User.Use(hooks...)
}

// BlobClient is a client for the Blob schema.
type BlobClient struct {
	config
}

// NewBlobClient returns a client for the Blob from the given config.
func NewBlobClient(c config) *BlobClient {
	return &BlobClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `blob.Hooks(f(g(h())))`.
func (c *BlobClient) Use(hooks ...Hook) {
	c.hooks.Blob = append(c.hooks.Blob, hooks...)
}

// Create returns a builder for creating a Blob entity.
func (c *BlobClient) Create() *BlobCreate {
	mutation := newBlobMutation(c.config, OpCreate)
	return &BlobCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Blob entities.
func (c *BlobClient) CreateBulk(builders ...*BlobCreate) *BlobCreateBulk {
	return &BlobCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Blob.
func (c *BlobClient) Update() *BlobUpdate {
	mutation := newBlobMutation(c.config, OpUpdate)
	return &BlobUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BlobClient) UpdateOne(b *Blob) *BlobUpdateOne {
	mutation := newBlobMutation(c.config, OpUpdateOne, withBlob(b))
	return &BlobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BlobClient) UpdateOneID(id uuid.UUID) *BlobUpdateOne {
	mutation := newBlobMutation(c.config, OpUpdateOne, withBlobID(id))
	return &BlobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Blob.
func (c *BlobClient) Delete() *BlobDelete {
	mutation := newBlobMutation(c.config, OpDelete)
	return &BlobDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BlobClient) DeleteOne(b *Blob) *BlobDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BlobClient) DeleteOneID(id uuid.UUID) *BlobDeleteOne {
	builder := c.Delete().Where(blob.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BlobDeleteOne{builder}
}

// Query returns a query builder for Blob.
func (c *BlobClient) Query() *BlobQuery {
	return &BlobQuery{
		config: c.config,
	}
}

// Get returns a Blob entity by its id.
func (c *BlobClient) Get(ctx context.Context, id uuid.UUID) (*Blob, error) {
	return c.Query().Where(blob.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BlobClient) GetX(ctx context.Context, id uuid.UUID) *Blob {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Blob.
func (c *BlobClient) QueryParent(b *Blob) *BlobQuery {
	query := &BlobQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := b.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(blob.Table, blob.FieldID, id),
			dynamodbgraph.To(blob.Table, blob.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2O, false, true, blob.ParentTable, blob.ParentAttribute),
		)
		fromV = dynamodbgraph.Neighbors(step, b.config.driver)
		return fromV, nil
	}
	return query
}

// QueryLinks queries the links edge of a Blob.
func (c *BlobClient) QueryLinks(b *Blob) *BlobQuery {
	query := &BlobQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := b.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(blob.Table, blob.FieldID, id),
			dynamodbgraph.To(blob.Table, blob.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, false, true, blob.LinksTable, blob.LinksAttributes...),
		)
		fromV = dynamodbgraph.Neighbors(step, b.config.driver)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BlobClient) Hooks() []Hook {
	return c.hooks.Blob
}

// CarClient is a client for the Car schema.
type CarClient struct {
	config
}

// NewCarClient returns a client for the Car from the given config.
func NewCarClient(c config) *CarClient {
	return &CarClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `car.Hooks(f(g(h())))`.
func (c *CarClient) Use(hooks ...Hook) {
	c.hooks.Car = append(c.hooks.Car, hooks...)
}

// Create returns a builder for creating a Car entity.
func (c *CarClient) Create() *CarCreate {
	mutation := newCarMutation(c.config, OpCreate)
	return &CarCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Car entities.
func (c *CarClient) CreateBulk(builders ...*CarCreate) *CarCreateBulk {
	return &CarCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Car.
func (c *CarClient) Update() *CarUpdate {
	mutation := newCarMutation(c.config, OpUpdate)
	return &CarUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CarClient) UpdateOne(ca *Car) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCar(ca))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CarClient) UpdateOneID(id int) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCarID(id))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Car.
func (c *CarClient) Delete() *CarDelete {
	mutation := newCarMutation(c.config, OpDelete)
	return &CarDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CarClient) DeleteOne(ca *Car) *CarDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CarClient) DeleteOneID(id int) *CarDeleteOne {
	builder := c.Delete().Where(car.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CarDeleteOne{builder}
}

// Query returns a query builder for Car.
func (c *CarClient) Query() *CarQuery {
	return &CarQuery{
		config: c.config,
	}
}

// Get returns a Car entity by its id.
func (c *CarClient) Get(ctx context.Context, id int) (*Car, error) {
	return c.Query().Where(car.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CarClient) GetX(ctx context.Context, id int) *Car {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Car.
func (c *CarClient) QueryOwner(ca *Car) *PetQuery {
	query := &PetQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := ca.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(car.Table, car.FieldID, id),
			dynamodbgraph.To(pet.Table, pet.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2O, true, false, car.OwnerTable, car.OwnerAttribute),
		)
		fromV = dynamodbgraph.Neighbors(step, ca.config.driver)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CarClient) Hooks() []Hook {
	return c.hooks.Car
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a builder for creating a Group entity.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := gr.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(group.Table, group.FieldID, id),
			dynamodbgraph.To(user.Table, user.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, false, false, group.UsersTable, group.UsersAttributes...),
		)
		fromV = dynamodbgraph.Neighbors(step, gr.config.driver)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// PetClient is a client for the Pet schema.
type PetClient struct {
	config
}

// NewPetClient returns a client for the Pet from the given config.
func NewPetClient(c config) *PetClient {
	return &PetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pet.Hooks(f(g(h())))`.
func (c *PetClient) Use(hooks ...Hook) {
	c.hooks.Pet = append(c.hooks.Pet, hooks...)
}

// Create returns a builder for creating a Pet entity.
func (c *PetClient) Create() *PetCreate {
	mutation := newPetMutation(c.config, OpCreate)
	return &PetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Pet entities.
func (c *PetClient) CreateBulk(builders ...*PetCreate) *PetCreateBulk {
	return &PetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Pet.
func (c *PetClient) Update() *PetUpdate {
	mutation := newPetMutation(c.config, OpUpdate)
	return &PetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PetClient) UpdateOne(pe *Pet) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPet(pe))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PetClient) UpdateOneID(id string) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPetID(id))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Pet.
func (c *PetClient) Delete() *PetDelete {
	mutation := newPetMutation(c.config, OpDelete)
	return &PetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PetClient) DeleteOne(pe *Pet) *PetDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PetClient) DeleteOneID(id string) *PetDeleteOne {
	builder := c.Delete().Where(pet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PetDeleteOne{builder}
}

// Query returns a query builder for Pet.
func (c *PetClient) Query() *PetQuery {
	return &PetQuery{
		config: c.config,
	}
}

// Get returns a Pet entity by its id.
func (c *PetClient) Get(ctx context.Context, id string) (*Pet, error) {
	return c.Query().Where(pet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PetClient) GetX(ctx context.Context, id string) *Pet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Pet.
func (c *PetClient) QueryOwner(pe *Pet) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := pe.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(pet.Table, pet.FieldID, id),
			dynamodbgraph.To(user.Table, user.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2O, true, false, pet.OwnerTable, pet.OwnerAttribute),
		)
		fromV = dynamodbgraph.Neighbors(step, pe.config.driver)
		return fromV, nil
	}
	return query
}

// QueryCars queries the cars edge of a Pet.
func (c *PetClient) QueryCars(pe *Pet) *CarQuery {
	query := &CarQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := pe.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(pet.Table, pet.FieldID, id),
			dynamodbgraph.To(car.Table, car.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, pet.CarsTable, pet.CarsAttribute),
		)
		fromV = dynamodbgraph.Neighbors(step, pe.config.driver)
		return fromV, nil
	}
	return query
}

// QueryFriends queries the friends edge of a Pet.
func (c *PetClient) QueryFriends(pe *Pet) *PetQuery {
	query := &PetQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := pe.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(pet.Table, pet.FieldID, id),
			dynamodbgraph.To(pet.Table, pet.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, false, true, pet.FriendsTable, pet.FriendsAttributes...),
		)
		fromV = dynamodbgraph.Neighbors(step, pe.config.driver)
		return fromV, nil
	}
	return query
}

// QueryBestFriend queries the best_friend edge of a Pet.
func (c *PetClient) QueryBestFriend(pe *Pet) *PetQuery {
	query := &PetQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := pe.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(pet.Table, pet.FieldID, id),
			dynamodbgraph.To(pet.Table, pet.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2O, false, true, pet.BestFriendTable, pet.BestFriendAttribute),
		)
		fromV = dynamodbgraph.Neighbors(step, pe.config.driver)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PetClient) Hooks() []Hook {
	return c.hooks.Pet
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroups queries the groups edge of a User.
func (c *UserClient) QueryGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := u.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(user.Table, user.FieldID, id),
			dynamodbgraph.To(group.Table, group.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2M, true, false, user.GroupsTable, user.GroupsAttributes...),
		)
		fromV = dynamodbgraph.Neighbors(step, u.config.driver)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a User.
func (c *UserClient) QueryParent(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := u.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(user.Table, user.FieldID, id),
			dynamodbgraph.To(user.Table, user.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.M2O, true, false, user.ParentTable, user.ParentAttribute),
		)
		fromV = dynamodbgraph.Neighbors(step, u.config.driver)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a User.
func (c *UserClient) QueryChildren(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := u.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(user.Table, user.FieldID, id),
			dynamodbgraph.To(user.Table, user.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, user.ChildrenTable, user.ChildrenAttribute),
		)
		fromV = dynamodbgraph.Neighbors(step, u.config.driver)
		return fromV, nil
	}
	return query
}

// QueryPets queries the pets edge of a User.
func (c *UserClient) QueryPets(u *User) *PetQuery {
	query := &PetQuery{config: c.config}
	query.path = func(context.Context) (fromV *dynamodb.Selector, _ error) {
		id := u.ID
		step := dynamodbgraph.NewStep(
			dynamodbgraph.From(user.Table, user.FieldID, id),
			dynamodbgraph.To(pet.Table, pet.FieldID, []string{}),
			dynamodbgraph.Edge(dynamodbgraph.O2M, false, false, user.PetsTable, user.PetsAttribute),
		)
		fromV = dynamodbgraph.Neighbors(step, u.config.driver)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
