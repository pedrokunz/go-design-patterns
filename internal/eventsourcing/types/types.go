package types

type EventType string

func (e EventType) IsValid() bool {
	if e == "" {
		return false
	}

	return true
}

type AggregateType string

func (a AggregateType) IsValid() bool {
	if a == "" {
		return false
	}

	return true
}

const (
	ErrInvalidAggregateID      = "invalid aggregate ID"
	ErrInvalidAggregateType    = "invalid aggregate type"
	ErrInvalidAggregateVersion = "invalid aggregate version"

	ErrInvalidEventID         = "invalid event ID"
	ErrInvalidEventPayload    = "invalid event payload"
	ErrInvalidEventRecordedAt = "invalid event recorded at"
	ErrInvalidEventType       = "invalid event type"
	ErrInvalidCausationID     = "invalid causation ID"
)
