package player_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/internal"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer(t *testing.T) {
	t.Run("constructs a player", func(t *testing.T) {
		actual := player.New("Elmster")
		expected := &player.Player{
			Name:   "Elmster",
			Armour: internal.Armour{Value: 0},
			Attack: internal.Attack{Min: 1, Max: 100},
			Life:   internal.Life{Value: 100},
		}

		require.Equal(t, actual, expected, "actual %v, expected %v", actual, expected)
	})

	t.Run("takes damage", func(t *testing.T) {
		actual := player.New("Elmster")
		damage := actual.TakeDamage(internal.Attack{Min: 1, Max: 10})

		require.Greater(t, damage, 0)
		require.Less(t, damage, 10)
		require.GreaterOrEqual(t, actual.Life.Value, 90)
		require.Less(t, actual.Life.Value, 100)
	})
}
