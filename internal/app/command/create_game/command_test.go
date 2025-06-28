package create_game_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/app/command/create_game"
	"github.com/pedrokunz/go-design-patterns/internal/common/test"
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExecute(t *testing.T) {
	t.Run("returns a valid game", func(t *testing.T) {
		eventStore := domain.NewEventStore()
		command := create_game.NewCommand(eventStore, create_game.Input{}, uuid.NewRandom)

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

	t.Run("should return error when game creation fails", func(t *testing.T) {
		expectedErr := errors.New("uuid error")
		generator := func() (uuid.UUID, error) {
			return uuid.Nil, expectedErr
		}
		eventStore := domain.NewEventStore()
		command := create_game.NewCommand(eventStore, create_game.Input{}, generator)

		output := command.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Game)
	})

	t.Run("should return error when saving game fails", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		expectedErr := errors.New("save error")
		mockEventStore.On("Save", mock.Anything, mock.AnythingOfType("[]aggregate.Event")).Return(expectedErr)

		command := create_game.NewCommand(mockEventStore, create_game.Input{}, nil)
		output := command.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Game)

		mockEventStore.AssertExpectations(t)
	})
}
