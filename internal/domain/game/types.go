package game

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/types"
)

const (
	Aggregate types.AggregateType = "game"

	Created      types.EventType = "game.created"
	RoomsCreated types.EventType = "game.room.created"
	PlayerAdded  types.EventType = "game.player.added"
)

func init() {
	types.RegisterAggregateType(Aggregate)
	types.RegisterEventType(Created)
	types.RegisterEventType(RoomsCreated)
	types.RegisterEventType(PlayerAdded)
}
