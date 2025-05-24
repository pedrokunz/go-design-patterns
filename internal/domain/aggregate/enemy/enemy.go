package enemy

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/aggregate/internal"
	"math/rand"
)

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
		Attack: internal.Attack{Min: 1, Max: 100},
		Life:   internal.Life{Value: 100},
	}
}

func (e *Enemy) TakeDamage(attack internal.Attack) int {
	damage := rand.Intn(attack.Max-attack.Min) + attack.Min

	e.Life.Value -= damage - e.Armour.Value

	return damage
}
