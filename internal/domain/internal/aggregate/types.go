package aggregate

import "sync"

type Type string

// validTypes holds all registered aggregate types
var validTypes = sync.Map{}

// RegisterType registers an aggregate type as valid
func RegisterType(aggregateType Type) {
	validTypes.Store(aggregateType, true)
}

func (a Type) IsValid() bool {
	if a == "" {
		return false
	}

	_, isValid := validTypes.Load(a)
	return isValid
}

func Types() []Type {
	types := make([]Type, 0)

	validTypes.Range(func(key, value any) bool {
		aggregateType, found := key.(Type)
		if found {
			types = append(types, aggregateType)
		}

		return found
	})

	return types
}
