// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/dynamodb/m2mbidi/ent"
)

func Example_M2MBidi() {
	client, err := ent.Open("dynamodb", "")
	if err != nil {
		log.Fatalf("failed opening connection to dynamodb: %v", err)
	}
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Output:
	// [User(id=1, age=30, name=a8m)]
	// [User(id=2, age=28, name=nati)]
	// [User(id=1, age=30, name=a8m) User(id=2, age=28, name=nati)]
}

func Do(ctx context.Context, client *ent.Client) error {
	// Unlike `Save`, `SaveX` panics if an error occurs.
	a8m := client.User.
		Create().
		SetID(1).
		SetAge(30).
		SetName("a8m").
		SaveX(ctx)
	nati := client.User.
		Create().
		SetID(2).
		SetAge(28).
		SetName("nati").
		AddFriends(a8m).
		SaveX(ctx)

	// Query friends. Unlike `All`, `AllX` panics if an error occurs.
	friends := nati.
		QueryFriends().
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=1, age=30, name=a8m)]

	friends = a8m.
		QueryFriends().
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=2, age=28, name=nati)]

	//// Query the graph:
	//friends = client.User.
	//	Query().
	//	Where(user.HasFriends()).
	//	AllX(ctx)
	//fmt.Println(friends)
	//// Output: [User(id=1, age=30, name=a8m) User(id=2, age=28, name=nati)]
	return nil
}
