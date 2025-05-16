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

type domainEvent struct {
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

type EventBuilder struct {
	event *domainEvent
	err   error
}

func NewEventBuilder(
	domainAggregate Aggregate,
	payload []byte,
	eventType types.EventType,
) *EventBuilder {
	err := validateDomainEventInput(
		payload,
		eventType,
	)
	if err != nil {
		return &EventBuilder{err: err}
	}

	id := uuid.New()
	recordedAt := time.Now()

	event := &domainEvent{
		aggregateID:      domainAggregate.ID(),
		aggregateType:    domainAggregate.Type(),
		aggregateVersion: domainAggregate.Version(),
		id:               id,
		payload:          payload,
		recordedAt:       recordedAt,
		eventType:        eventType,
		causationID:      nil,
		metadata:         make(map[string]string),
	}

	return &EventBuilder{event: event, err: nil}
}

func (b *EventBuilder) WithCausationID(causationID *uuid.UUID) *EventBuilder {
	if b.event == nil {
		return b
	}

	if causationID != nil && *causationID == uuid.Nil {
		b.err = errors.New(ErrInvalidCausationID)
		return b
	}

	b.event.causationID = causationID
	return b
}

func (b *EventBuilder) WithMetadata(metadata map[string]string) *EventBuilder {
	if b.event == nil || metadata == nil {
		return b
	}

	b.event.metadata = metadata
	return b
}

func (b *EventBuilder) Build() (Event, error) {
	if b.err != nil {
		return nil, b.err
	}

	return b.event, nil
}

func (d domainEvent) AggregateID() uuid.UUID {
	return d.aggregateID
}

func (d domainEvent) AggregateType() types.AggregateType {
	return d.aggregateType
}

func (d domainEvent) AggregateVersion() int {
	return d.aggregateVersion
}

func (d domainEvent) ID() uuid.UUID {
	return d.id
}

func (d domainEvent) Payload() []byte {
	return d.payload
}

func (d domainEvent) RecordedAt() time.Time {
	return d.recordedAt
}

func (d domainEvent) Type() types.EventType {
	return d.eventType
}

func (d domainEvent) CausationID() *uuid.UUID {
	return d.causationID
}

func (d domainEvent) Metadata() map[string]string {
	return d.metadata
}

const (
	ErrInvalidEventPayload = "invalid event payload"
	ErrInvalidEventType    = "invalid event type"
	ErrInvalidCausationID  = "invalid causation ID"
)

func validateDomainEventInput(payload []byte, eventType types.EventType) error {
	if len(payload) == 0 {
		return errors.New(ErrInvalidEventPayload)
	}

	if !eventType.IsValid() {
		return errors.New(ErrInvalidEventType)
	}

	return nil
}
