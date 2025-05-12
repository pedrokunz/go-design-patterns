package types_test

import (
	eventTypes "github.com/pedrokunz/go-design-patterns/internal/eventsourcing/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterEventType(t *testing.T) {
	t.Run("all event types should be registered", func(t *testing.T) {
		expectedEventTypes := []eventTypes.EventType{
			eventTypes.PlayerCreatedEventType,
		}

		require.Equal(t, expectedEventTypes, eventTypes.EventTypes())
	})

	t.Run("invalid event", func(t *testing.T) {
		t.Run("with invalid event type", func(t *testing.T) {
			invalidEventType := eventTypes.EventType("InvalidEventType")
			if invalidEventType.IsValid() {
				t.Errorf("InvalidEventType should be invalid")
			}
		})

		t.Run("with empty event type", func(t *testing.T) {
			emptyEventType := eventTypes.EventType("")
			if emptyEventType.IsValid() {
				t.Errorf("Empty EventType should be invalid")
			}
		})
	})
}
