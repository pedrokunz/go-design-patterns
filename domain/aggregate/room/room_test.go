package room_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room"
	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/item"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("constructs a room", func(t *testing.T) {
		items := []item.Item{}
		enemies := []enemy.Enemy{}
		kind := room.KindTreasure

		actual := room.New(items, enemies, kind)
		expected := &room.Room{
			Items:   []item.Item{},
			Enemies: []enemy.Enemy{},
			Kind:    room.KindTreasure,
		}

		require.Equal(t, expected, actual)
	})
}
