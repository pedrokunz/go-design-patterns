package internal

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal/aggregate"
)

const (
	ErrInvalidAggregateID      = "invalid aggregate ID"
	ErrInvalidAggregateType    = "invalid aggregate type"
	ErrInvalidAggregateVersion = "invalid aggregate version"
)

type Aggregate struct {
	ID      uuid.UUID      `json:"id"`
	Type    aggregate.Type `json:"type"`
	Version int            `json:"version"`

	events []Event
}

func NewAggregate(
	id uuid.UUID,
	aggregateType aggregate.Type,
) (*Aggregate, error) {
	err := validateDomainAggregateInput(id, aggregateType)
	if err != nil {
		return nil, err
	}

	return &Aggregate{
		ID:     id,
		Type:   aggregateType,
		events: []Event{},
	}, nil
}

func (d *Aggregate) ApplyEvent(event Event) {
	d.Version++
	d.events = append(d.events, event)
}

func (d *Aggregate) Events() []Event {
	return d.events
}

func (d *Aggregate) FlushEvents() {
	d.events = []Event{}
}

func validateDomainAggregateInput(id uuid.UUID, aggregateType aggregate.Type) error {
	if id == uuid.Nil {
		return fmt.Errorf(ErrInvalidAggregateID)
	}

	if !aggregateType.IsValid() {
		return fmt.Errorf(ErrInvalidAggregateType)
	}

	return nil
}
