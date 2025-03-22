package observer

import (
	"errors"
	"fmt"
)

type Factory struct {
	kind Kind
}

func NewFactory(kind Kind) *Factory {
	return &Factory{kind: kind}
}

func (factory *Factory) Build(config any) (o Observer, _ error) {
	invalidConfigType := fmt.Errorf("invalid config type")

	switch factory.kind {
	case PlayerObserver:
		playerObserverConfig, ok := config.(PlayerObserverConfig)
		if !ok {
			return nil, invalidConfigType
		}

		return newPlayerObserver(playerObserverConfig)
	default:
		return o, errors.New("invalid observer type")
	}
}
