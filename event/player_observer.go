package event

import "fmt"

type playerObserver struct {
	name string
}

func NewPlayerObserver(name string) Observer {
	return &playerObserver{name: name}
}

func (p *playerObserver) Update(event Event) {
	fmt.Printf("Player %s received event %s", p.name, event.Type())
}
