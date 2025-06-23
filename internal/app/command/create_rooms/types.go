package create_rooms

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
)

type Command struct {
	eventStore domain.EventStore
	input      Input
}

type Input struct {
	GameID uuid.UUID `json:"game_id"`
}

type Output struct {
	Game  *game.Game `json:"game,omitempty"`
	Error error      `json:"error,omitempty"`
}
