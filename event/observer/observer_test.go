package observer_test

import (
	"github.com/pedrokunz/go-design-patterns/event"
	"github.com/pedrokunz/go-design-patterns/event/observer"
	"github.com/stretchr/testify/require"
	"testing"
)

type MockObserver struct {
	on      func(event event.Event) error
	events  []event.Event
	onCalls int
}

func NewMockObserver() *MockObserver {
	return &MockObserver{
		events: make([]event.Event, 0),
	}
}

func (observer *MockObserver) On(event event.Event) error {
	observer.onCalls++
	observer.events = append(observer.events, event)

	if observer.on != nil {
		return observer.on(event)
	}

	return nil
}

func TestObserver(t *testing.T) {
	t.Run("succeeds", func(t *testing.T) {
		t.Run("when player observer reacts", func(t *testing.T) {
			playerObserver, newErr := observer.New(
				observer.PlayerObserver,
				observer.PlayerObserverConfig{
					Name: "player_name",
				},
			)

			require.NoError(t, newErr, "error building player observer")
			require.NotNil(t, playerObserver, "player observer should not be nil")

			playerJoinedEvent := event.New(event.PlayerJoined)
			onErr := playerObserver.On(playerJoinedEvent)

			require.NoError(t, onErr, "error notifying player observer")
		})
	})

	t.Run("fails", func(t *testing.T) {
		t.Run("when observer kind is not supported", func(t *testing.T) {
			_, newErr := observer.New("unknown", nil)

			require.EqualError(t, newErr, "invalid observer kind", "error message should be 'invalid observer kind'")
		})

		t.Run("when observer config is invalid", func(t *testing.T) {
			_, newErr := observer.New(observer.PlayerObserver, nil)

			require.EqualError(
				t,
				newErr,
				"invalid config type",
				"error message should be 'invalid config type'",
			)
		})

		t.Run("when observer event is invalid", func(t *testing.T) {
			playerObserver, newErr := observer.New(
				observer.PlayerObserver,
				observer.PlayerObserverConfig{
					Name: "player_name",
				},
			)

			require.NoError(t, newErr, "error building player observer")
			require.NotNil(t, playerObserver, "player observer should not be nil")

			onErr := playerObserver.On(nil)
			require.EqualError(
				t,
				onErr,
				"event cannot be nil",
				"error message should be 'event cannot be nil'",
			)
		})
	})
}
