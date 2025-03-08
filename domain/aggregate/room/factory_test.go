package room_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room"
	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room/internal"
	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/item"
	"github.com/stretchr/testify/require"
)

func TestRoomFactory(t *testing.T) {
	t.Run("fails to construct when kind is invalid", func(t *testing.T) {
		items := []item.Item{}
		enemies := []enemy.Enemy{}
		kind := room.Kind("invalid")

		actual := room.Factory(room.FactoryInput{
			Kind:    kind,
			Items:   items,
			Enemies: enemies,
		})

		require.Nil(t, actual)
	})

	t.Run("constructs a treasure room", func(t *testing.T) {
		items := []item.Item{}
		enemies := []enemy.Enemy{}
		kind := room.KindTreasure

		actual := room.Factory(room.FactoryInput{
			Kind:    kind,
			Items:   items,
			Enemies: enemies,
		})

		expected := internal.NewTreasureRoom(items)

		require.NotNil(t, actual)
		require.Equal(t, expected, actual)
	})

	t.Run("constructs an enemy room", func(t *testing.T) {
		items := []item.Item{}
		enemies := []enemy.Enemy{}
		kind := room.KindEnemy

		actual := room.Factory(room.FactoryInput{
			Kind:    kind,
			Items:   items,
			Enemies: enemies,
		})

		expected := internal.NewEnemyRoom(items, enemies)

		require.NotNil(t, actual)
		require.Equal(t, expected, actual)
	})
}
