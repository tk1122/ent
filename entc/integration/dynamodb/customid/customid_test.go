// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package customid

import (
	"context"
	"log"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	"entgo.io/ent/entc/integration/dynamodb/customid/ent"
)

func TestDynamoDBCustomID(t *testing.T) {
	client, err := ent.Open("dynamodb", "")
	if err != nil {
		log.Fatalf("failed opening connection to dynamodb: %v", err)
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	client.User.Delete().ExecX(ctx)
	client.Blob.Delete().ExecX(ctx)
	client.Car.Delete().ExecX(ctx)
	client.Group.Delete().ExecX(ctx)
	client.Pet.Delete().ExecX(ctx)

	CustomID(t, client)
}

func CustomID(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	nat := client.User.Create().SetID(1).SaveX(ctx)
	require.Equal(t, 1, nat.ID)
	//_, err := client.User.Create().SetID(1).Save(ctx)
	//require.True(t, ent.IsConstraintError(err), "duplicate id")
	a8m := client.User.Create().SetID(5).SaveX(ctx)
	require.Equal(t, 5, a8m.ID)

	hub := client.Group.Create().SetID(3).AddUsers(a8m, nat).SaveX(ctx)
	require.Equal(t, 3, hub.ID)
	hubUsers, err := hub.QueryUsers().All(ctx)
	require.Nil(t, err)
	var ids []int
	for _, user := range hubUsers {
		ids = append(ids, user.ID)
	}
	require.Equal(t, []int{1, 5}, ids)

	blb := client.Blob.Create().SaveX(ctx)
	require.NotEmpty(t, blb.ID, "use default value")
	id := uuid.NewV4()
	chd := client.Blob.Create().SetID(id).SetParent(blb).SaveX(ctx)
	require.Equal(t, id, chd.ID, "use provided id")
	require.Equal(t, blb.ID, chd.QueryParent().OnlyX(ctx).ID)
	lnk := client.Blob.Create().SetID(uuid.NewV4()).AddLinks(chd, blb).SaveX(ctx)
	require.Equal(t, 2, lnk.QueryLinks().CountX(ctx))
	require.Equal(t, lnk.ID, chd.QueryLinks().OnlyX(ctx).ID)
	require.Equal(t, lnk.ID, blb.QueryLinks().OnlyX(ctx).ID)
	blobs, err := client.Blob.Query().All(ctx)
	require.Nil(t, err)
	var blobIds []uuid.UUID
	for _, blob := range blobs {
		blobIds = append(blobIds, blob.ID)
	}
	require.Len(t, blobIds, 3)

	//pedro := client.Pet.Create().SetID("pedro").SetOwner(a8m).SaveX(ctx)
	//require.Equal(t, a8m.ID, pedro.QueryOwner().OnlyIDX(ctx))
	//require.Equal(t, pedro.ID, a8m.QueryPets().OnlyIDX(ctx))
	//xabi := client.Pet.Create().SetID("xabi").AddFriends(pedro).SetBestFriend(pedro).SaveX(ctx)
	//require.Equal(t, "xabi", xabi.ID)
	//pedro = client.Pet.Query().Where(pet.HasOwnerWith(user.ID(a8m.ID))).OnlyX(ctx)
	//require.Equal(t, "pedro", pedro.ID)
	//
	//pets := client.Pet.Query().WithFriends().WithBestFriend().Order().AllX(ctx)
	//require.Len(t, pets, 2)
	//
	//require.Equal(t, pedro.ID, pets[0].ID)
	//require.NotNil(t, pets[0].Edges.BestFriend)
	//require.Equal(t, xabi.ID, pets[0].Edges.BestFriend.ID)
	//require.Len(t, pets[0].Edges.Friends, 1)
	//require.Equal(t, xabi.ID, pets[0].Edges.Friends[0].ID)
	//
	//require.Equal(t, xabi.ID, pets[1].ID)
	//require.NotNil(t, pets[1].Edges.BestFriend)
	//require.Equal(t, pedro.ID, pets[1].Edges.BestFriend.ID)
	//require.Len(t, pets[1].Edges.Friends, 1)
	//require.Equal(t, pedro.ID, pets[1].Edges.Friends[0].ID)
	//
	//bee := client.Car.Create().SetID(1).SetModel("Chevrolet Camaro").SetOwner(pedro).SaveX(ctx)
	//require.NotNil(t, bee)
	//bee = client.Car.Query().WithOwner().OnlyX(ctx)
	//require.Equal(t, "Chevrolet Camaro", bee.Model)
	//require.NotNil(t, bee.Edges.Owner)
	//require.Equal(t, pedro.ID, bee.Edges.Owner.ID)
}
