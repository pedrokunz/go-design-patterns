package player_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/core/internal"
	"github.com/pedrokunz/go-design-patterns/domain/core/player"
)

func TestPlayer(t *testing.T) {
	t.Run("constructs a player", func(t *testing.T) {
		got := player.New("Elmster")
		want := player.Player{
			Name:   "Elmster",
			Armour: internal.Armour{Value: 0},
			Attack: internal.Attack{Value: 1},
			Life:   internal.Life{Value: 100},
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
