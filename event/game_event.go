package event

type gameEvent struct {
	kind Kind
}

func NewGameEvent(kind Kind) Event {
	return &gameEvent{kind: kind}
}

func (gameEvent *gameEvent) Type() Kind {
	return gameEvent.kind
}

