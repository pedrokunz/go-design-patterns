package game

import (
	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room"
	"github.com/pedrokunz/go-design-patterns/domain/core/player"
	"github.com/pedrokunz/go-design-patterns/event"
	"github.com/pedrokunz/go-design-patterns/event/observer"
)

var state *State = nil

type State struct {
	Player       *player.Player
	Rooms        []room.Room
	Notifier     observer.Notifier
	IsPlayerTurn bool
}

func NewState() *State {
	if state == nil {
		state = &State{
			Rooms:    make([]room.Room, 0),
			Notifier: observer.NewNotifier(),
		}
	}

	return state
}

func (state *State) AddObserver(observer observer.Observer) {
	state.Notifier.Attach(observer)
}

func (state *State) NotifyEvent(event event.Event) {
	state.Notifier.Notify(event)
}
