// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/dynamodb/o2mrecur/ent"
)

func Example_O2MRecur() {
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
	// Tree leafs [1 3 5]
	// [1 3 5]
	// Node(id=1, value=2)
}

func Do(ctx context.Context, client *ent.Client) error {
	root, err := client.Node.
		Create().
		SetID(2).
		SetValue(2).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating the root: %w", err)
	}

	// Add additional nodes to the tree:
	//
	//       2
	//     /   \
	//    1     4
	//        /   \
	//       3     5
	//

	// Unlike `Save`, `SaveX` panics if an error occurs.
	n1 := client.Node.
		Create().
		SetID(1).
		SetValue(1).
		SetParent(root).
		SaveX(ctx)
	n4 := client.Node.
		Create().
		SetID(4).
		SetValue(4).
		SetParent(root).
		SaveX(ctx)
	n3 := client.Node.
		Create().
		SetID(3).
		SetValue(3).
		SetParent(n4).
		SaveX(ctx)
	n5 := client.Node.
		Create().
		SetID(5).
		SetValue(5).
		SetParent(n4).
		SaveX(ctx)

	fmt.Println("Tree leafs", []int{n1.Value, n3.Value, n5.Value})
	////Output: Tree leafs [1 3 5]
	////
	////Get all leafs (nodes without children).
	////Unlike `Int`, `IntX` panics if an error occurs.
	//ints := client.Node.
	//	Query().                             // All nodes.
	//	Where(node.Not(node.HasChildren())). // Only leafs.
	//	Order(ent.Asc()).                    // Order by their `value` field.
	//	AllX(ctx)
	//fmt.Println(ints)
	////Output: [1 3 5]
	//
	//// Get orphan nodes (nodes without parent).
	//// Unlike `Only`, `OnlyX` panics if an error occurs.
	//orphan := client.Node.
	//	Query().
	//	Where(node.Not(node.HasParent())).
	//	OnlyX(ctx)
	//fmt.Println(orphan)
	//// Output: Node(id=1, value=2)

	return nil
}
