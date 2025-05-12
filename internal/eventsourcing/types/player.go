package types

const (
	PlayerAggregateType AggregateType = "player"

	PlayerCreatedEventType EventType = "player.created"
)

func init() {
	RegisterAggregateType(PlayerAggregateType)
	RegisterEventType(PlayerCreatedEventType)
}
