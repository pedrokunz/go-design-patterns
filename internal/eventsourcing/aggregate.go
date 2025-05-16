package eventsourcing

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing/types"
)

// Aggregate represents a cluster of domain objects treated as a single unit for data changes.
// Aggregates are the primary building blocks in event sourcing that maintain consistency
// boundaries and encapsulate business rules. Each aggregate has a unique identity and a
// sequence of events that represent its state changes over time.
type Aggregate interface {
	// ID returns the unique identifier for this aggregate instance
	ID() uuid.UUID

	// Type returns the type name of this aggregate, used for categorization and routing
	Type() types.AggregateType

	// Version returns the current version of the aggregate, which increases with each applied event
	Version() int

	// ApplyEvent applies an event to the aggregate, updating its internal state
	ApplyEvent(event Event)

	// Events returns all new events that have been generated but not yet persisted
	Events() []Event

	// FlushEvents removes all events after they have been successfully stored
	FlushEvents()
}

type DomainAggregate struct {
	id            uuid.UUID
	aggregateType types.AggregateType
	version       int
	events        []Event
}

func NewDomainAggregate(
	id uuid.UUID,
	aggregateType types.AggregateType,
) (Aggregate, error) {
	err := validateDomainAggregateInput(id, aggregateType)
	if err != nil {
		return nil, err
	}

	return &DomainAggregate{
		id:            id,
		aggregateType: aggregateType,
		events:        []Event{},
	}, nil
}

func (d *DomainAggregate) ID() uuid.UUID {
	return d.id
}

func (d *DomainAggregate) Type() types.AggregateType {
	return d.aggregateType
}

func (d *DomainAggregate) Version() int {
	return d.version
}

func (d *DomainAggregate) ApplyEvent(event Event) {
	d.version++
	d.events = append(d.events, event)
}

func (d *DomainAggregate) Events() []Event {
	return d.events
}

func (d *DomainAggregate) FlushEvents() {
	d.events = []Event{}
}

const (
	ErrInvalidAggregateID   = "invalid aggregate ID"
	ErrInvalidAggregateType = "invalid aggregate type"
)

func validateDomainAggregateInput(id uuid.UUID, aggregateType types.AggregateType) error {
	if id == uuid.Nil {
		return fmt.Errorf(ErrInvalidAggregateID)
	}

	if !aggregateType.IsValid() {
		return fmt.Errorf(ErrInvalidAggregateType)
	}

	return nil
}
