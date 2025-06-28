package game

import (
	"encoding/json"
	"errors"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"github.com/pedrokunz/go-design-patterns/internal/domain/room"
)

type Game struct {
	Aggregate    *aggregate.Aggregate `json:"aggregate"`
	Player       *player.Player       `json:"player"`
	Rooms        []room.Room          `json:"rooms"`
	IsPlayerTurn bool                 `json:"-"`
}

func (g *Game) Apply(event aggregate.Event) error {
	switch event.Type() {
	case Created:
		var createdGame Game
		if err := json.Unmarshal(event.Payload(), &createdGame); err != nil {
			return err
		}
		*g = createdGame
	case PlayerAdded:
		playerAdded := &player.Player{}
		if err := json.Unmarshal(event.Payload(), playerAdded); err != nil {
			return err
		}
		g.Player = playerAdded
	case RoomsCreated:
		var roomsCreated []room.Room
		if err := json.Unmarshal(event.Payload(), &roomsCreated); err != nil {
			return err
		}
		g.Rooms = roomsCreated
	}

	return nil
}

func LoadFromHistory(events []aggregate.Event) (*Game, error) {
	if len(events) == 0 {
		return nil, errors.New("invalid game state: no events found")
	}

	game := &Game{
		Aggregate: &aggregate.Aggregate{},
	}

	for _, event := range events {
		if err := game.Apply(event); err != nil {
			return nil, err
		}
		game.Aggregate.ApplyEvent(event)
	}
	return game, nil
}
