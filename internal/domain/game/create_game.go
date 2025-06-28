package game

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate"
	"github.com/pedrokunz/go-design-patterns/internal/domain/room"
)

func Create(generator common.UUIDGenerator) (*Game, error) {
	if generator == nil {
		generator = uuid.NewRandom
	}

	gameID, err := generator()
	if err != nil {
		return nil, err
	}

	domainAggregate, newDomainAggregateErr := aggregate.NewAggregate(gameID, Aggregate)
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

	event, err := aggregate.NewEventBuilder(
		game.Aggregate,
		payload,
		Created,
	).Build()
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
