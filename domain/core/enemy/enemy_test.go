package enemy_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/internal"
)

func TestEnemy(t *testing.T) {
	t.Run("constructs an enemy", func(t *testing.T) {
		got := enemy.New(enemy.Goblin)
		want := enemy.Enemy{
			Type:   enemy.Goblin,
			Armour: internal.Armour{Value: 0},
			Attack: internal.Attack{Value: 1},
			Life:   internal.Life{Value: 100},
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
