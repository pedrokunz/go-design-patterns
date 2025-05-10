package event_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/event"
	"testing"
)

func TestEvent(t *testing.T) {
	t.Run("succeeds", func(t *testing.T) {
		t.Run("when event is created", func(t *testing.T) {
			evt := event.New(event.PlayerJoined)

			if evt == nil {
				t.Error("event should not be nil")
			}

			if evt.Type() != event.PlayerJoined {
				t.Errorf("event type should be %s, got %s", event.PlayerJoined, evt.Type())
			}
		})
	})

	t.Run("fails", func(t *testing.T) {
		t.Run("when event type is unknown", func(t *testing.T) {
			evt := event.New("unknown")

			if evt == nil {
				t.Error("event should not be nil")
			}

			if evt.Type() != "unknown" {
				t.Errorf("event type should be 'unknown', got %s", evt.Type())
			}
		})
	})
}
