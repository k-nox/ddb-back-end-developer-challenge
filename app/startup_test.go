package app_test

import (
	"github.com/k-nox/ddb-backend-developer-challenge/app"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApp_Startup(t *testing.T) {
	t.Run("should correctly parse startup character", func(t *testing.T) {
		a, err := app.New("file::memory:?cache=shared", "../db/migrations")
		defer a.CloseDB()
		require.NoError(t, err)

		startupErr := a.Startup("../briv.json")
		require.NoError(t, startupErr)

		char, err := a.GetCharacterByName("Briv")
		require.NoError(t, err)
		require.NotNil(t, char)
		require.NotZero(t, char.ID)
		require.Equal(t, "Briv", char.Name)
	})
}
