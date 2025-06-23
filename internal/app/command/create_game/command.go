package create_game

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
)

func NewCommand(eventStore domain.EventStore, input Input) *Command {
	return &Command{
		eventStore: eventStore,
		input:      input,
	}
}

func (c *Command) Execute() Output {
	Game, createErr := game.Create()
	if createErr != nil {
		return Output{
			Game:  nil,
			Error: createErr,
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
