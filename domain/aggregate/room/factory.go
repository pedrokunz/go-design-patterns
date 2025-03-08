package room

import (
	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room/internal"
	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/item"
)

type FactoryInput struct {
	Kind    Kind
	Items   []item.Item
	Enemies []enemy.Enemy
}

func Factory(input FactoryInput) Room {
	switch input.Kind {
	case KindTreasure:
		return internal.NewTreasureRoom(input.Items)
	case KindEnemy:
		return internal.NewEnemyRoom(input.Items, input.Enemies)
	default:
		return nil
	}
}
