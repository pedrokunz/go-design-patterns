package observer

import (
	"errors"
	"fmt"
)

type Factory struct {
	kind Kind
}

func newFactory(kind Kind) *Factory {
	return &Factory{kind: kind}
}

var invalidConfigType = fmt.Errorf("invalid config type")

func (factory *Factory) Build(config any) (o Observer, _ error) {
	switch factory.kind {
	case PlayerObserver:
		return newPlayerObserver(config)
	default:
		return o, errors.New("invalid observer kind")
	}
}
