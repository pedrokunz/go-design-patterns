package create_player_test

import (
	"github.com/pedrokunz/go-design-patterns/internal/app/command/create_player"
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExecute(t *testing.T) {
	t.Run("returns a valid player", func(t *testing.T) {
		eventStore := domain.NewEventStore()
		command := create_player.NewCommand(
			eventStore,
			create_player.Input{
				PlayerName: "Player1",
			},
		)

		output := command.Execute()
		if output.Error != nil {
			t.Errorf("Execute should not return an error, got: %v", output.Error)
		}

		require.NotNil(t, output.Player, "Player should not be nil")

		require.Equal(
			t,
			1,
			len(eventStore.GetEvents()[output.Player.Aggregate.ID]),
			"Should have one event in the store",
		)
	})
}
