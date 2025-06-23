package event

import "sync"

type Type string

// validEventTypes holds all registered event types
var validEventTypes = sync.Map{}

// RegisterType registers an event type as valid
func RegisterType(eventType Type) {
	validEventTypes.Store(eventType, true)
}

func (e Type) IsValid() bool {
	if e == "" {
		return false
	}

	_, isValid := validEventTypes.Load(e)
	return isValid
}

func Types() []Type {
	eventTypes := make([]Type, 0)

	validEventTypes.Range(func(key, value any) bool {
		eventType, found := key.(Type)
		if found {
			eventTypes = append(eventTypes, eventType)
		}

		return found
	})

	return eventTypes
}
