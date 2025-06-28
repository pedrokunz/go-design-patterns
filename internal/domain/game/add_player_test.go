package game_test

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
	"testing"

	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"github.com/stretchr/testify/assert"
)

func TestGame_AddPlayer(t *testing.T) {
	t.Run("should return error when player is nil", func(t *testing.T) {
		g := &game.Game{}
		err := g.AddPlayer(nil)
		assert.Error(t, err)
		assert.Equal(t, "player cannot be nil", err.Error())
	})

	t.Run("should return error when game is nil", func(t *testing.T) {
		var g *game.Game
		err := g.AddPlayer(&player.Player{})
		assert.Error(t, err)
		assert.Equal(t, "game cannot be nil", err.Error())
	})

	t.Run("should return error when game already has a player", func(t *testing.T) {
		g := &game.Game{
			Player: &player.Player{},
		}
		err := g.AddPlayer(&player.Player{})
		assert.Error(t, err)
		assert.Equal(t, "game already has a player", err.Error())
	})

	t.Run("should add player successfully", func(t *testing.T) {
		g, err := game.Create(uuid.Nil)
		assert.NoError(t, err)
		p, err := player.Create("player 1")
		assert.NoError(t, err)
		err = g.AddPlayer(p)
		assert.NoError(t, err)
		assert.NotNil(t, g.Player)
		assert.Equal(t, "player 1", g.Player.Name)
		assert.NotEmpty(t, g.Aggregate.Events)
		assert.Equal(t, game.PlayerAdded, g.Aggregate.Events()[1].Type())
	})
}
