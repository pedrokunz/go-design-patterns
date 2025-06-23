package create_game

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
)

type Command struct {
	eventStore domain.EventStore
	input      Input
}

type Input struct{}

type Output struct {
	Game  *game.Game `json:"game,omitempty"`
	Error error      `json:"error,omitempty"`
}
