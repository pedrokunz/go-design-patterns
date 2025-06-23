package create_player

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
)

type Command struct {
	eventStore domain.EventStore
	input      Input
}

type Input struct {
	PlayerName string `json:"player_name"`
}

type Output struct {
	Player *player.Player `json:"player,omitempty"`
	Error  error          `json:"error,omitempty"`
}
