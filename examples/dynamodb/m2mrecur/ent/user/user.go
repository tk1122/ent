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
	// EdgeFollowers holds the string denoting the followers edge name in mutations.
	EdgeFollowers = "followers"
	// EdgeFollowing holds the string denoting the following edge name in mutations.
	EdgeFollowing = "following"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FollowersTable is the table that holds the followers relation/edge. The primary key declared below.
	FollowersTable = "user_following"
	// FollowingTable is the table that holds the following relation/edge. The primary key declared below.
	FollowingTable = "user_following"
)

// Keys holds all DynamoDB keys for user fields.
var Keys = []string{
	FieldID,
	FieldAge,
	FieldName,
	"user_id",
	"follower_id",
}

var (
	// FollowersAttribute and FollowersAttribute2 are the collection keys denoting the
	// primary key for the followers relation (M2M).
	FollowersAttributes = []string{"user_id", "follower_id"}
	// FollowingAttribute and FollowingAttribute2 are the collection keys denoting the
	// primary key for the following relation (M2M).
	FollowingAttributes = []string{"user_id", "follower_id"}
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
