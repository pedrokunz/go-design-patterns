package player

import "github.com/pedrokunz/go-design-patterns/internal/eventsourcing/types"

const (
	Aggregate types.AggregateType = "player"

	Created types.EventType = "player.created"
)

func init() {
	types.RegisterAggregateType(Aggregate)
	types.RegisterEventType(Created)
}
