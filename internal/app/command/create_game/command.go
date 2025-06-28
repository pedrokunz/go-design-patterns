package create_game

import (
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
)

func NewCommand(
	eventStore domain.EventStore,
	input Input,
	uuidGenerator common.UUIDGenerator,
) *Command {
	return &Command{
		eventStore:    eventStore,
		input:         input,
		uuidGenerator: uuidGenerator,
	}
}

func (c *Command) Execute() Output {
	Game, createGameErr := game.Create(c.uuidGenerator)
	if createGameErr != nil {
		return Output{
			Game:  nil,
			Error: createGameErr,
		}
	}

	saveErr := c.eventStore.Save(Game.Aggregate.ID, Game.Aggregate.Events())
	if saveErr != nil {
		return Output{
			Game:  nil,
			Error: saveErr,
		}
	}

	return Output{
		Game:  Game,
		Error: nil,
	}
}
