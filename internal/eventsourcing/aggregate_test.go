package eventsourcing_test

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAggregate(t *testing.T) {
	// Arrange
	aggregateID := common.NewDeterministicUUID("player-1")
	aggregateType := types.PlayerAggregate

	t.Run("should create a new aggregate", func(t *testing.T) {
		// Act
		aggregate, newAggregateErr := eventsourcing.NewDomainAggregate(
			aggregateID,
			aggregateType,
		)

		// Assert
		require.NoError(t, newAggregateErr)
		require.Equal(t, aggregateID, aggregate.ID())
		require.Equal(t, aggregateType, aggregate.Type())
	})

	t.Run("should return error when creating a new aggregate", func(t *testing.T) {
		t.Run("with invalid aggregate ID", func(t *testing.T) {
			invalidAggregateID := uuid.Nil

			aggregate, newAggregateErr := eventsourcing.NewDomainAggregate(
				invalidAggregateID,
				types.PlayerAggregate,
			)

			// Assert
			require.Error(t, newAggregateErr)
			require.Nil(t, aggregate)
			require.EqualError(t, newAggregateErr, eventsourcing.ErrInvalidAggregateID)
		})

		t.Run("with invalid aggregate type", func(t *testing.T) {
			invalidAggregateType := types.AggregateType("")

			aggregate, newAggregateErr := eventsourcing.NewDomainAggregate(
				aggregateID,
				invalidAggregateType,
			)

			// Assert
			require.Error(t, newAggregateErr)
			require.Nil(t, aggregate)
			require.EqualError(t, newAggregateErr, eventsourcing.ErrInvalidAggregateType)
		})
	})

	t.Run("should apply event to aggregate", func(t *testing.T) {
		aggregate, newAggregateErr := eventsourcing.NewDomainAggregate(
			aggregateID,
			aggregateType,
		)

		require.NoError(t, newAggregateErr)

		event, newDomainEventErr := eventsourcing.NewEventBuilder(
			aggregate,
			[]byte("event payload"),
			types.PlayerCreated,
		).Build()

		require.NoError(t, newDomainEventErr)

		aggregate.ApplyEvent(event)

		t.Run("should update aggregate version and events", func(t *testing.T) {
			require.Equal(t, 1, aggregate.Version())
			require.Len(t, aggregate.Events(), 1)
			require.Equal(t, event, aggregate.Events()[0])
		})

		t.Run("should flush events", func(t *testing.T) {
			aggregate.FlushEvents()

			require.Len(t, aggregate.Events(), 0)
		})
	})
}
