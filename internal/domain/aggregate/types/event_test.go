package types_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/types"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"testing"
)

func TestRegisterEventType(t *testing.T) {
	t.Run("all event types should be registered", func(t *testing.T) {
		expectedEventTypes := map[types.EventType]bool{
			game.Created:      true,
			game.RoomsCreated: true,
			game.PlayerAdded:  true,
			player.Created:    true,
		}

		for _, eventType := range types.EventTypes() {
			if _, exists := expectedEventTypes[eventType]; !exists {
				t.Errorf("Event type %s is not registered", eventType)
			}
		}
	})

	t.Run("invalid event", func(t *testing.T) {
		t.Run("with invalid type", func(t *testing.T) {
			invalidEventType := types.EventType("InvalidEventType")

			if invalidEventType.IsValid() {
				t.Errorf("InvalidEventType should be invalid")
			}
		})

		t.Run("with empty type", func(t *testing.T) {
			var emptyEventType types.EventType
			if emptyEventType.IsValid() {
				t.Errorf("empty EventType should be invalid")
			}
		})
	})
}
