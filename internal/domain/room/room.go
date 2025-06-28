package room

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/domain/enemy"
	"github.com/pedrokunz/go-design-patterns/internal/domain/item"
)

type Room struct {
	ID      uuid.UUID      `json:"id"`
	Kind    Kind           `json:"kind"`
	Items   []item.Item    `json:"items,omitempty"`
	Enemies []*enemy.Enemy `json:"enemies,omitempty"`
}

type Config struct {
	Kind    Kind
	Items   []item.Item
	Enemies []*enemy.Enemy
}
