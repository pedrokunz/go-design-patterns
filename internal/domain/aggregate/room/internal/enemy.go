package internal

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/item"
)

type EnemyRoom struct {
	items   []item.Item
	enemies []*enemy.Enemy
}

func NewEnemyRoom(items []item.Item, enemies []*enemy.Enemy) *EnemyRoom {
	return &EnemyRoom{items: items, enemies: enemies}
}

func (e *EnemyRoom) Items() []item.Item {
	return e.items
}

func (e *EnemyRoom) Enemies() []*enemy.Enemy {
	return e.enemies
}
