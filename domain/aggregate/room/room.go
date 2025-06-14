package room

import (
	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/item"
)

type Room interface {
	Items() []item.Item
	Enemies() []*enemy.Enemy
}
