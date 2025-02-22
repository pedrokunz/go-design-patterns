package enemy

import "github.com/pedrokunz/go-design-patterns/domain/core/internal"

type Enemy struct {
	Type   Type
	Armour internal.Armour
	Life   internal.Life
	Attack internal.Attack
}
