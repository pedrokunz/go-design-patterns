package event

type Bus struct {
	subscribers map[Kind][]func(Event)
}

func NewEventBus() *Bus {
	return &Bus{
		subscribers: make(map[Kind][]func(Event)),
	}
}

func (bus *Bus) Subscribe(kind Kind, subscriber func(Event)) {
	bus.subscribers[kind] = append(bus.subscribers[kind], subscriber)
}

func (bus *Bus) Publish(event Event) {
	handlers, found := bus.subscribers[event.Type()]
	if !found {
		return
	}

	for _, handler := range handlers {
		handler(event)
	}
}
