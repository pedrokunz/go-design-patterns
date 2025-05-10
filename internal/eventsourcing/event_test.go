package eventsourcing_test

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	// Arrange
	aggregateID := common.NewDeterministicUUID("player-1")
	aggregateType := types.PlayerAggregateType
	aggregateVersion := 0
	id := common.NewDeterministicUUID("event-1")
	payload := []byte(`{"name":"player1"}`)
	recordedAt := time.Now()
	eventType := types.PlayerCreatedEventType
	causationID := common.NewDeterministicUUID("command-or-event-1")
	metadata := map[string]string{"key": "value"}

	t.Run("should create a new event", func(t *testing.T) {
		t.Run("with all parameters", func(t *testing.T) {
			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				aggregateType,
				aggregateVersion,
				id,
				payload,
				recordedAt,
				eventType,
				&causationID,
				metadata,
			)

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, event.AggregateID())
			require.Equal(t, aggregateType, event.AggregateType())
			require.Equal(t, aggregateVersion, event.AggregateVersion())
			require.Equal(t, id, event.ID())
			require.Equal(t, payload, event.Payload())
			require.Equal(t, recordedAt, event.RecordedAt())
			require.Equal(t, eventType, event.Type())
			require.Equal(t, &causationID, event.CausationID())
			require.Equal(t, metadata, event.Metadata())
		})

		t.Run("without causation ID", func(t *testing.T) {
			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				aggregateType,
				aggregateVersion,
				id,
				payload,
				recordedAt,
				eventType,
				nil,
				metadata,
			)

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, event.AggregateID())
			require.Equal(t, aggregateType, event.AggregateType())
			require.Equal(t, aggregateVersion, event.AggregateVersion())
			require.Equal(t, id, event.ID())
			require.Equal(t, payload, event.Payload())
			require.Equal(t, recordedAt, event.RecordedAt())
			require.Equal(t, eventType, event.Type())
			require.Nil(t, event.CausationID())
			require.Equal(t, metadata, event.Metadata())
		})

		t.Run("without metadata", func(t *testing.T) {
			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				aggregateType,
				aggregateVersion,
				id,
				payload,
				recordedAt,
				eventType,
				&causationID,
				nil,
			)

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, event.AggregateID())
			require.Equal(t, aggregateType, event.AggregateType())
			require.Equal(t, aggregateVersion, event.AggregateVersion())
			require.Equal(t, id, event.ID())
			require.Equal(t, payload, event.Payload())
			require.Equal(t, recordedAt, event.RecordedAt())
			require.Equal(t, eventType, event.Type())
			require.Equal(t, &causationID, event.CausationID())
			require.Nil(t, event.Metadata())
		})
	})

	t.Run("should return error when creating event with invalid parameters", func(t *testing.T) {
		t.Run("invalid aggregate ID", func(t *testing.T) {
			invalidAggregateID := uuid.Nil

			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				invalidAggregateID,
				aggregateType,
				aggregateVersion,
				id,
				payload,
				recordedAt,
				eventType,
				&causationID,
				metadata,
			)

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, types.ErrInvalidAggregateID)
		})

		t.Run("invalid aggregate type", func(t *testing.T) {
			invalidAggregateType := types.AggregateType("")

			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				invalidAggregateType,
				aggregateVersion,
				id,
				payload,
				recordedAt,
				eventType,
				&causationID,
				metadata,
			)

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, types.ErrInvalidAggregateType)
		})

		t.Run("invalid aggregate version", func(t *testing.T) {
			invalidAggregateVersion := -1

			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				aggregateType,
				invalidAggregateVersion,
				id,
				payload,
				recordedAt,
				eventType,
				&causationID,
				metadata,
			)

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, types.ErrInvalidAggregateVersion)
		})

		t.Run("invalid event ID", func(t *testing.T) {
			invalidEventID := uuid.Nil

			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				aggregateType,
				aggregateVersion,
				invalidEventID,
				payload,
				recordedAt,
				eventType,
				&causationID,
				metadata,
			)

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, types.ErrInvalidEventID)
		})

		t.Run("invalid event payload", func(t *testing.T) {
			invalidPayload := []byte("")

			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				aggregateType,
				aggregateVersion,
				id,
				invalidPayload,
				recordedAt,
				eventType,
				&causationID,
				metadata,
			)

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, types.ErrInvalidEventPayload)
		})

		t.Run("invalid event recorded at", func(t *testing.T) {
			invalidRecordedAt := time.Time{}

			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				aggregateType,
				aggregateVersion,
				id,
				payload,
				invalidRecordedAt,
				eventType,
				&causationID,
				metadata,
			)

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, types.ErrInvalidEventRecordedAt)
		})

		t.Run("invalid event type", func(t *testing.T) {
			invalidEventType := types.EventType("")

			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				aggregateType,
				aggregateVersion,
				id,
				payload,
				recordedAt,
				invalidEventType,
				&causationID,
				metadata,
			)

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, types.ErrInvalidEventType)
		})

		t.Run("invalid causation ID", func(t *testing.T) {
			invalidCausationID := uuid.Nil

			event, newDomainEventErr := eventsourcing.NewDomainEvent(
				aggregateID,
				aggregateType,
				aggregateVersion,
				id,
				payload,
				recordedAt,
				eventType,
				&invalidCausationID,
				metadata,
			)

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, types.ErrInvalidCausationID)
		})
	})
}
