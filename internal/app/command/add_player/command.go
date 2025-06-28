package add_player

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
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

	Player, loadPlayerErr := c.loadPlayer()
	if loadPlayerErr != nil {
		return Output{
			Game:  nil,
			Error: loadPlayerErr,
		}
	}

	addPlayerErr := Game.AddPlayer(Player)
	if addPlayerErr != nil {
		return Output{
			Game:  nil,
			Error: addPlayerErr,
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

func (c *Command) loadPlayer() (*player.Player, error) {
	events, loadPlayerErr := c.eventStore.Load(c.input.PlayerID)
	if loadPlayerErr != nil {
		return nil, loadPlayerErr
	}

	return player.LoadFromHistory(events)
}
