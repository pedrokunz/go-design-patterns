package types

import "sync"

type AggregateType string

// validAggregateTypes holds all registered aggregate types
var validAggregateTypes = sync.Map{}

// RegisterAggregateType registers an aggregate type as valid
func RegisterAggregateType(aggregateType AggregateType) {
	validAggregateTypes.Store(aggregateType, true)
}

func (a AggregateType) IsValid() bool {
	if a == "" {
		return false
	}

	_, isValid := validAggregateTypes.Load(a)
	return isValid
}

func AggregateTypes() []AggregateType {
	aggregateTypes := make([]AggregateType, 0)

	sync.OnceFunc(func() {
		validAggregateTypes.Range(func(key, value any) bool {
			aggregateType, found := key.(AggregateType)
			if found {
				aggregateTypes = append(aggregateTypes, aggregateType)
			}

			return found
		})
	})()

	return aggregateTypes
}
