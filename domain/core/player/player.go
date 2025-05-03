package player

import "github.com/pedrokunz/go-design-patterns/domain/core/internal"

type Player struct {
	Name   string
	Armour internal.Armour
	Life   internal.Life
	Attack internal.Attack
}

const (
	initialArmour = 0
	initialAttack = 1
	initialLife   = 100
)

func New(name string) *Player {
	return &Player{
		Name:   name,
		Armour: internal.Armour{Value: initialArmour},
		Attack: internal.Attack{Value: initialAttack},
		Life:   internal.Life{Value: initialLife},
	}
}
