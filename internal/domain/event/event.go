package event

type Kind string

type Event interface {
	Type() Kind
}

type event struct {
	kind Kind
}

func New(kind Kind) Event {
	return &event{kind: kind}
}

func (event *event) Type() Kind {
	return event.kind
}

const (
	PlayerJoined Kind = "player_joined"
	PlayerLeft   Kind = "player_left"
)
