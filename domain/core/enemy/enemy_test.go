package enemy_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/internal"
)

func TestEnemy(t *testing.T) {
	t.Run("constructs an enemy", func(t *testing.T) {
		actual := enemy.New(enemy.Goblin)
		expected := &enemy.Enemy{
			Type:   enemy.Goblin,
			Armour: internal.Armour{Value: 0},
			Attack: internal.Attack{Value: 1},
			Life:   internal.Life{Value: 100},
		}

		require.Equal(t, actual, expected, "actual %v, expected %v", actual, expected)
	})
}
