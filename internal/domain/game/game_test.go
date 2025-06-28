package game_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("returns a valid game", func(t *testing.T) {
		actual, err := game.Create()

		require.NoError(t, err, "Create should not return an error")
		require.NotNil(t, actual, "Create should not be nil")
		require.NotNil(t, actual.Rooms, "Rooms should not be nil")
		require.Len(t, actual.Rooms, 0, "Rooms should be empty")
	})

	t.Run("returns different games", func(t *testing.T) {
		actual, actualCreateErr := game.Create()
		expected, expectedCreateErr := game.Create()

		require.NoError(t, actualCreateErr, "Create should not return an error")
		require.NoError(t, expectedCreateErr, "Create should not return an error")

		require.NotEqual(
			t,
			actual,
			expected,
			"Games should not be equal",
		)
	})
}
