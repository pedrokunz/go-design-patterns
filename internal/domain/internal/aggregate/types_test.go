package aggregate_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal/aggregate"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterAggregateType(t *testing.T) {
	t.Run("all aggregate types should be registered", func(t *testing.T) {
		expectedAggregateTypes := []aggregate.Type{
			player.Aggregate,
		}

		require.Equal(t, expectedAggregateTypes, aggregate.Types())
	})

	t.Run("invalid aggregate", func(t *testing.T) {
		t.Run("with invalid aggregate type", func(t *testing.T) {
			invalidAggregateType := aggregate.Type("InvalidAggregateType")
			if invalidAggregateType.IsValid() {
				t.Errorf("InvalidAggregateType should be invalid")
			}
		})

		t.Run("with empty aggregate type", func(t *testing.T) {
			emptyAggregateType := aggregate.Type("")
			if emptyAggregateType.IsValid() {
				t.Errorf("Empty AggregateType should be invalid")
			}
		})
	})
}
