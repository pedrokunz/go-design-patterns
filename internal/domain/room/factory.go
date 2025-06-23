package room

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/item"
)

type FactoryInput struct {
	Kind    Kind
	Items   []item.Item
	Enemies []*enemy.Enemy
}

func Factory(input FactoryInput) Room {
	room := Room{
		Kind: input.Kind,
	}

	switch input.Kind {
	case KindTreasure:

		room.Items = input.Items
	case KindEnemy:
		room.Items = input.Items
		room.Enemies = input.Enemies
	}

	return room
}
