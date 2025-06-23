package game

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal/aggregate"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal/event"
)

const (
	Aggregate aggregate.Type = "game"

	Created      event.Type = "game.created"
	RoomsCreated event.Type = "game.room.created"
	PlayerAdded  event.Type = "game.player.added"
)

func init() {
	aggregate.RegisterType(Aggregate)
	event.RegisterType(Created)
	event.RegisterType(RoomsCreated)
	event.RegisterType(PlayerAdded)
}
