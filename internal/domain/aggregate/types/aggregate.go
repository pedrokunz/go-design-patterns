package types

import "sync"

type AggregateType string

// validTypes holds all registered aggregate types
var validTypes = sync.Map{}

// RegisterAggregateType registers an aggregate type as valid
func RegisterAggregateType(aggregateType AggregateType) {
	validTypes.Store(aggregateType, true)
}

func (a AggregateType) IsValid() bool {
	if a == "" {
		return false
	}

	_, isValid := validTypes.Load(a)
	return isValid
}

func AggregateTypes() []AggregateType {
	types := make([]AggregateType, 0)

	validTypes.Range(func(key, value any) bool {
		aggregateType, found := key.(AggregateType)
		if found {
			types = append(types, aggregateType)
		}

		return found
	})

	return types
}
