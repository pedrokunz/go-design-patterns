package internal_test

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal/aggregate"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAggregate(t *testing.T) {
	// Arrange
	aggregateID := common.NewDeterministicUUID("player-1")
	aggregateType := player.Aggregate

	t.Run("should create a new aggregate", func(t *testing.T) {
		// Act
		Aggregate, newAggregateErr := internal.NewAggregate(
			aggregateID,
			aggregateType,
		)

		// Assert
		require.NoError(t, newAggregateErr)
		require.Equal(t, aggregateID, Aggregate.ID)
		require.Equal(t, aggregateType, Aggregate.Type)
	})

	t.Run("should return error when creating a new aggregate", func(t *testing.T) {
		t.Run("with invalid aggregate ID", func(t *testing.T) {
			invalidAggregateID := uuid.Nil

			Aggregate, newAggregateErr := internal.NewAggregate(
				invalidAggregateID,
				player.Aggregate,
			)

			// Assert
			require.Error(t, newAggregateErr)
			require.Nil(t, Aggregate)
			require.EqualError(t, newAggregateErr, internal.ErrInvalidAggregateID)
		})

		t.Run("with invalid aggregate type", func(t *testing.T) {
			invalidAggregateType := aggregate.Type("")

			Aggregate, newAggregateErr := internal.NewAggregate(
				aggregateID,
				invalidAggregateType,
			)

			// Assert
			require.Error(t, newAggregateErr)
			require.Nil(t, Aggregate)
			require.EqualError(t, newAggregateErr, internal.ErrInvalidAggregateType)
		})
	})

	t.Run("should apply event to aggregate", func(t *testing.T) {
		Aggregate, newAggregateErr := internal.NewAggregate(
			aggregateID,
			aggregateType,
		)

		require.NoError(t, newAggregateErr)

		event, newDomainEventErr := internal.NewEventBuilder(
			Aggregate,
			[]byte("event payload"),
			player.Created,
		).Build()

		require.NoError(t, newDomainEventErr)

		Aggregate.ApplyEvent(event)

		t.Run("should update aggregate version and events", func(t *testing.T) {
			require.Equal(t, 1, Aggregate.Version)
			require.Len(t, Aggregate.Events(), 1)
			require.Equal(t, event, Aggregate.Events()[0])
		})

		t.Run("should flush events", func(t *testing.T) {
			Aggregate.FlushEvents()

			require.Len(t, Aggregate.Events(), 0)
		})
	})
}
