// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package hooks

import (
	"context"
	"log"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"entgo.io/ent/entc/integration/dynamodb/hooks/ent"
	"entgo.io/ent/entc/integration/dynamodb/hooks/ent/card"
	"entgo.io/ent/entc/integration/dynamodb/hooks/ent/hook"
	_ "entgo.io/ent/entc/integration/dynamodb/hooks/ent/runtime"
	"entgo.io/ent/entc/integration/dynamodb/hooks/ent/user"
)

func TestSchemaHooks(t *testing.T) {
	ctx := context.Background()
	client := initDynamoDBClient(ctx)

	_, err := client.Card.Create().SetID(1).SetNumber("123").Save(ctx)
	require.EqualError(t, err, "card number is too short", "error is returned from hook")
	crd := client.Card.Create().SetID(2).SetNumber("1234").SaveX(ctx)
	require.Equal(t, "unknown", crd.Name, "name was set by hook")
	client.Use(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			name, ok := m.Name()
			require.True(t, !ok && name == "", "should be the first hook to execute")
			return next.Mutate(ctx, m)
		})
	})
	client.Card.Create().SetID(2).SetNumber("1234").SaveX(ctx)
	_, err = client.Card.Update().Save(ctx)
	require.EqualError(t, err, "OpUpdate operation is not allowed")
}

func TestRuntimeHooks(t *testing.T) {
	ctx := context.Background()
	client := initDynamoDBClient(ctx)

	var calls int
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			calls++
			return next.Mutate(ctx, m)
		})
	})
	client.Card.Create().SetID(3).SetNumber("1234").SaveX(ctx)
	client.User.Create().SetID(1).SetName("a8m").SaveX(ctx)
	require.Equal(t, 2, calls)
	client.Card.Create().SetID(4).SetNumber("1234").SaveX(ctx)
	client.User.Create().SetID(2).SetName("a8m").SaveX(ctx)
	require.Equal(t, 4, calls, "debug client should keep thr same hooks")
}

func TestRuntimeChain(t *testing.T) {
	ctx := context.Background()
	client := initDynamoDBClient(ctx)
	var (
		chain  hook.Chain
		values []int
	)
	for value := 0; value < 5; value++ {
		chain = chain.Append(func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				values = append(values, value)
				return next.Mutate(ctx, m)
			})
		})
	}
	client.User.Use(chain.Hook())
	client.User.Create().SetID(3).SetName("alexsn").SaveX(ctx)
	require.Len(t, values, 5)
	require.True(t, sort.IntsAreSorted(values))
}

func TestMutationClient(t *testing.T) {
	ctx := context.Background()
	client := initDynamoDBClient(ctx)

	client.Card.Use(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			id, _ := m.OwnerID()
			usr := m.Client().User.GetX(ctx, id)
			m.SetName(usr.Name)
			return next.Mutate(ctx, m)
		})
	})
	a8m := client.User.Create().SetID(1).SetName("a8m").SaveX(ctx)
	crd := client.Card.Create().SetID(1).SetNumber("1234").SetOwner(a8m).SaveX(ctx)
	require.Equal(t, a8m.Name, crd.Name)
}

//func TestMutationTx(t *testing.T) {
//	ctx := context.Background()
//	client := initDynamoDBClient(ctx)
//
//	client.Card.Use(func(next ent.Mutator) ent.Mutator {
//		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
//			tx, err := m.Tx()
//			if err != nil {
//				return nil, err
//			}
//			if err := tx.Rollback(); err != nil {
//				return nil, err
//			}
//			return nil, fmt.Errorf("rolled back")
//		})
//	})
//	tx, err := client.Tx(ctx)
//	require.NoError(t, err)
//	a8m := tx.User.Create().SetName("a8m").SaveX(ctx)
//	crd, err := tx.Card.Create().SetNumber("1234").SetOwner(a8m).Save(ctx)
//	require.EqualError(t, err, "rolled back")
//	require.Nil(t, crd)
//	_, err = tx.Card.Query().All(ctx)
//	require.Error(t, err, "tx already rolled back")
//}

//func TestDeletion(t *testing.T) {
//	ctx := context.Background()
//	client := initDynamoDBClient(ctx)
//
//	client.User.Use(func(next ent.Mutator) ent.Mutator {
//		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
//			if !m.Op().Is(ent.OpDeleteOne) {
//				return next.Mutate(ctx, m)
//			}
//			id, ok := m.ID()
//			if !ok {
//				return nil, fmt.Errorf("missing id")
//			}
//			owner, err := m.Client().User.Get(ctx, id)
//			require.Nil(t, err)
//			cards, err := owner.QueryCards().All(ctx)
//			require.Nil(t, err)
//			require.NotEmpty(t, cards)
//			m.Client().Card.Delete().Where(card.IDEQ(cards[0].ID)).ExecX(ctx)
//			return next.Mutate(ctx, m)
//		})
//	})
//	a8m := client.User.Create().SetID(1).SetName("a8m").SaveX(ctx)
//	for i := 0; i < 5; i++ {
//		client.Card.Create().SetID(i + 1).SetNumber(fmt.Sprintf("card-%d", i)).SetOwner(a8m).SaveX(ctx)
//	}
//	client.User.DeleteOne(a8m).ExecX(ctx)
//	require.Zero(t, client.User.Query().CountX(ctx))
//	require.Zero(t, client.Card.Query().CountX(ctx))
//}

