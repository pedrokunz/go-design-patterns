package event

type Kind string

type Event interface {
	Type() Kind
}

