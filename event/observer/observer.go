package observer

import (
	"github.com/pedrokunz/go-design-patterns/event"
)

type Observer interface {
	On(event event.Event) error
}

func New(kind Kind, config any) (Observer, error) {
	return newFactory(kind).Build(config)
}

type Kind string

const (
	PlayerObserver Kind = "player"
)
