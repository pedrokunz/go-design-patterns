package player_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer(t *testing.T) {
	t.Run("constructs a player", func(t *testing.T) {
		actual, createErr := player.Create("Player 1")
		expected := &player.Player{
			Name:   "Player 1",
			Armour: internal.Armour{Value: 0},
			Attack: internal.Attack{Min: 1, Max: 100},
			Life:   internal.Life{Value: 100},
		}

		require.NoError(t, createErr, "unexpected error creating player: %v", createErr)
		require.Equal(t, expected.Name, actual.Name, "actual %v, expected %v", actual.Name, expected.Name)
		require.Equal(t, expected.Armour, actual.Armour, "actual %v, expected %v", actual.Armour, expected.Armour)
		require.Equal(t, expected.Attack, actual.Attack, "actual %v, expected %v", actual.Attack, expected.Attack)
		require.Equal(t, expected.Life, actual.Life, "actual %v, expected %v", actual.Life, expected.Life)
		require.NotEmpty(t, actual.Aggregate.ID, "aggregate ID should not be empty")
		require.Equal(t, player.Aggregate, actual.Aggregate.Type, "aggregate type should be player")
	})

	t.Run("takes damage", func(t *testing.T) {
		actual, _ := player.Create("Player 1")
		damage := actual.TakeDamage(internal.Attack{Min: 1, Max: 10})

		require.Greater(t, damage, 0)
		require.Less(t, damage, 10)
		require.GreaterOrEqual(t, actual.Life.Value, 90)
		require.Less(t, actual.Life.Value, 100)
	})
}
