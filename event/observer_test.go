package event_test

import (
	"testing"

	"github.com/pedrokunz/go-design-patterns/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubject(t *testing.T) {
	t.Run("attaches and notifies observers", func(t *testing.T) {
		subject := &event.Subject{}
		observer := &MockObserver{}
		subject.Attach(observer)

		event := &ExampleEvent{}
		subject.Notify(event)

		assert.Equal(t, observer.calls, 1, "observer must be called only once")
		require.Len(t, observer.events, 1, "observer must have only one event")
		require.Equal(
			t,
			observer.events[0].Type(),
			event.Type(),
			"observer must be called with 'hello' event type",
		)
	})
}

type ExampleEvent struct{}

func (event *ExampleEvent) Type() string {
	return "hello"
}

type MockObserver struct {
	onUpdate func(event event.Event)
	events   []event.Event
	calls    int
}

func (observer *MockObserver) Update(event event.Event) {
	observer.calls++
	if observer.onUpdate != nil {
		observer.onUpdate(event)
		return
	}

	observer.events = append(observer.events, event)
}
