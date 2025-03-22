package observer_test

import (
	"github.com/pedrokunz/go-design-patterns/event"
)

type MockObserver struct {
	onUpdate func(event event.Event) error
	events   []event.Event
	onCalls  int
}

func NewMockObserver() *MockObserver {
	return &MockObserver{
		events: make([]event.Event, 0),
	}
}

func (observer *MockObserver) On(event event.Event) error {
	observer.onCalls++
	if observer.onUpdate != nil {
		return observer.onUpdate(event)
	}

	observer.events = append(observer.events, event)

	return nil
}
