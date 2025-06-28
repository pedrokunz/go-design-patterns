package create_player

import (
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
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
	Player, createPlayerErr := player.Create(c.uuidGenerator, c.input.PlayerName)
	if createPlayerErr != nil {
		return Output{
			Player: nil,
			Error:  createPlayerErr,
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
