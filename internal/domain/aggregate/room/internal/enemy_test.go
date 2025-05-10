package internal_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/item"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/room/internal"
	"testing"

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
