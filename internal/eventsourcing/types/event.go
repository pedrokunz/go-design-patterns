package types

import "sync"

type EventType string

// validEventTypes holds all registered event types
var validEventTypes = sync.Map{}

// RegisterEventType registers an event type as valid
func RegisterEventType(eventType EventType) {
	validEventTypes.Store(eventType, true)
}

func (e EventType) IsValid() bool {
	if e == "" {
		return false
	}

	_, isValid := validEventTypes.Load(e)
	return isValid
}

func EventTypes() []EventType {
	eventTypes := make([]EventType, 0)

	sync.OnceFunc(func() {
		validEventTypes.Range(func(key, value any) bool {
			eventType, found := key.(EventType)
			if found {
				eventTypes = append(eventTypes, eventType)
			}

			return found
		})
	})()

	return eventTypes
}
