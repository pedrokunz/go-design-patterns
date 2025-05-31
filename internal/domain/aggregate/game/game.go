package game

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/player"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/room"
	"github.com/pedrokunz/go-design-patterns/internal/domain/event"
	"github.com/pedrokunz/go-design-patterns/internal/domain/event/observer"
)

var game *Game = nil

type Game struct {
	Player       *player.Player
	Rooms        []room.Room
	Notifier     observer.Notifier
	IsPlayerTurn bool
}

func NewGame() *Game {
	if game == nil {
		game = &Game{
			Rooms:    make([]room.Room, 0),
			Notifier: observer.NewNotifier(),
		}
	}

	return game
}

func (state *Game) AddObserver(observer observer.Observer) {
	_ = state.Notifier.Attach(observer)
}

func (state *Game) NotifyEvent(event event.Event) {
	_ = state.Notifier.Notify(event)
}
