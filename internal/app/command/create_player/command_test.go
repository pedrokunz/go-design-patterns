package create_player_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/app/command/create_player"
	"github.com/pedrokunz/go-design-patterns/internal/common/test"
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
			uuid.NewRandom,
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

	t.Run("should return error when player creation fails", func(t *testing.T) {
		expectedErr := errors.New("uuid error")
		generator := func() (uuid.UUID, error) {
			return uuid.Nil, expectedErr
		}
		eventStore := domain.NewEventStore()
		command := create_player.NewCommand(
			eventStore,
			create_player.Input{
				PlayerName: "Player1",
			},
			generator,
		)

		output := command.Execute()

		require.Error(t, output.Error, "Execute should return an error")
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Player)
	})

	t.Run("should return error when saving player fails", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		expectedErr := errors.New("save error")
		mockEventStore.On("Save", mock.Anything, mock.AnythingOfType("[]aggregate.Event")).Return(expectedErr)

		command := create_player.NewCommand(
			mockEventStore,
			create_player.Input{
				PlayerName: "Player1",
			},
			uuid.NewRandom,
		)

		output := command.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Player)

		mockEventStore.AssertExpectations(t)
	})
}
