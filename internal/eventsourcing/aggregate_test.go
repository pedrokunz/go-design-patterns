package eventsourcing_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing"
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAggregate(t *testing.T) {
	t.Run("should create a new aggregate", func(t *testing.T) {
		// Arrange
		aggregateID := common.NewDeterministicUUID("player-1")
		aggregateType := types.PlayerAggregateType
		aggregateVersion := 1

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
}
