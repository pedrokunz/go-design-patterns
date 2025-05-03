package internal

import (
	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/item"
)

type TreasureRoom struct {
	items []item.Item
}

func NewTreasureRoom(items []item.Item) *TreasureRoom {
	return &TreasureRoom{items: items}
}

func (t *TreasureRoom) Items() []item.Item {
	return t.items
}

func (t *TreasureRoom) Enemies() []*enemy.Enemy {
	return make([]*enemy.Enemy, 0)
}
