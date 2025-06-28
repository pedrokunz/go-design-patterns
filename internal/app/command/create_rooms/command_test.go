package create_rooms_test

import (
	"errors"
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate"
	"testing"

	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/app/command/create_rooms"
	"github.com/pedrokunz/go-design-patterns/internal/common/test"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	t.Run("should create rooms successfully", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		g, err := game.Create(uuid.NewRandom)
		assert.NoError(t, err)
		gameID := g.Aggregate.ID
		gameEvents := g.Aggregate.Events()
		g.Aggregate.ClearEvents()

		input := create_rooms.Input{
			GameID: gameID,
		}

		mockEventStore.On("Load", gameID).Return(gameEvents, nil)
		mockEventStore.On("Save", gameID, mock.Anything).Return(nil)

		command := create_rooms.NewCommand(mockEventStore, input, uuid.NewRandom)
		output := command.Execute()

		assert.NoError(t, output.Error)
		assert.NotNil(t, output.Game)
		assert.NotEmpty(t, output.Game.Rooms)
		mockEventStore.AssertExpectations(t)
	})

	t.Run("should return error when loading game fails", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		gameID := uuid.New()
		input := create_rooms.Input{
			GameID: gameID,
		}
		expectedErr := errors.New("load game error")

		mockEventStore.On("Load", gameID).Return(nil, expectedErr)

		command := create_rooms.NewCommand(mockEventStore, input, uuid.NewRandom)
		output := command.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Game)
		mockEventStore.AssertExpectations(t)
	})

	t.Run("should return error when creating rooms fails", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		g, err := game.Create(uuid.NewRandom)
		assert.NoError(t, err)
		gameID := g.Aggregate.ID
		gameEvents := g.Aggregate.Events()
		g.Aggregate.ClearEvents()

		input := create_rooms.Input{
			GameID: gameID,
		}
		expectedErr := errors.New("uuid error")
		generator := func() (uuid.UUID, error) {
			return uuid.Nil, expectedErr
		}

		mockEventStore.On("Load", gameID).Return(gameEvents, nil)

		command := create_rooms.NewCommand(mockEventStore, input, generator)
		output := command.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Game)

		mockEventStore.AssertExpectations(t)
	})

	t.Run("should return error when saving game fails", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		g, err := game.Create(uuid.NewRandom)
		assert.NoError(t, err)
		gameID := g.Aggregate.ID
		gameEvents := g.Aggregate.Events()
		g.Aggregate.ClearEvents()

		input := create_rooms.Input{
			GameID: gameID,
		}
		expectedErr := errors.New("save error")

		mockEventStore.On("Load", gameID).Return(gameEvents, nil)
		mockEventStore.On("Save", gameID, mock.Anything).Return(expectedErr)

		command := create_rooms.NewCommand(mockEventStore, input, uuid.NewRandom)
		output := command.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Game)
		mockEventStore.AssertExpectations(t)
	})

	t.Run("should return error when game does not exist", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		gameID := uuid.New()
		input := create_rooms.Input{
			GameID: gameID,
		}
		// Simulate game not found: Load returns empty events and no error
		mockEventStore.On("Load", gameID).Return([]aggregate.Event{}, nil)

		command := create_rooms.NewCommand(mockEventStore, input, uuid.NewRandom)
		output := command.Execute()

		assert.Error(t, output.Error)
		assert.Contains(t, output.Error.Error(), "invalid game state")
		assert.Nil(t, output.Game)
		mockEventStore.AssertExpectations(t)
	})
}
