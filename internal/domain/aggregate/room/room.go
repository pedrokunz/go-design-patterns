package room

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/item"
)

type Room interface {
	Items() []item.Item
	Enemies() []*enemy.Enemy
}
