package room

import (
	"github.com/pedrokunz/go-design-patterns/internal/domain/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/item"
)

type Room struct {
	Kind    Kind           `json:"kind"`
	Items   []item.Item    `json:"items,omitempty"`
	Enemies []*enemy.Enemy `json:"enemies,omitempty"`
}
