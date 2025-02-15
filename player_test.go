package godesignpatterns_test

import (
	"testing"

	godesignpatterns "github.com/pedrokunz/go-design-patterns"
)

func TestPlayer(t *testing.T) {
	t.Run("constructs a player", func(t *testing.T) {
		got := godesignpatterns.NewPlayer("Elmster")
		want := godesignpatterns.Player{Name: "Elmster"}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
