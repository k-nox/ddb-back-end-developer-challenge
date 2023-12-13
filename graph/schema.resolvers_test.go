package graph_test

import (
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/k-nox/ddb-backend-developer-challenge/app"
	"github.com/k-nox/ddb-backend-developer-challenge/graph"
	"github.com/k-nox/ddb-backend-developer-challenge/graph/generated"
	"github.com/k-nox/ddb-backend-developer-challenge/graph/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMutationResolver_DamageCharacter(t *testing.T) {
	cases := []struct {
		name              string
		roll              int
		damageType        model.DamageType
		expectedHitPoints int
		expectedErr       error
	}{
		{
			name:              "should reject if roll is less than zero",
			roll:              -1,
			damageType:        model.DamageTypeThunder,
			expectedHitPoints: 25,
			expectedErr:       errors.New("[{\"message\":\"roll -1 is invalid; must be positive value\",\"path\":[\"damageCharacter\"]}]"),
		},
		{
			name:              "should correctly apply damage when character has no resistances",
			roll:              2,
			damageType:        model.DamageTypeThunder,
			expectedHitPoints: 23,
		},
		{
			name:              "should correctly apply damage when character has resistance",
			roll:              2,
			damageType:        model.DamageTypeSLAShing,
			expectedHitPoints: 24,
		},
		{
			name:              "should correctly apply no damage when character has immunity",
			roll:              2,
			damageType:        model.DamageTypeFire,
			expectedHitPoints: 25,
		},
		{
			name:              "should correctly set hit points to zero when character suffers damage greater than their current hit points",
			roll:              30,
			damageType:        model.DamageTypeBludgeoning,
			expectedHitPoints: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// startup in-memory db and populate with starting character
			app, err := app.New("file::memory:?cache=shared", "../db/migrations")
			defer app.CloseDB()
			require.NoError(t, err)
			err = app.Startup("../briv.json")
			require.NoError(t, err)

			client := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.New(app)})))

			var resp struct {
				DamageCharacter struct{ CurrentHitPoints int }
			}
			query := fmt.Sprintf("mutation { damageCharacter(input: { characterId: 1, damageType: %s, roll: %d }) { currentHitPoints } }", c.damageType.String(), c.roll)
			err = client.Post(query, &resp)
			if c.expectedErr != nil {
				require.Equal(t, c.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, c.expectedHitPoints, resp.DamageCharacter.CurrentHitPoints)
			}

			char, err := app.GetCharacterByID(1)
			require.NoError(t, err)
			require.Equal(t, c.expectedHitPoints, char.CurrentHitPoints)
		})
	}
}

func TestMutationResolver_HealCharacter(t *testing.T) {
	cases := []struct {
		name              string
		roll              int
		expectedHitPoints int
		expectedErr       error
	}{
		{
			name:              "should correctly apply healing when current hit points + roll is less than character's max hit points",
			roll:              2,
			expectedHitPoints: 12,
		},
		{
			name:              "should only heal to character's max hit points if current hit points + roll is greater than max hit points",
			roll:              20,
			expectedHitPoints: 25,
		},
		{
			name:              "should reject negative values for healing rolls",
			roll:              -1,
			expectedHitPoints: 10,
			expectedErr:       errors.New("[{\"message\":\"roll -1 is invalid; must be positive value\",\"path\":[\"healCharacter\"]}]"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			app, err := app.New("file::memory:?cache=shared", "../db/migrations")
			defer app.CloseDB()
			require.NoError(t, err)
			err = app.Startup("../briv.json")
			require.NoError(t, err)

			// set current hit points to 10
			err = app.UpdateHitPoints(1, 10)
			require.NoError(t, err)

			client := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.New(app)})))

			var resp struct {
				HealCharacter struct{ CurrentHitPoints int }
			}
			query := fmt.Sprintf("mutation { healCharacter(input: { characterId: 1, roll: %d }) { currentHitPoints } }", c.roll)

			err = client.Post(query, &resp)
			if c.expectedErr != nil {
				require.Equal(t, c.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, c.expectedHitPoints, resp.HealCharacter.CurrentHitPoints)
			}

			char, err := app.GetCharacterByID(1)
			require.NoError(t, err)
			require.Equal(t, c.expectedHitPoints, char.CurrentHitPoints)
		})
	}
}
