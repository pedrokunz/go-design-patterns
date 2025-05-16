package types

const (
	PlayerAggregate AggregateType = "player"

	PlayerCreatedEventType EventType = "player.created"
)

func init() {
	RegisterAggregateType(PlayerAggregate)
	RegisterEventType(PlayerCreatedEventType)
}
