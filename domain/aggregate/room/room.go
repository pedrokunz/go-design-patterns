package room

import (
	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/item"
)

type Room struct {
	Items   []item.Item
	Enemies []enemy.Enemy
	Kind    Kind
}

func New(items []item.Item, enemies []enemy.Enemy, kind Kind) *Room {
	return &Room{
		Items:   items,
		Enemies: enemies,
		Kind:    kind,
	}
}
