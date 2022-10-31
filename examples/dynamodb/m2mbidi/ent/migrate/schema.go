// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/dynamodb/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersAttributes  holds the attributes for the "users " table.
	UsersAttributes = []*schema.Attribute{
		{Name: "id", Type: field.TypeInt},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Attributes: UsersAttributes,
		PrimaryKey: []*schema.KeySchema{
			{AttributeName: "id", KeyType: schema.KeyType("HASH")},
		},
	}
	// UserFriendsAttributes  holds the attributes for the "user_friends " table.
	UserFriendsAttributes = []*schema.Attribute{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// UserFriendsTable holds the schema information for the "user_friends" table.
	UserFriendsTable = &schema.Table{
		Name:       "user_friends",
		Attributes: UserFriendsAttributes,
		PrimaryKey: []*schema.KeySchema{
			{AttributeName: "user_id", KeyType: schema.KeyType("HASH")},
			{AttributeName: "friend_id", KeyType: schema.KeyType("RANGE")},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
		UserFriendsTable,
	}
)
