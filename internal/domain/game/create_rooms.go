package game

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate"
	"github.com/pedrokunz/go-design-patterns/internal/domain/room"
)

func (g *Game) CreateRooms(generator common.UUIDGenerator, configs []room.Config) error {
	if generator == nil {
		generator = uuid.NewRandom
	}

	var rooms []room.Room
	for _, config := range configs {
		roomID, err := generator()
		if err != nil {
			return err
		}
		rooms = append(rooms, room.Room{
			ID:      roomID,
			Kind:    config.Kind,
			Items:   config.Items,
			Enemies: config.Enemies,
		})
	}

	payload, err := json.Marshal(rooms)
	if err != nil {
		return err
	}

	event, err := aggregate.NewEventBuilder(
		g.Aggregate,
		payload,
		RoomsCreated,
	).Build()
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
