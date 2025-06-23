package internal_test

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal/event"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvent(t *testing.T) {
	// Arrange
	aggregateID := common.NewDeterministicUUID("player-1")
	aggregateType := player.Aggregate
	aggregate, newDomainAggregateErr := internal.NewAggregate(
		aggregateID,
		aggregateType,
	)

	require.NoError(t, newDomainAggregateErr)

	payload := []byte(`{"name":"player1"}`)
	eventType := player.Created
	causationID := common.NewDeterministicUUID("command-or-event-1")
	metadata := map[string]string{"key": "value"}

	t.Run("should create a new event", func(t *testing.T) {
		t.Run("with all parameters", func(t *testing.T) {
			Event, newDomainEventErr := internal.NewEventBuilder(
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
			require.Equal(t, aggregateID, Event.AggregateID())
			require.Equal(t, aggregateType, Event.AggregateType())
			require.Equal(t, payload, Event.Payload())
			require.Equal(t, eventType, Event.Type())
			require.Equal(t, &causationID, Event.CausationID())
			require.Equal(t, metadata, Event.Metadata())
		})

		t.Run("with minimal parameters", func(t *testing.T) {
			Event, newDomainEventErr := internal.NewEventBuilder(
				aggregate,
				payload,
				eventType,
			).Build()

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, Event.AggregateID())
			require.Equal(t, aggregateType, Event.AggregateType())
			require.Equal(t, payload, Event.Payload())
			require.Equal(t, eventType, Event.Type())
			require.Nil(t, Event.CausationID())
			require.Empty(t, Event.Metadata())
		})

		t.Run("without causation ID", func(t *testing.T) {
			Event, newDomainEventErr := internal.NewEventBuilder(
				aggregate,
				payload,
				eventType,
			).WithMetadata(
				metadata,
			).Build()

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, Event.AggregateID())
			require.Equal(t, aggregateType, Event.AggregateType())
			require.Equal(t, payload, Event.Payload())
			require.Equal(t, eventType, Event.Type())
			require.Nil(t, Event.CausationID())
			require.Equal(t, metadata, Event.Metadata())
		})

		t.Run("without metadata", func(t *testing.T) {
			Event, newDomainEventErr := internal.NewEventBuilder(
				aggregate,
				payload,
				eventType,
			).WithCausationID(
				&causationID,
			).Build()

			// Assert
			require.NoError(t, newDomainEventErr)
			require.Equal(t, aggregateID, Event.AggregateID())
			require.Equal(t, aggregateType, Event.AggregateType())
			require.Equal(t, payload, Event.Payload())
			require.Equal(t, eventType, Event.Type())
			require.Equal(t, &causationID, Event.CausationID())
			require.Empty(t, Event.Metadata())
		})
	})

	t.Run("should return error when creating event with invalid parameters", func(t *testing.T) {
		t.Run("invalid event payload", func(t *testing.T) {
			invalidPayload := []byte("")

			Event, newDomainEventErr := internal.NewEventBuilder(
				aggregate,
				invalidPayload,
				eventType,
			).WithCausationID(
				&causationID,
			).WithMetadata(
				metadata,
			).Build()

			require.Error(t, newDomainEventErr)
			require.Nil(t, Event)
			require.EqualError(t, newDomainEventErr, internal.ErrInvalidEventPayload)
		})

		t.Run("invalid event type", func(t *testing.T) {
			invalidEventType := event.Type("")

			Event, newDomainEventErr := internal.NewEventBuilder(
				aggregate,
				payload,
				invalidEventType,
			).WithCausationID(
				&causationID,
			).WithMetadata(
				metadata,
			).Build()

			require.Error(t, newDomainEventErr)
			require.Nil(t, Event)
			require.EqualError(t, newDomainEventErr, internal.ErrInvalidEventType)
		})

		t.Run("invalid causation ID", func(t *testing.T) {
			invalidCausationID := uuid.Nil

			Event, newDomainEventErr := internal.NewEventBuilder(
				aggregate,
				payload,
				eventType,
			).WithCausationID(
				&invalidCausationID,
			).WithMetadata(
				metadata,
			).Build()

			require.Error(t, newDomainEventErr)
			require.Nil(t, Event)
			require.EqualError(t, newDomainEventErr, internal.ErrInvalidCausationID)
		})
	})
}
