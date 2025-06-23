package event_test

import (
	eventTypes "github.com/pedrokunz/go-design-patterns/internal/domain/internal/event"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterEventType(t *testing.T) {
	t.Run("all event types should be registered", func(t *testing.T) {
		expectedEventTypes := []eventTypes.Type{
			player.Created,
		}

		require.Equal(t, expectedEventTypes, eventTypes.Types())
	})

	t.Run("invalid event", func(t *testing.T) {
		t.Run("with invalid event type", func(t *testing.T) {
			invalidEventType := eventTypes.Type("InvalidEventType")
			if invalidEventType.IsValid() {
				t.Errorf("InvalidEventType should be invalid")
			}
		})

		t.Run("with empty event type", func(t *testing.T) {
			emptyEventType := eventTypes.Type("")
			if emptyEventType.IsValid() {
				t.Errorf("Empty EventType should be invalid")
			}
		})
	})
}
