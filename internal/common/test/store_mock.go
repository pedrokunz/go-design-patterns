package test

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate"
	"github.com/stretchr/testify/mock"
)

type MockEventStore struct {
	mock.Mock
}

func (m *MockEventStore) Clear() {
	m.Called()
}

func (m *MockEventStore) GetEvents() map[uuid.UUID][]aggregate.Event {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(map[uuid.UUID][]aggregate.Event)
}

func (m *MockEventStore) Load(id uuid.UUID) ([]aggregate.Event, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]aggregate.Event), args.Error(1)
}

func (m *MockEventStore) Save(id uuid.UUID, events []aggregate.Event) error {
	args := m.Called(id, events)
	return args.Error(0)
}
