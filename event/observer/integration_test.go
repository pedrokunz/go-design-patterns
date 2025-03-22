package observer_test

import (
	"github.com/pedrokunz/go-design-patterns/event"
	"github.com/pedrokunz/go-design-patterns/event/observer"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestObserver(t *testing.T) {
	notifier := observer.NewNotifier()
	playerObserver, buildErr := observer.NewFactory(observer.PlayerObserver).Build(
		observer.PlayerObserverConfig{
			Name: "player_name",
		},
	)

	require.NoError(t, buildErr, "error building player observer")

	notifier.Attach(playerObserver)

	notifier.Notify(event.New(event.PlayerJoined))

	// Output:
	// Player player_name received event player_joined

	notifier.Notify(event.New(event.PlayerLeft))

	// Output:
	// Player player_name received event player_left
}
