package room_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/item"
	"github.com/pedrokunz/go-design-patterns/internal/domain/room"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoomFactory(t *testing.T) {
	t.Run("fails to construct when kind is invalid", func(t *testing.T) {
		items := make([]item.Item, 0)
		enemies := make([]*enemy.Enemy, 0)
		kind := room.Kind("invalid")

		actual := room.Factory(room.FactoryInput{
			Kind:    kind,
			Items:   items,
			Enemies: enemies,
		})

		require.Nil(t, actual)
	})

	t.Run("constructs a treasure room", func(t *testing.T) {
		items := make([]item.Item, 0)
		enemies := make([]*enemy.Enemy, 0)
		kind := room.KindTreasure

		actual := room.Factory(room.FactoryInput{
			Kind:    kind,
			Items:   items,
			Enemies: enemies,
		})

		expected := room.Room{Kind: room.KindTreasure, Items: items}

		require.NotNil(t, actual)
		require.Equal(t, expected, actual)
	})

	t.Run("constructs an enemy room", func(t *testing.T) {
		items := make([]item.Item, 0)
		enemies := make([]*enemy.Enemy, 0)
		kind := room.KindEnemy

		actual := room.Factory(room.FactoryInput{
			Kind:    kind,
			Items:   items,
			Enemies: enemies,
		})

		expected := room.Room{Kind: room.KindEnemy, Items: items, Enemies: enemies}

		require.NotNil(t, actual)
		require.Equal(t, expected, actual)
	})
}
