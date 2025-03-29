package observer_test

import (
	"errors"
	"github.com/pedrokunz/go-design-patterns/event/observer"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/pedrokunz/go-design-patterns/event"
)

func TestNotifier(t *testing.T) {
	t.Run("succeeds", func(t *testing.T) {
		t.Run("when attach observers", func(t *testing.T) {
			notifier := observer.NewNotifier()
			playerObserver, newErr := observer.New(
				observer.PlayerObserver,
				observer.PlayerObserverConfig{
					Name: "player_name",
				})

			require.NoError(t, newErr, "error building player observer")

			attachErr := notifier.Attach(playerObserver)

			require.NoError(t, attachErr, "error attaching player observer")
		})

		t.Run("when notifies observers", func(t *testing.T) {
			notifier := observer.NewNotifier()

			mockObserver := NewMockObserver()

			attachErr := notifier.Attach(mockObserver)
			require.NoError(t, attachErr, "error attaching observer")

			notifyPlayerJoinedErr := notifier.Notify(event.New(event.PlayerJoined))
			require.NoError(t, notifyPlayerJoinedErr, "error notifying player joined")

			notifyPlayerLeftErr := notifier.Notify(event.New(event.PlayerLeft))
			require.NoError(t, notifyPlayerLeftErr, "error notifying player left")

			require.Equal(t, 2, mockObserver.onCalls, "Observer should call 2 events")
			require.Len(t, mockObserver.events, 2, "Observer should have received 2 events")
		})
	})

	t.Run("fails", func(t *testing.T) {
		t.Run("when attaching with invalid observer", func(t *testing.T) {
			notifier := observer.NewNotifier()

			attachErr := notifier.Attach(nil)
			require.EqualError(
				t,
				errors.New("observer cannot be nil"),
				attachErr.Error(),
				"error message should be 'observer cannot be nil'",
			)
		})

		t.Run("when notifying observers with invalid event", func(t *testing.T) {
			notifier := observer.NewNotifier()

			observer1, newErr := observer.New(
				observer.PlayerObserver,
				observer.PlayerObserverConfig{
					Name: "player_name_1",
				},
			)
			require.NoError(t, newErr, "error building player observer 1")

			observer2, newErr := observer.New(
				observer.PlayerObserver,
				observer.PlayerObserverConfig{
					Name: "player_name_2",
				},
			)
			require.NoError(t, newErr, "error building player observer 2")

			attachErr := notifier.Attach(observer1)
			require.NoError(t, attachErr, "error attaching observer 1")

			attachErr = notifier.Attach(observer2)
			require.NoError(t, attachErr, "error attaching observer 2")

			notifyPlayerJoinedErr := notifier.Notify(nil)
			require.EqualError(
				t,
				errors.New("event cannot be nil"),
				notifyPlayerJoinedErr.Error(),
				"error message should be 'event cannot be nil'",
			)
		})

		t.Run("when notifying observers with error", func(t *testing.T) {
			notifier := observer.NewNotifier()

			mockObserver1WithError := NewMockObserver()
			mockObserver1WithError.on = func(event event.Event) error {
				return errors.New("observer error")
			}

			mockObserver2WithError := NewMockObserver()
			mockObserver2WithError.on = func(event event.Event) error {
				return errors.New("observer error 2")
			}

			mockObserver3WithoutError := NewMockObserver()

			attachErr := notifier.Attach(mockObserver1WithError)
			require.NoError(t, attachErr, "error attaching observer")

			notifyPlayerJoinedErr := notifier.Notify(event.New(event.PlayerJoined))
			require.EqualError(
				t,
				errors.New("observer error"),
				notifyPlayerJoinedErr.Error(),
				"error message should be 'observer error'",
			)

			require.Equal(t, 1, mockObserver1WithError.onCalls, "Observer should call 1 event")
			require.Len(t, mockObserver1WithError.events, 1, "Observer should have received 1 event")
			require.Equal(
				t,
				event.PlayerJoined,
				mockObserver1WithError.events[0].Type(),
				"Observer should have received player joined event",
			)

			attachErr = notifier.Attach(mockObserver2WithError)
			require.NoError(t, attachErr, "error attaching observer 2")

			notifyPlayerJoined2Err := notifier.Notify(event.New(event.PlayerJoined))
			require.EqualError(
				t,
				errors.New("observer error; observer error 2"),
				notifyPlayerJoined2Err.Error(),
				"error message should be 'observer error; observer error 2'",
			)

			require.Equal(t, 1, mockObserver2WithError.onCalls, "Observer should call 1 event")
			require.Len(t, mockObserver2WithError.events, 1, "Observer should have received 1 event")
			require.Equal(
				t,
				event.PlayerJoined,
				mockObserver2WithError.events[0].Type(),
				"Observer should have received player joined event",
			)

			require.Equal(t, 2, mockObserver1WithError.onCalls, "Observer should call 2 events")
			require.Len(t, mockObserver1WithError.events, 2, "Observer should have received 2 events")
			require.Equal(
				t,
				event.PlayerJoined,
				mockObserver1WithError.events[1].Type(),
				"Observer should have received player joined event",
			)

			attachErr = notifier.Attach(mockObserver3WithoutError)
			require.NoError(t, attachErr, "error attaching observer 3")

			notifyPlayerLeftErr := notifier.Notify(event.New(event.PlayerLeft))
			require.EqualError(
				t,
				errors.New("observer error; observer error 2"),
				notifyPlayerLeftErr.Error(),
				"error message should be 'observer error; observer error 2'",
			)

			require.Equal(t, 1, mockObserver3WithoutError.onCalls, "Observer should call 1 event")
			require.Len(t, mockObserver3WithoutError.events, 1, "Observer should have received 1 event")
			require.Equal(
				t,
				event.PlayerLeft,
				mockObserver3WithoutError.events[0].Type(),
				"Observer should have received player left event",
			)

			require.Equal(t, 3, mockObserver1WithError.onCalls, "Observer should call 3 events")
			require.Len(t, mockObserver1WithError.events, 3, "Observer should have received 3 events")
			require.Equal(
				t,
				event.PlayerLeft,
				mockObserver1WithError.events[2].Type(),
				"Observer should have received player left event",
			)

			require.Equal(t, 2, mockObserver2WithError.onCalls, "Observer should call 2 events")
			require.Len(t, mockObserver2WithError.events, 2, "Observer should have received 2 events")
			require.Equal(
				t,
				event.PlayerLeft,
				mockObserver2WithError.events[1].Type(),
				"Observer should have received player left event",
			)
		})
	})
}

type ExampleEvent struct{}

func (event *ExampleEvent) Type() event.Kind {
	return "hello"
}
