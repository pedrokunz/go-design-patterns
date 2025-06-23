package game

import (
	"encoding/json"
	"errors"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
)

func (g *Game) AddPlayer(player *player.Player) error {
	if player == nil {
		return errors.New("player cannot be nil")
	}

	if g == nil {
		return errors.New("game cannot be nil")
	}

	if g.Player != nil {
		return errors.New("game already has a player")
	}

	payload, err := json.Marshal(player)
	if err != nil {
		return err
	}

	event, err := internal.NewEventBuilder(
		g.Aggregate,
		payload,
		PlayerAdded,
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
