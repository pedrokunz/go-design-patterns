package game

import (
	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room"
	"github.com/pedrokunz/go-design-patterns/domain/core/player"
	"github.com/pedrokunz/go-design-patterns/event"
)

var state *State = nil

type State struct {
	Player  player.Player
	Rooms   []room.Room
	Subject event.Notifier
}

func NewState() *State {
	if state == nil {
		state = &State{Subject: &event.Subject{}}
	}

	return state
}

func (state *State) AddObserver(observer event.Observer) {
	state.Subject.Attach(observer)
}

func (state *State) NotifyEvent(event event.Event) {
	state.Subject.Notify(event)
}
