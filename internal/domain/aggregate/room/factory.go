package room

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/item"
	internal2 "github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/room/internal"
)

type FactoryInput struct {
	Kind    Kind
	Items   []item.Item
	Enemies []*enemy.Enemy
}

func Factory(input FactoryInput) Room {
	switch input.Kind {
	case KindTreasure:
		return internal2.NewTreasureRoom(input.Items)
	case KindEnemy:
		return internal2.NewEnemyRoom(input.Items, input.Enemies)
	default:
		return nil
	}
}
