package enemy

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
	"math/rand"
)

type Enemy struct {
	Race   Race            `json:"race"`
	Armour internal.Armour `json:"armour"`
	Life   internal.Life   `json:"life"`
	Attack internal.Attack `json:"attack"`
}

func New(t Race) *Enemy {
	return &Enemy{
		Race:   t,
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
