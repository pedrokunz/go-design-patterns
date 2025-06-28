package domain

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate"
)

type EventStore interface {
	Load(id uuid.UUID) ([]aggregate.Event, error)
	Save(id uuid.UUID, events []aggregate.Event) error
	Clear()
	GetEvents() map[uuid.UUID][]aggregate.Event
}

type store struct {
	events map[uuid.UUID][]aggregate.Event
}

func NewEventStore() EventStore {
	return &store{
		events: make(map[uuid.UUID][]aggregate.Event),
	}
}

func (s *store) Load(id uuid.UUID) ([]aggregate.Event, error) {
	if events, exists := s.events[id]; exists {
		return events, nil
	}

	return nil, errors.New("events not found")
}

func (s *store) Save(id uuid.UUID, events []aggregate.Event) error {
	for _, event := range events {
		if event == nil {
			return errors.New("event cannot be nil")
		}
	}

	s.events[id] = events

	return nil
}

func (s *store) Clear() {
	s.events = make(map[uuid.UUID][]aggregate.Event)
}

func (s *store) GetEvents() map[uuid.UUID][]aggregate.Event {
	return s.events
}
