package create_rooms

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
	Game, loadGameErr := c.loadGame()
	if loadGameErr != nil {
		return Output{
			Game:  nil,
			Error: loadGameErr,
		}
	}

	err := Game.CreateRooms()
	if err != nil {
		return Output{
			Game:  nil,
			Error: err,
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

func (c *Command) loadGame() (*game.Game, error) {
	events, loadGameErr := c.eventStore.Load(c.input.GameID)
	if loadGameErr != nil {
		return nil, loadGameErr
	}

	return game.LoadFromHistory(events)
}
