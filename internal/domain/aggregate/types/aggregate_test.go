package types_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/types"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"testing"
)

func TestRegisterAggregateType(t *testing.T) {
	t.Run("all aggregate types should be registered", func(t *testing.T) {
		expectedAggregateTypes := map[types.AggregateType]bool{
			player.Aggregate: true,
			game.Aggregate:   true,
		}

		for _, aggregateType := range types.AggregateTypes() {
			if _, exists := expectedAggregateTypes[aggregateType]; !exists {
				t.Errorf("Aggregate type %s is not registered", aggregateType)
			}
		}
	})

	t.Run("invalid aggregate", func(t *testing.T) {
		t.Run("with invalid type", func(t *testing.T) {
			invalidAggregateType := types.AggregateType("InvalidAggregateType")
			if invalidAggregateType.IsValid() {
				t.Errorf("InvalidAggregateType should be invalid")
			}
		})

		t.Run("with empty type", func(t *testing.T) {
			var emptyAggregateType types.AggregateType
			if emptyAggregateType.IsValid() {
				t.Errorf("empty aggregate type should be invalid")
			}
		})
	})
}
