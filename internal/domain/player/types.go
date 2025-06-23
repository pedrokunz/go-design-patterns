package player

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal/aggregate"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal/event"
)

const (
	Aggregate aggregate.Type = "player"

	Created event.Type = "player.created"
)

func init() {
	aggregate.RegisterType(Aggregate)
	event.RegisterType(Created)
}
