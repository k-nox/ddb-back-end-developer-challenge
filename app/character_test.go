package app

import (
	"github.com/k-nox/ddb-backend-developer-challenge/graph/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApp_InsertCharacter(t *testing.T) {
	cases := []struct {
		name        string
		in          *model.Character
		expectedErr error
	}{
		{
			name: "should correctly insert valid character",
			in: &model.Character{
				Name:             "Briv",
				MaxHitPoints:     23,
				CurrentHitPoints: 24,
				Level:            5,
				Stats: &model.Stats{
					Strength:     10,
					Dexterity:    11,
					Constitution: 12,
					Intelligence: 13,
					Wisdom:       14,
					Charisma:     15,
				},
			},
		},
		{
			name:        "should reject if character is nil",
			expectedErr: InvalidCharError,
		},
		{
			name:        "should refect if stats are nil",
			in:          &model.Character{},
			expectedErr: InvalidCharError,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a, err := New("file::memory:?cache=shared", "../db/migrations")
			defer a.CloseDB()
			require.NoError(t, err)

			sess := a.db.NewSession(nil)
			// ensure there are no records already in the db
			count, countErr := sess.Select("count(*)").From(characterTable).ReturnInt64()
			require.NoError(t, countErr)
			require.Zero(t, count)

			// now insert the character
			insertErr := a.InsertCharacter(c.in)

			// select all rows in table now
			var records []characterRecord
			n, selectErr := sess.Select("*").From(characterTable).Load(&records)
			require.NoError(t, selectErr)

			if c.expectedErr != nil {
				require.ErrorIs(t, c.expectedErr, insertErr)
				require.Zero(t, n)
				require.Empty(t, records)
			} else {
				require.NoError(t, insertErr)
				// should only have inserted 1 row
				require.EqualValues(t, 1, n)
				require.NotEmpty(t, records)
				require.NotZero(t, records[0].Id)
				records[0].Id = 0
				require.Equal(t, c.in, records[0].toModel())
			}
		})
	}
}

func TestApp_GetCharacterByName(t *testing.T) {
	cases := []struct {
		name        string
		charName    string
		toInsert    *model.Character
		expectedErr error
	}{
		{
			name:     "should return correct character if character is in db",
			charName: "Briv",
			toInsert: &model.Character{
				Name:             "Briv",
				MaxHitPoints:     23,
				CurrentHitPoints: 24,
				Level:            5,
				Stats: &model.Stats{
					Strength:     10,
					Dexterity:    11,
					Constitution: 12,
					Intelligence: 13,
					Wisdom:       14,
					Charisma:     15,
				},
			},
		},
		{
			name:        "should return CharNotFoundError if no character is found by the requested name",
			charName:    "Max",
			expectedErr: CharNotFoundError,
			toInsert: &model.Character{
				Name:             "Briv",
				MaxHitPoints:     23,
				CurrentHitPoints: 24,
				Level:            5,
				Stats: &model.Stats{
					Strength:     10,
					Dexterity:    11,
					Constitution: 12,
					Intelligence: 13,
					Wisdom:       14,
					Charisma:     15,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a, err := New("file::memory:?cache=shared", "../db/migrations")
			defer a.CloseDB()
			require.NoError(t, err)

			sess := a.db.NewSession(nil)
			// ensure there are no records already in the db
			count, countErr := sess.Select("count(*)").From(characterTable).ReturnInt64()
			require.NoError(t, countErr)
			require.Zero(t, count)

			// setup
			insertErr := a.InsertCharacter(c.toInsert)
			require.NoError(t, insertErr)

			out, outErr := a.GetCharacterByName(c.charName)
			if c.expectedErr != nil {
				require.ErrorIs(t, c.expectedErr, outErr)
				require.Nil(t, out)
			} else {
				require.Equal(t, 1, out.ID)
				c.toInsert.ID = 1
				require.Equal(t, c.toInsert, out)
			}
		})
	}
}
