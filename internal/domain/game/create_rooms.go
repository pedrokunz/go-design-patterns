package game

import (
	"encoding/json"
	"github.com/pedrokunz/go-design-patterns/internal/domain/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
	"github.com/pedrokunz/go-design-patterns/internal/domain/item"
	"github.com/pedrokunz/go-design-patterns/internal/domain/room"
)

func (g *Game) CreateRooms() error {
	treasuryRoom := room.Room{
		Kind: room.KindTreasure,
		Items: []item.Item{
			{
				Name: "Sword",
				Type: item.Weapon,
			},
			{
				Name: "Shield",
				Type: item.Armour,
			},
		},
	}

	enemyRoom := room.Room{
		Kind: room.KindEnemy,
		Enemies: []*enemy.Enemy{
			enemy.New(enemy.Goblin),
		},
	}

	rooms := []room.Room{treasuryRoom, enemyRoom}
	payload, err := json.Marshal(rooms)
	if err != nil {
		return err
	}

	event, err := internal.NewEventBuilder(
		g.Aggregate,
		payload,
		RoomsCreated,
	).
		Build()
	if err != nil {
		return err
	}

	err = g.Apply(event)
	if err != nil {
		return err
	}

	g.Aggregate.ApplyEvent(event)

	return nil
}
