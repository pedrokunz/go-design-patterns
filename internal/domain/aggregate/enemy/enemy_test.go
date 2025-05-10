package enemy_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/internal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEnemy(t *testing.T) {
	t.Run("constructs an enemy", func(t *testing.T) {
		actual := enemy.New(enemy.Goblin)
		expected := &enemy.Enemy{
			Type:   enemy.Goblin,
			Armour: internal.Armour{Value: 0},
			Attack: internal.Attack{Min: 1, Max: 100},
			Life:   internal.Life{Value: 100},
		}

		require.Equal(t, actual, expected, "actual %v, expected %v", actual, expected)
	})
}
