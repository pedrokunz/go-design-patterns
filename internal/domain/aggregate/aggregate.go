package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/types"
)

const (
	ErrInvalidAggregateID = "invalid aggregate ID"
)

type Aggregate struct {
	ID      uuid.UUID
	Type    types.AggregateType
	Version int
	events  []Event
}

func NewAggregate(id uuid.UUID, aggregateType types.AggregateType) (*Aggregate, error) {
	if id == uuid.Nil {
		return nil, errors.New(ErrInvalidAggregateID)
	}

	return &Aggregate{
		ID:      id,
		Type:    aggregateType,
		Version: 0,
		events:  make([]Event, 0),
	}, nil
}

func (a *Aggregate) AddEvent(event Event) {
	a.events = append(a.events, event)
}

func (a *Aggregate) Events() []Event {
	return a.events
}

func (a *Aggregate) ApplyEvent(event Event) {
	a.AddEvent(event)
	a.Version++
}

func (a *Aggregate) ClearEvents() {
	a.events = make([]Event, 0)
}
