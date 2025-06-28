package player

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
)

func Create(generator common.UUIDGenerator, name string) (*Player, error) {
	if generator == nil {
		generator = uuid.NewRandom
	}

	playerID, err := generator()
	if err != nil {
		return nil, err
	}

	if len(name) < 3 || len(name) > 10 {
		return nil, errors.New("player name must be between 3 and 10 characters")
	}

	domainAggregate, newDomainAggregateErr := aggregate.NewAggregate(playerID, Aggregate)
	if newDomainAggregateErr != nil {
		return nil, newDomainAggregateErr
	}

	player := &Player{
		Aggregate: domainAggregate,
		Name:      name,
		Armour:    internal.Armour{Value: 0},
		Attack:    internal.Attack{Min: 1, Max: 100},
		Life:      internal.Life{Value: 100},
	}

	payload, err := json.Marshal(player)
	if err != nil {
		return nil, err
	}

	event, err := aggregate.NewEventBuilder(
		player.Aggregate,
		payload,
		Created,
	).
		Build()
	if err != nil {
		return nil, err
	}

	err = player.Apply(event)
	if err != nil {
		return nil, err
	}

	player.Aggregate.ApplyEvent(event)

	return player, nil
}
