package game_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/aggregate/game"
	"github.com/stretchr/testify/require"
)

func TestNewState(t *testing.T) {
	t.Run("returns a valid state", func(t *testing.T) {
		actual := game.NewState()

		require.NotNil(t, actual, "NewState should not be nil")
	})

	t.Run("returns the same state", func(t *testing.T) {
		actual := game.NewState()
		expected := game.NewState()

		require.True(
			t,
			expected == actual,
			"NewState should return the same state",
		)
	})
}
