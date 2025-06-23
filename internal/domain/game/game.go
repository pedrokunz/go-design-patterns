package game

import (
	"encoding/json"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"github.com/pedrokunz/go-design-patterns/internal/domain/room"
)

type Game struct {
	Aggregate    *internal.Aggregate `json:"aggregate"`
	Player       *player.Player      `json:"player"`
	Rooms        []room.Room         `json:"rooms"`
	IsPlayerTurn bool                `json:"-"`
}

func (g *Game) Apply(event internal.Event) error {
	switch event.Type() {
	case Created:
		return json.Unmarshal(event.Payload(), g)
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

func LoadFromHistory(events []internal.Event) (*Game, error) {
	game := &Game{
		Aggregate: &internal.Aggregate{},
	}
	for _, event := range events {
		if err := game.Apply(event); err != nil {
			return nil, err
		}
	}
	return game, nil
}
