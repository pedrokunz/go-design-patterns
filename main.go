package main

import (
	"bufio"
	"fmt"
	"github.com/pedrokunz/go-design-patterns/domain/aggregate/game"
	"github.com/pedrokunz/go-design-patterns/domain/aggregate/room"
	"github.com/pedrokunz/go-design-patterns/domain/core/enemy"
	"github.com/pedrokunz/go-design-patterns/domain/core/item"
	"github.com/pedrokunz/go-design-patterns/domain/core/player"
	"os"
)

func main() {
	fmt.Println("Hello player, what is your name?")

	scanner := bufio.NewScanner(os.Stdin)
	name := ""
	if scanner.Scan() {
		name = scanner.Text()
		fmt.Printf("Nice to meet you, %s!\n", name)
	}

	err := scanner.Err()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Reading standard input: %v\n", "test")
	}

	state := game.NewState()
	Player := player.New(name)

	state.Player = Player

	treasuryRoom := room.Factory(
		room.FactoryInput{
			Kind: room.KindTreasure,
			Items: []item.Item{
				{
					Name: "Sword",
					Type: item.Weapon,
				},
				{
					Name: "Shield",
					Type: item.Armour,
				},
			},
		},
	)

	enemyRoom := room.Factory(
		room.FactoryInput{
			Kind: room.KindEnemy,
			Enemies: []*enemy.Enemy{
				enemy.New(enemy.Goblin),
			},
		},
	)

	state.Rooms = []room.Room{
		treasuryRoom,
		enemyRoom,
	}

	fmt.Println("Initiate combat!")

	Enemy := state.Rooms[1].Enemies()[0]
	for Enemy.Life.Value > 0 {
		if state.IsPlayerTurn {
			damage := Enemy.TakeDamage(Player.Attack)
			state.IsPlayerTurn = false

			if Enemy.Life.Value <= 0 {
				fmt.Println("Enemy died! ☠️")
				break
			} else {
				fmt.Printf("👺 Enemy took %d damage ♥️[%d]\n", damage, Enemy.Life.Value)
			}
		} else {
			damage := Player.TakeDamage(Enemy.Attack)
			state.IsPlayerTurn = true

			if Player.Life.Value <= 0 {
				fmt.Println("Player died! ☠️")
				break
			} else {
				fmt.Printf("🤺 Player took %d damage ♥️[%d]\n", damage, Player.Life.Value)
			}
		}
	}

	fmt.Println("Game over!")
}
