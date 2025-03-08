package internal_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room/internal"
	"github.com/pedrokunz/go-design-patterns/domain/core/item"
	"github.com/stretchr/testify/require"
)

func TestNewTreasureRoom(t *testing.T) {
	t.Run("constructs a treasure room", func(t *testing.T) {
		items := []item.Item{}
		actual := internal.NewTreasureRoom(items)
		expected := internal.NewTreasureRoom(items)

		require.NotNil(t, actual)
		require.NotNil(t, expected)
		require.Equal(t, expected, actual)
		require.Equal(t, expected.Items(), actual.Items())
		require.Equal(t, expected.Enemies(), actual.Enemies())
	})
}
