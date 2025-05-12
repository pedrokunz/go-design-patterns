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

	// Changes returns all new events that have been generated but not yet persisted
	Changes() []Event

	// FlushChanges removes all events after they have been successfully stored
	FlushChanges()
}

type domainAggregate struct {
	id            uuid.UUID
	aggregateType types.AggregateType
	version       int
	changes       []Event
}

func NewDomainAggregate(
	id uuid.UUID,
	aggregateType types.AggregateType,
	version int,
) (Aggregate, error) {
	err := validateDomainAggregateInput(id, aggregateType, version)
	if err != nil {
		return nil, err
	}

	return &domainAggregate{
		id:            id,
		aggregateType: aggregateType,
		version:       version,
		changes:       []Event{},
	}, nil
}

func (d *domainAggregate) ID() uuid.UUID {
	return d.id
}

func (d *domainAggregate) Type() types.AggregateType {
	return d.aggregateType
}

func (d *domainAggregate) Version() int {
	return d.version
}

func (d *domainAggregate) ApplyEvent(event Event) {
	d.version++
	d.changes = append(d.changes, event)
}

func (d *domainAggregate) Changes() []Event {
	return d.changes
}

func (d *domainAggregate) FlushChanges() {
	d.changes = []Event{}
}

const (
	ErrInvalidAggregateID      = "invalid aggregate ID"
	ErrInvalidAggregateType    = "invalid aggregate type"
	ErrInvalidAggregateVersion = "invalid aggregate version"
)

func validateDomainAggregateInput(id uuid.UUID, aggregateType types.AggregateType, version int) error {
	if id == uuid.Nil {
		return fmt.Errorf(ErrInvalidAggregateID)
	}

	if !aggregateType.IsValid() {
		return fmt.Errorf(ErrInvalidAggregateType)
	}

	if version < 0 {
		return fmt.Errorf(ErrInvalidAggregateVersion)
	}

	return nil
}
