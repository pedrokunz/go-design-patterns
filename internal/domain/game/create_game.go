package game

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
	"github.com/pedrokunz/go-design-patterns/internal/domain/room"
)

func Create() (*Game, error) {
	domainAggregate, newDomainAggregateErr := internal.NewAggregate(
		uuid.New(),
		Aggregate,
	)
	if newDomainAggregateErr != nil {
		return nil, newDomainAggregateErr
	}

	game := &Game{
		Aggregate: domainAggregate,
		Rooms:     make([]room.Room, 0),
	}

	payload, err := json.Marshal(game)
	if err != nil {
		return nil, err
	}

	event, err := internal.NewEventBuilder(
		game.Aggregate,
		payload,
		Created,
	).
		Build()
	if err != nil {
		return nil, err
	}

	err = game.Apply(event)
	if err != nil {
		return nil, err
	}

	game.Aggregate.ApplyEvent(event)

	return game, nil
}
