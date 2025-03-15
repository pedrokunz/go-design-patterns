package event

type Kind string

type Event interface {
	Type() Kind
}

type GameEvent struct {
	kind Kind
}

func (gameEvent *GameEvent) Type() Kind {
	return gameEvent.kind
}
