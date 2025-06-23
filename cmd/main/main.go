package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/app/command/add_player"
	"github.com/pedrokunz/go-design-patterns/internal/app/command/create_game"
	"github.com/pedrokunz/go-design-patterns/internal/app/command/create_player"
	"github.com/pedrokunz/go-design-patterns/internal/app/command/create_rooms"
	"github.com/pedrokunz/go-design-patterns/internal/domain"
	"github.com/pedrokunz/go-design-patterns/internal/domain/game"
	"github.com/pedrokunz/go-design-patterns/internal/domain/player"
	"os"
)

var eventStore = domain.NewEventStore()

func main() {
	fmt.Println("Hello player, what is your name?")

	scanner := bufio.NewScanner(os.Stdin)
	playerName := ""
	if scanner.Scan() {
		playerName = scanner.Text()
		fmt.Printf("Nice to meet you, %s!\n", playerName)
	}

	err := scanner.Err()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Reading standard input: %v\n", "test")
	}

	Game := createGame()
	Game = createRooms(Game.Aggregate.ID)
	Player := createPlayer(playerName)
	Game = addPlayer(Game.Aggregate.ID, Player.Aggregate.ID)

	fmt.Println("Initiate combat!")

	Enemy := Game.Rooms[1].Enemies[0]
	for Enemy.Life.Value > 0 {
		if Game.IsPlayerTurn {
			damage := Enemy.TakeDamage(Game.Player.Attack)
			Game.IsPlayerTurn = false

			fmt.Printf("👺 Enemy took %d damage ♥️[%d]\n", damage, Enemy.Life.Value)

			if Enemy.Life.Value <= 0 {
				fmt.Println("Enemy died! ☠️")
				break
			}
		} else {
			damage := Game.Player.TakeDamage(Enemy.Attack)
			Game.IsPlayerTurn = true

			fmt.Printf("🤺 Player took %d damage ♥️[%d]\n", damage, Game.Player.Life.Value)

			if Game.Player.Life.Value <= 0 {
				fmt.Println("Player died! ☠️")
				break
			}
		}
	}

	fmt.Println("Game over!")
}

func createGame() *game.Game {
	command := create_game.NewCommand(eventStore, create_game.Input{})

	output := command.Execute()
	if output.Error != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Creating game: %v\n", output.Error)
	}

	return output.Game
}

func createRooms(gameID uuid.UUID) *game.Game {
	command := create_rooms.NewCommand(eventStore, create_rooms.Input{
		GameID: gameID,
	})

	output := command.Execute()
	if output.Error != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Creating rooms: %v\n", output.Error)
		return nil
	}

	return output.Game
}

func createPlayer(playerName string) *player.Player {
	command := create_player.NewCommand(
		eventStore,
		create_player.Input{
			PlayerName: playerName,
		},
	)

	output := command.Execute()
	if output.Error != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Creating player: %v\n", output.Error)
	}

	return output.Player
}

func addPlayer(gameID, playerID uuid.UUID) *game.Game {
	command := add_player.NewCommand(
		eventStore,
		add_player.Input{
			GameID:   gameID,
			PlayerID: playerID,
		},
	)

	output := command.Execute()
	if output.Error != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Adding player: %v\n", output.Error)
	}

	return output.Game
}
