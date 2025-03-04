// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/dynamodb/o2m2types/ent"
)

func Example_O2M2Types() {
	client, err := ent.Open("dynamodb", "")
	if err != nil {
		log.Fatalf("failed opening connection to dynamodb: %v", err)
	}
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Output:
	// User created: User(id=1, age=30, name=a8m)
	// a8m
	// 2
}

func Do(ctx context.Context, client *ent.Client) error {
	// Create the 2 pets.
	pedro, err := client.Pet.
		Create().
		SetID(1).
		SetName("pedro").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %w", err)
	}
	lola, err := client.Pet.
		Create().
		SetID(2).
		SetName("lola").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %w", err)
	}
	// Create the user, and add its pets on the creation.
	a8m, err := client.User.
		Create().
		SetID(1).
		SetAge(30).
		SetName("a8m").
		AddPets(pedro, lola).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	fmt.Println("User created:", a8m)
	//Output: User(id=1, age=30, name=a8m)

	// Query the owner. Unlike `Only`, `OnlyX` panics if an error occurs.
	owner := pedro.QueryOwner().OnlyX(ctx)
	fmt.Println(owner.Name)
	// Output: a8m

	//// Traverse the sub-graph. Unlike `Count`, `CountX` panics if an error occurs.
	//count := pedro.
	//	QueryOwner(). // a8m
	//	QueryPets().  // pedro, lola
	//	CountX(ctx)   // count
	//fmt.Println(count)
	//// Output: 2
	return nil
}
