// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "user_friends"
)

// Keys holds all DynamoDB keys for user fields.
var Keys = []string{
	FieldID,
	FieldAge,
	FieldName,
	"friend_id",
}

var (
	// FriendsAttribute and FriendsAttribute2 are the collection keys denoting the
	// primary key for the friends relation (M2M).
	FriendsAttributes = []string{"user_id", "friend_id"}
)

// ValidKey reports if the key is valid (one of keys of collection's fields).
func ValidKey(key string) bool {
	for i := range Keys {
		if key == Keys[i] {
			return true
		}
	}
	return false
}
