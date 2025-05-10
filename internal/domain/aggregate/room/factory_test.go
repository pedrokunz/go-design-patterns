package room_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/item"
	room2 "github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/room"
	internal2 "github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/room/internal"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoomFactory(t *testing.T) {
	t.Run("fails to construct when kind is invalid", func(t *testing.T) {
		items := []item.Item{}
		enemies := []enemy.Enemy{}
		kind := room2.Kind("invalid")

		actual := room2.Factory(room2.FactoryInput{
			Kind:    kind,
			Items:   items,
			Enemies: enemies,
		})

		require.Nil(t, actual)
	})

	t.Run("constructs a treasure room", func(t *testing.T) {
		items := []item.Item{}
		enemies := []enemy.Enemy{}
		kind := room2.KindTreasure

		actual := room2.Factory(room2.FactoryInput{
			Kind:    kind,
			Items:   items,
			Enemies: enemies,
		})

		expected := internal2.NewTreasureRoom(items)

		require.NotNil(t, actual)
		require.Equal(t, expected, actual)
	})

	t.Run("constructs an enemy room", func(t *testing.T) {
		items := []item.Item{}
		enemies := []enemy.Enemy{}
		kind := room2.KindEnemy

		actual := room2.Factory(room2.FactoryInput{
			Kind:    kind,
			Items:   items,
			Enemies: enemies,
		})

		expected := internal2.NewEnemyRoom(items, enemies)

		require.NotNil(t, actual)
		require.Equal(t, expected, actual)
	})
}
