package observer

import (
	"fmt"
	"github.com/pedrokunz/go-design-patterns/internal/domain/event"
)

type PlayerObserverConfig struct {
	Name string
}

type playerObserver struct {
	name string
}

func newPlayerObserver(config any) (Observer, error) {
	playerObserverConfig, ok := config.(PlayerObserverConfig)
	if !ok {
		return nil, invalidConfigType
	}

	return &playerObserver{name: playerObserverConfig.Name}, nil
}

func (p *playerObserver) On(event event.Event) error {
	if event == nil {
		return fmt.Errorf("event cannot be nil")
	}

	fmt.Printf("Player %s received event %s\n", p.name, event.Type())

	return nil
}
