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

func TestAggregate(t *testing.T) {
	// Arrange
	aggregateID := common.NewDeterministicUUID("player-1")
	aggregateType := types.PlayerAggregateType
	aggregateVersion := 1

	t.Run("should create a new aggregate", func(t *testing.T) {
		// Act
		aggregate, newAggregateErr := eventsourcing.NewDomainAggregate(
			aggregateID,
			aggregateType,
			aggregateVersion,
		)

		// Assert
		require.NoError(t, newAggregateErr)
		require.Equal(t, aggregateID, aggregate.ID())
		require.Equal(t, aggregateType, aggregate.Type())
		require.Equal(t, aggregateVersion, aggregate.Version())
	})

	t.Run("should return error when creating a new aggregate", func(t *testing.T) {
		t.Run("with invalid aggregate ID", func(t *testing.T) {
			invalidAggregateID := uuid.Nil

			aggregate, newAggregateErr := eventsourcing.NewDomainAggregate(
				invalidAggregateID,
				types.PlayerAggregateType,
				aggregateVersion,
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
				aggregateVersion,
			)

			// Assert
			require.Error(t, newAggregateErr)
			require.Nil(t, aggregate)
			require.EqualError(t, newAggregateErr, eventsourcing.ErrInvalidAggregateType)
		})

		t.Run("with invalid aggregate version", func(t *testing.T) {
			invalidAggregateVersion := -1

			aggregate, newAggregateErr := eventsourcing.NewDomainAggregate(
				aggregateID,
				aggregateType,
				invalidAggregateVersion,
			)

			// Assert
			require.Error(t, newAggregateErr)
			require.Nil(t, aggregate)
			require.EqualError(t, newAggregateErr, eventsourcing.ErrInvalidAggregateVersion)
		})
	})

	t.Run("should apply event to aggregate", func(t *testing.T) {
		aggregate, newAggregateErr := eventsourcing.NewDomainAggregate(
			aggregateID,
			aggregateType,
			aggregateVersion,
		)

		require.NoError(t, newAggregateErr)

		event, newDomainEventErr := eventsourcing.NewDomainEvent(
			aggregateID,
			aggregateType,
			aggregateVersion,
			common.NewDeterministicUUID("event-1"),
			[]byte("event payload"),
			time.Date(2025, 5, 10, 1, 2, 3, 0, time.UTC),
			types.PlayerCreatedEventType,
			nil,
			nil,
		)

		require.NoError(t, newDomainEventErr)

		aggregate.ApplyEvent(event)

		t.Run("should update aggregate version and changes", func(t *testing.T) {
			require.Equal(t, aggregateVersion+1, aggregate.Version())
			require.Equal(t, 1, len(aggregate.Changes()))
			require.Equal(t, event, aggregate.Changes()[0])
		})

		t.Run("should flush changes", func(t *testing.T) {
			aggregate.FlushChanges()

			require.Equal(t, 0, len(aggregate.Changes()))
		})
	})
}
