package game_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/domain/aggregate/game"
	"github.com/pedrokunz/go-design-patterns/event"
	"github.com/stretchr/testify/require"
)

func TestNewState(t *testing.T) {
	t.Run("returns a valid state", func(t *testing.T) {
		actual := game.NewState()

		require.NotNil(t, actual, "NewState should not be nil")
	})

	t.Run("returns the same state", func(t *testing.T) {
		actual := game.NewState()
		expected := game.NewState()

		require.True(
			t,
			expected == actual,
			"NewState should return the same state",
		)
	})

	t.Run("notifies multiple observers of events", func(t *testing.T) {
		state := game.NewState()
		mockSubject := &MockSubject{}
		state.Subject = mockSubject

		observer1 := event.NewPlayerObserver("george")
		observer2 := event.NewPlayerObserver("washington")
		state.AddObserver(observer1)
		state.AddObserver(observer2)

		state.NotifyEvent(event.NewGameEvent("hello"))

		require.Len(t, mockSubject.attachCalls, 2)
		require.Len(t, mockSubject.notifyCalls, 1)
	})
}

type MockSubject struct {
	attachCalls []event.Observer
	notifyCalls []event.Event
}

func (subject *MockSubject) Attach(observer event.Observer) {
	subject.attachCalls = append(subject.attachCalls, observer)
}

func (subject *MockSubject) Notify(event event.Event) {
	subject.notifyCalls = append(subject.notifyCalls, event)
}
