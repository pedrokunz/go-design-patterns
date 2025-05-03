package enemy

import "github.com/pedrokunz/go-design-patterns/domain/core/internal"

type Enemy struct {
	Type   Kind
	Armour internal.Armour
	Life   internal.Life
	Attack internal.Attack
}

func New(t Kind) *Enemy {
	return &Enemy{
		Type:   t,
		Armour: internal.Armour{Value: 0},
		Attack: internal.Attack{Value: 1},
		Life:   internal.Life{Value: 100},
	}
}
