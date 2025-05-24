package eventsourcing_test

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvent(t *testing.T) {
	// Arrange
	aggregateID := common.NewDeterministicUUID("player-1")
	aggregateType := types.PlayerAggregate
	aggregate, newDomainAggregateErr := eventsourcing.NewDomainAggregate(
		aggregateID,
		aggregateType,
	)

	require.NoError(t, newDomainAggregateErr)

	payload := []byte(`{"name":"player1"}`)
	eventType := types.PlayerCreated
	causationID := common.NewDeterministicUUID("command-or-event-1")
	metadata := map[string]string{"key": "value"}

	t.Run("should create a new event", func(t *testing.T) {
		t.Run("with all parameters", func(t *testing.T) {
			event, newDomainEventErr := eventsourcing.NewEventBuilder(
				aggregate,
				payload,
				eventType,
			).WithCausationID(
				&causationID,
			).WithMetadata(
				metadata,
			).Build()

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, event.AggregateID())
			require.Equal(t, aggregateType, event.AggregateType())
			require.Equal(t, payload, event.Payload())
			require.Equal(t, eventType, event.Type())
			require.Equal(t, &causationID, event.CausationID())
			require.Equal(t, metadata, event.Metadata())
		})

		t.Run("with minimal parameters", func(t *testing.T) {
			event, newDomainEventErr := eventsourcing.NewEventBuilder(
				aggregate,
				payload,
				eventType,
			).Build()

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, event.AggregateID())
			require.Equal(t, aggregateType, event.AggregateType())
			require.Equal(t, payload, event.Payload())
			require.Equal(t, eventType, event.Type())
			require.Nil(t, event.CausationID())
			require.Empty(t, event.Metadata())
		})

		t.Run("without causation ID", func(t *testing.T) {
			event, newDomainEventErr := eventsourcing.NewEventBuilder(
				aggregate,
				payload,
				eventType,
			).WithMetadata(
				metadata,
			).Build()

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, event.AggregateID())
			require.Equal(t, aggregateType, event.AggregateType())
			require.Equal(t, payload, event.Payload())
			require.Equal(t, eventType, event.Type())
			require.Nil(t, event.CausationID())
			require.Equal(t, metadata, event.Metadata())
		})

		t.Run("without metadata", func(t *testing.T) {
			event, newDomainEventErr := eventsourcing.NewEventBuilder(
				aggregate,
				payload,
				eventType,
			).WithCausationID(
				&causationID,
			).Build()

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, event.AggregateID())
			require.Equal(t, aggregateType, event.AggregateType())
			require.Equal(t, payload, event.Payload())
			require.Equal(t, eventType, event.Type())
			require.Equal(t, &causationID, event.CausationID())
			require.Empty(t, event.Metadata())
		})
	})

	t.Run("should return error when creating event with invalid parameters", func(t *testing.T) {
		t.Run("invalid event payload", func(t *testing.T) {
			invalidPayload := []byte("")

			event, newDomainEventErr := eventsourcing.NewEventBuilder(
				aggregate,
				invalidPayload,
				eventType,
			).WithCausationID(
				&causationID,
			).WithMetadata(
				metadata,
			).Build()

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, eventsourcing.ErrInvalidEventPayload)
		})

		t.Run("invalid event type", func(t *testing.T) {
			invalidEventType := types.EventType("")

			event, newDomainEventErr := eventsourcing.NewEventBuilder(
				aggregate,
				payload,
				invalidEventType,
			).WithCausationID(
				&causationID,
			).WithMetadata(
				metadata,
			).Build()

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, eventsourcing.ErrInvalidEventType)
		})

		t.Run("invalid causation ID", func(t *testing.T) {
			invalidCausationID := uuid.Nil

			event, newDomainEventErr := eventsourcing.NewEventBuilder(
				aggregate,
				payload,
				eventType,
			).WithCausationID(
				&invalidCausationID,
			).WithMetadata(
				metadata,
			).Build()

			require.Error(t, newDomainEventErr)
			require.Nil(t, event)
			require.EqualError(t, newDomainEventErr, eventsourcing.ErrInvalidCausationID)
		})
	})
}
