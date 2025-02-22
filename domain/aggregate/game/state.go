package game

import (
	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room"
	"github.com/pedrokunz/go-design-patterns/domain/core/player"
)

type State struct {
	Player player.Player
	Rooms  []room.Room
}
