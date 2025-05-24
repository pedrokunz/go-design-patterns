package game

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/player"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/room"
	"github.com/pedrokunz/go-design-patterns/internal/domain/event"
	"github.com/pedrokunz/go-design-patterns/internal/domain/event/observer"
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
	_ = state.Notifier.Attach(observer)
}

func (state *State) NotifyEvent(event event.Event) {
	_ = state.Notifier.Notify(event)
}
