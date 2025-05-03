package player

import "github.com/pedrokunz/go-design-patterns/domain/core/internal"

type Player struct {
	Name   string
	Armour internal.Armour
	Life   internal.Life
	Attack internal.Attack
}

func New(name string) *Player {
	return &Player{
		Name:   name,
		Armour: internal.Armour{Value: 0},
		Attack: internal.Attack{Min: 1, Max: 100},
		Life:   internal.Life{Value: 100},
	}
}
