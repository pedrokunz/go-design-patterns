package player_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/core/internal"
	"github.com/pedrokunz/go-design-patterns/domain/core/player"
)

func TestPlayer(t *testing.T) {
	t.Run("constructs a player", func(t *testing.T) {
		actual := player.New("Elmster")
		expected := &player.Player{
			Name:   "Elmster",
			Armour: internal.Armour{Value: 0},
			Attack: internal.Attack{Value: 1},
			Life:   internal.Life{Value: 100},
		}

		require.Equal(t, actual, expected, "actual %v, expected %v", actual, expected)
	})
}
