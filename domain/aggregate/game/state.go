package game

import (
	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room"
	"github.com/pedrokunz/go-design-patterns/domain/core/player"
)

var state *State = nil

type State struct {
	Player player.Player
	Rooms  []room.Room
}

func NewState() *State {
	if state == nil {
		state = &State{}
	}

	return state
}
