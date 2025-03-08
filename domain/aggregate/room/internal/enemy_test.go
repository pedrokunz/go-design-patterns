package internal_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room/internal"
	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/item"
	"github.com/stretchr/testify/require"
)

func TestNewEnemyRoom(t *testing.T) {
	t.Run("constructs an enemy room", func(t *testing.T) {
		items := []item.Item{}
		enemies := []enemy.Enemy{}
		actual := internal.NewEnemyRoom(items, enemies)
		expected := internal.NewEnemyRoom(items, enemies)

		require.NotNil(t, actual)
		require.NotNil(t, expected)
		require.Equal(t, expected, actual)
		require.Equal(t, expected.Items(), actual.Items())
		require.Equal(t, expected.Enemies(), actual.Enemies())
	})
}
