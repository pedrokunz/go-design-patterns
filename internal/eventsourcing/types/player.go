package types

const (
	PlayerAggregate AggregateType = "player"

	PlayerCreated EventType = "player.created"
)

func init() {
	RegisterAggregateType(PlayerAggregate)
	RegisterEventType(PlayerCreated)
}
