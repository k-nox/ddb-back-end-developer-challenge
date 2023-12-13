package graph_test

import (
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
	}{
		{
			name:              "should correctly apply damage when character has no resistances",
			roll:              2,
			damageType:        model.DamageTypeThunder,
			expectedHitPoints: 23,
		},
		{
			name:              "should correctly apply damnage when character has resistance",
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
			client.MustPost(query, &resp)
			require.Equal(t, c.expectedHitPoints, resp.DamageCharacter.CurrentHitPoints)

			char, err := app.GetCharacterByID(1)
			require.NoError(t, err)
			require.Equal(t, c.expectedHitPoints, char.CurrentHitPoints)
		})
	}
}
