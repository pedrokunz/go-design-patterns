package player

import (
	"encoding/json"
	"github.com/pedrokunz/go-design-patterns/internal/domain/internal"
	"math/rand"
)

type Player struct {
	Aggregate *internal.Aggregate `json:"aggregate"`
	Name      string              `json:"name"`
	Armour    internal.Armour     `json:"armour"`
	Life      internal.Life       `json:"life"`
	Attack    internal.Attack     `json:"attack"`
}

func (p *Player) TakeDamage(attack internal.Attack) int {
	damage := rand.Intn(attack.Max-attack.Min) + attack.Min

	p.Life.Value -= damage - p.Armour.Value

	return damage
}

func (p *Player) Apply(event internal.Event) error {
	switch event.Type() {
	case Created:
		return json.Unmarshal(event.Payload(), p)
	}
	return nil
}

func LoadFromHistory(events []internal.Event) (*Player, error) {
	player := &Player{}
	for _, event := range events {
		if err := player.Apply(event); err != nil {
			return nil, err
		}
	}
	return player, nil
}
