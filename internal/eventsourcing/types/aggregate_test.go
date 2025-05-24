package types_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/eventsourcing/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterAggregateType(t *testing.T) {
	t.Run("all aggregate types should be registered", func(t *testing.T) {
		expectedAggregateTypes := []types.AggregateType{
			types.PlayerAggregate,
		}

		require.Equal(t, expectedAggregateTypes, types.AggregateTypes())
	})

	t.Run("invalid aggregate", func(t *testing.T) {
		t.Run("with invalid aggregate type", func(t *testing.T) {
			invalidAggregateType := types.AggregateType("InvalidAggregateType")
			if invalidAggregateType.IsValid() {
				t.Errorf("InvalidAggregateType should be invalid")
			}
		})

		t.Run("with empty aggregate type", func(t *testing.T) {
			emptyAggregateType := types.AggregateType("")
			if emptyAggregateType.IsValid() {
				t.Errorf("Empty AggregateType should be invalid")
			}
		})
	})
}
