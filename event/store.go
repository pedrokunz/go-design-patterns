package event

type Store struct {
	events []Event
}

func NewEventStore() *Store {
	return &Store{
		events: make([]Event, 0),
	}
}

func (store *Store) Add(event Event) {
	store.events = append(store.events, event)
}

func (store *Store) Events() []Event {
	return store.events
}
