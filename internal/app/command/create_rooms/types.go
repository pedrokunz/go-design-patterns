package create_rooms

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
	"github.com/pedrokunz/go-design-patterns/internal/domain/room"
)

type Command struct {
	eventStore    domain.EventStore
	input         Input
	uuidGenerator common.UUIDGenerator
}

type Input struct {
	GameID  uuid.UUID     `json:"game_id"`
	Configs []room.Config `json:"configs"`
}

type Output struct {
	Game  *game.Game `json:"game,omitempty"`
	Error error      `json:"error,omitempty"`
}