func TestPostCreation(t *testing.T) {
	ctx := context.Background()
	client := initDynamoDBClient(ctx)

	client.Card.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			//id, exists := m.ID()
			//require.False(t, exists, "id should not exist pre mutation")
			//require.Zero(t, id)
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			id, exists := m.ID()
			require.True(t, exists, "id should exist post mutation")
			require.NotZero(t, id)
			require.True(t, id == value.(*ent.Card).ID)
			return value, nil
		})
	}, ent.OpCreate))
	client.Card.Create().SetID(1).SetNumber("12345").SetName("a8m").SaveX(ctx)
	//client.Card.CreateBulk(client.Card.Create().SetNumber("12345")).SaveX(ctx)
}

func TestOldValues(t *testing.T) {
	ctx := context.Background()
	client := initDynamoDBClient(ctx)

	// Querying old fields post mutation should fail.
	client.Card.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			value, err := next.Mutate(ctx, m)
			require.NoError(t, err)
			_, err = m.OldNumber(ctx)
			require.Error(t, err)
			return value, nil
		})
	}, ent.OpUpdateOne))
	crd := client.Card.Create().SetID(1).SetNumber("1234").SetName("a8m").SaveX(ctx)
	client.Card.UpdateOneID(crd.ID).SetName("a8m").SaveX(ctx)

	// A typed hook.
	client.User.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			name, err := m.OldName(ctx)
			if err != nil {
				return nil, err
			}
			require.Equal(t, "a8m", name)
			return next.Mutate(ctx, m)
		})
	}, ent.OpUpdateOne))
	// A generic hook (executed on all types).
	client.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			namer, ok := m.(interface {
				OldName(context.Context) (string, error)
			})
			if !ok {
				// Skip if the mutation does not have
				// a method for getting the old name.
				return next.Mutate(ctx, m)
			}
			name, err := namer.OldName(ctx)
			if err != nil {
				return nil, err
			}
			require.Equal(t, "a8m", name)
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			_, err = namer.OldName(ctx)
			require.NoError(t, err)
			return value, nil
		})
	}, ent.OpUpdateOne))
	// A generic hook (executed on all types).
	client.Use(hook.Unless(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			namer, ok := m.(interface {
				OldName(context.Context) (string, error)
			})
			if !ok {
				// Skip if the mutation does not have
				// a method for getting the old name.
				return next.Mutate(ctx, m)
			}
			name, err := namer.OldName(ctx)
			if err != nil {
				return nil, err
			}
			require.Equal(t, "a8m", name)
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			_, err = namer.OldName(ctx)
			require.NoError(t, err)
			return value, nil
		})
	}, ^ent.OpUpdateOne))
	a8m := client.User.Create().SetID(1).SetName("a8m").SaveX(ctx)
	require.Equal(t, "a8m", a8m.Name)
	_, err := client.User.UpdateOne(a8m).SetName("Ariel").SetVersion(a8m.Version).Save(ctx)
	require.EqualError(t, err, "version field must be incremented by 1")
	a8m = client.User.UpdateOne(a8m).SetName("Ariel").SetVersion(a8m.Version + 1).SaveX(ctx)
	require.Equal(t, "Ariel", a8m.Name)
}

func TestConditions(t *testing.T) {
	ctx := context.Background()
	client := initDynamoDBClient(ctx)

	var calls int
	defer func() { require.Equal(t, 2, calls) }()
	client.Card.Use(hook.If(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			require.True(t, m.Op().Is(ent.OpUpdateOne))
			calls++
			return next.Mutate(ctx, m)
		})
	}, hook.Or(
		hook.HasFields(card.FieldName),
		hook.HasClearedFields(card.FieldName),
	)))
	client.User.Use(hook.If(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			require.True(t, m.Op().Is(ent.OpUpdate))
			incr, exists := m.AddedWorth()
			require.True(t, exists)
			require.EqualValues(t, 100, incr)
			return next.Mutate(ctx, m)
		})
	}, hook.HasAddedFields(user.FieldWorth)))

	crd := client.Card.Create().SetID(2).SetNumber("9876").SaveX(ctx)
	crd = crd.Update().SetName("alexsn").SaveX(ctx)
	crd = crd.Update().ClearName().SaveX(ctx)
	client.Card.DeleteOne(crd).ExecX(ctx)

	alexsn := client.User.Create().SetID(2).SetName("alexsn").SaveX(ctx)
	client.User.Update().Where(user.ID(alexsn.ID)).AddWorth(100).SaveX(ctx)
	client.User.DeleteOne(alexsn).ExecX(ctx)
}

func initDynamoDBClient(ctx context.Context) *ent.Client {
	client, err := ent.Open("dynamodb", "")
	if err != nil {
		log.Fatalf("failed opening connection to dynamodb: %v", err)
	}
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
