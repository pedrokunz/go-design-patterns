package game_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/aggregate/game"
	"github.com/pedrokunz/go-design-patterns/event"
	"github.com/pedrokunz/go-design-patterns/event/observer"
	"github.com/stretchr/testify/require"
)

func TestNewState(t *testing.T) {
	t.Run("returns a valid state", func(t *testing.T) {
		actual := game.NewState()

		require.NotNil(t, actual, "NewState should not be nil")
		require.NotNil(t, actual.Notifier, "Notifier should not be nil")
		require.NotNil(t, actual.Rooms, "Rooms should not be nil")
		require.Len(t, actual.Rooms, 0, "Rooms should be empty")
	})

	t.Run("returns the same state", func(t *testing.T) {
		actual := game.NewState()
		expected := game.NewState()

		require.Equal(
			t,
			actual,
			expected,
			"States should be equal",
		)
	})

	t.Run("notifies multiple observers of events", func(t *testing.T) {
		state := game.NewState()
		mockSubject := &MockSubject{}
		state.Notifier = mockSubject

		observer1, observer1BuildErr := observer.NewFactory(observer.PlayerObserver).Build(
			observer.PlayerObserverConfig{
				Name: "george",
			},
		)

		require.NoError(t, observer1BuildErr, "error building player observer 1")

		observer2, observer2BuildErr := observer.NewFactory(observer.PlayerObserver).Build(
			observer.PlayerObserverConfig{
				Name: "washington",
			},
		)

		require.NoError(t, observer2BuildErr, "error building player observer 2")

		state.AddObserver(observer1)
		state.AddObserver(observer2)

		state.NotifyEvent(event.New("hello"))

		require.Len(t, mockSubject.attachCalls, 2)
		require.Len(t, mockSubject.notifyCalls, 1)
	})
}

type MockSubject struct {
	attachCalls []observer.Observer
	notifyCalls []event.Event
}

func (subject *MockSubject) Attach(observer observer.Observer) error {
	subject.attachCalls = append(subject.attachCalls, observer)
	return nil
}

func (subject *MockSubject) Notify(event event.Event) error {
	subject.notifyCalls = append(subject.notifyCalls, event)
	return nil
}
