package observer

import (
	"fmt"
	"github.com/pedrokunz/go-design-patterns/event"
)

type PlayerObserverConfig struct {
	Name string
}

type playerObserver struct {
	name string
}

func newPlayerObserver(config PlayerObserverConfig) (Observer, error) {
	return &playerObserver{name: config.Name}, nil
}

func (p *playerObserver) On(event event.Event) error {
	if event == nil {
		return fmt.Errorf("event cannot be nil")
	}

	fmt.Printf("Player %s received event %s\n", p.name, event.Type())

	return nil
}
