package add_player_test

import (
	"errors"
	"github.com/pedrokunz/go-design-patterns/internal/app/command/add_player"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"github.com/pedrokunz/go-design-patterns/internal/common/test"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCommand_Execute(t *testing.T) {
	t.Run("should add player to game successfully", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		gameID := common.NewDeterministicUUID("game-id")
		playerID := common.NewDeterministicUUID("player-id")

		g, err := game.Create()
		assert.NoError(t, err)
		gameEvents := g.Aggregate.Events()
		g.Aggregate.ClearEvents()

		p, err := player.Create("player 1")
		assert.NoError(t, err)
		playerEvents := p.Aggregate.Events()
		p.Aggregate.ClearEvents()

		input := add_player.Input{
			GameID:   gameID,
			PlayerID: playerID,
		}

		mockEventStore.On("Load", gameID).Return(gameEvents, nil)
		mockEventStore.On("Load", playerID).Return(playerEvents, nil)
		mockEventStore.On("Save", mock.Anything, mock.Anything).Return(nil)

		cmd := add_player.NewCommand(mockEventStore, input)
		output := cmd.Execute()

		assert.NoError(t, output.Error)
		assert.NotNil(t, output.Game)
		assert.NotNil(t, output.Game.Player)
		assert.Equal(t, "player 1", output.Game.Player.Name)
		mockEventStore.AssertExpectations(t)
	})

	t.Run("should return error when loading game fails", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		input := add_player.Input{
			GameID:   common.NewDeterministicUUID("game-id"),
			PlayerID: common.NewDeterministicUUID("player-id"),
		}
		expectedErr := errors.New("load game error")

		mockEventStore.On("Load", input.GameID).Return(nil, expectedErr)

		cmd := add_player.NewCommand(mockEventStore, input)
		output := cmd.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Game)
		mockEventStore.AssertExpectations(t)
	})

	t.Run("should return error when loading player fails", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		gameID := common.NewDeterministicUUID("game-id")
		playerID := common.NewDeterministicUUID("player-id")

		g, err := game.Create()
		assert.NoError(t, err)
		gameEvents := g.Aggregate.Events()

		input := add_player.Input{
			GameID:   gameID,
			PlayerID: playerID,
		}
		expectedErr := errors.New("load player error")

		mockEventStore.On("Load", gameID).Return(gameEvents, nil)
		mockEventStore.On("Load", playerID).Return(nil, expectedErr)

		cmd := add_player.NewCommand(mockEventStore, input)
		output := cmd.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Game)
		mockEventStore.AssertExpectations(t)
	})

	t.Run("should return error when adding player to game fails", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		gameID := common.NewDeterministicUUID("game-id")
		playerID := common.NewDeterministicUUID("player-id")

		g, err := game.Create()
		assert.NoError(t, err)
		p1, err := player.Create("player-1")
		assert.NoError(t, err)
		err = g.AddPlayer(p1)
		assert.NoError(t, err)
		gameEventsWithPlayer := g.Aggregate.Events()

		p2, err := player.Create("player-2")
		assert.NoError(t, err)
		playerEvents := p2.Aggregate.Events()

		input := add_player.Input{
			GameID:   gameID,
			PlayerID: playerID,
		}

		mockEventStore.On("Load", gameID).Return(gameEventsWithPlayer, nil)
		mockEventStore.On("Load", playerID).Return(playerEvents, nil)

		cmd := add_player.NewCommand(mockEventStore, input)
		output := cmd.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, "game already has a player", output.Error.Error())
		assert.Nil(t, output.Game)
		mockEventStore.AssertExpectations(t)
	})

	t.Run("should return error when saving game fails", func(t *testing.T) {
		mockEventStore := &test.MockEventStore{}
		gameID := common.NewDeterministicUUID("game-id")
		playerID := common.NewDeterministicUUID("player-id")

		g, err := game.Create()
		assert.NoError(t, err)
		gameEvents := g.Aggregate.Events()

		p, err := player.Create("player 1")
		assert.NoError(t, err)
		playerEvents := p.Aggregate.Events()

		input := add_player.Input{
			GameID:   gameID,
			PlayerID: playerID,
		}
		expectedErr := errors.New("save error")

		mockEventStore.On("Load", gameID).Return(gameEvents, nil)
		mockEventStore.On("Load", playerID).Return(playerEvents, nil)
		mockEventStore.On("Save", mock.Anything, mock.Anything).Return(expectedErr)

		cmd := add_player.NewCommand(mockEventStore, input)
		output := cmd.Execute()

		assert.Error(t, output.Error)
		assert.Equal(t, expectedErr, output.Error)
		assert.Nil(t, output.Game)
		mockEventStore.AssertExpectations(t)
	})
}
