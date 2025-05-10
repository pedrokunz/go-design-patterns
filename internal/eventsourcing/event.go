package eventsourcing

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing/types"
	"time"
)

// Event represents an immutable record of something that happened in the domain.
// Events are the building blocks of event sourcing, capturing all state changes
// within the system. Each event belongs to a specific aggregate and is part of
// that aggregate's event stream.
type Event interface {
	// AggregateID Identifies which aggregate instance this event belongs to
	AggregateID() uuid.UUID
	// AggregateType Defines the type of aggregate
	AggregateType() types.AggregateType
	// AggregateVersion Indicates the version of the aggregate after this event is applied
	AggregateVersion() int

	// ID Unique identifier for this specific event instance
	ID() uuid.UUID
	// Payload Raw event data as byte array, allowing for storage-agnostic serialization
	Payload() []byte
	// RecordedAt When this event was recorded in the event store
	RecordedAt() time.Time
	// Type The descriptive type name of this event
	Type() types.EventType

	// CausationID Identifies the command or event that triggered this event,
	// creating a causal chain for tracing purposes
	CausationID() *uuid.UUID
	// Metadata Additional contextual information about the event
	Metadata() map[string]string
}

type DomainEvent struct {
	aggregateID      uuid.UUID
	aggregateType    types.AggregateType
	aggregateVersion int
	id               uuid.UUID
	payload          []byte
	recordedAt       time.Time
	eventType        types.EventType
	causationID      *uuid.UUID
	metadata         map[string]string
}

func NewDomainEvent(
	aggregateID uuid.UUID,
	aggregateType types.AggregateType,
	aggregateVersion int,
	id uuid.UUID,
	payload []byte,
	recordedAt time.Time,
	eventType types.EventType,
	causationID *uuid.UUID,
	metadata map[string]string,
) (Event, error) {
	err := validateDomainEventInput(
		aggregateID,
		aggregateType,
		aggregateVersion,
		id,
		payload,
		recordedAt,
		eventType,
		causationID,
	)
	if err != nil {
		return nil, err
	}

	return &DomainEvent{
		aggregateID:      aggregateID,
		aggregateType:    aggregateType,
		aggregateVersion: aggregateVersion,
		id:               id,
		payload:          payload,
		recordedAt:       recordedAt,
		eventType:        eventType,
		causationID:      causationID,
		metadata:         metadata,
	}, nil
}

func (d DomainEvent) AggregateID() uuid.UUID {
	return d.aggregateID
}

func (d DomainEvent) AggregateType() types.AggregateType {
	return d.aggregateType
}

func (d DomainEvent) AggregateVersion() int {
	return d.aggregateVersion
}

func (d DomainEvent) ID() uuid.UUID {
	return d.id
}

func (d DomainEvent) Payload() []byte {
	return d.payload
}

func (d DomainEvent) RecordedAt() time.Time {
	return d.recordedAt
}

func (d DomainEvent) Type() types.EventType {
	return d.eventType
}

func (d DomainEvent) CausationID() *uuid.UUID {
	return d.causationID
}

func (d DomainEvent) Metadata() map[string]string {
	return d.metadata
}

func validateDomainEventInput(
	aggregateID uuid.UUID,
	aggregateType types.AggregateType,
	aggregateVersion int,
	id uuid.UUID,
	payload []byte,
	recordedAt time.Time,
	eventType types.EventType,
	causationID *uuid.UUID,
) error {
	if aggregateID == uuid.Nil {
		return errors.New(types.ErrInvalidAggregateID)
	}

	if !aggregateType.IsValid() {
		return errors.New(types.ErrInvalidAggregateType)
	}

	if aggregateVersion < 0 {
		return errors.New(types.ErrInvalidAggregateVersion)
	}

	if id == uuid.Nil {
		return errors.New(types.ErrInvalidEventID)
	}

	if len(payload) == 0 {
		return errors.New(types.ErrInvalidEventPayload)
	}

	if recordedAt.IsZero() {
		return errors.New(types.ErrInvalidEventRecordedAt)
	}

	if !eventType.IsValid() {
		return errors.New(types.ErrInvalidEventType)
	}

	if causationID != nil && *causationID == uuid.Nil {
		return errors.New(types.ErrInvalidCausationID)
	}

	return nil
}
