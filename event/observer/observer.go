package observer

import (
	"github.com/pedrokunz/go-design-patterns/event"
)

type Observer interface {
	On(event event.Event) error
}

type Kind string

func New(kind Kind, config any) (Observer, error) {
	return NewFactory(kind).Build(config)
}

const (
	PlayerObserver Kind = "player"
)
