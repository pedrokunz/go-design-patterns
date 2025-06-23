package create_player

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
)

func NewCommand(eventStore domain.EventStore, input Input) *Command {
	return &Command{
		eventStore: eventStore,
		input:      input,
	}
}

func (c *Command) Execute() Output {
	Player, createErr := player.Create(c.input.PlayerName)
	if createErr != nil {
		return Output{
			Player: nil,
			Error:  createErr,
		}
	}

	saveErr := c.eventStore.Save(Player.Aggregate.ID, Player.Aggregate.Events())
	if saveErr != nil {
		return Output{
			Player: nil,
			Error:  saveErr,
		}
	}

	return Output{
		Player: Player,
		Error:  nil,
	}
}
