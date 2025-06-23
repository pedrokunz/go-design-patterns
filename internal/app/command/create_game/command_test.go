package create_game_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/app/command/create_game"
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExecute(t *testing.T) {
	t.Run("returns a valid game", func(t *testing.T) {
		eventStore := domain.NewEventStore()
		command := create_game.NewCommand(eventStore, create_game.Input{})

		output := command.Execute()
		if output.Error != nil {
			t.Errorf("Execute should not return an error, got: %v", output.Error)
		}

		require.NotNil(t, output.Game, "Game should not be nil")

		require.Equal(
			t,
			1,
			len(eventStore.GetEvents()[output.Game.Aggregate.ID]),
			"Should have one event in the store",
		)
	})
}
