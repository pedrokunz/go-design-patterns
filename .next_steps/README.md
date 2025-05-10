# Event Sourcing Implementation Guide

This document outlines the steps to implement event sourcing in the Adventure Quest game project.

- [High-Level Event Sourcing Architecture](./HIGH_LEVEL.md)
    - Overview of the entire system
    - Shows the relationship between commands, events, aggregates, event store, projections, and reactors
    - Demonstrates the separation of writing and read paths
- [Command Processing Flow](./COMMAND_PROCESSING_FLOW.md)
    - Details how commands are processed
    - Shows validation, handling, and event generation
    - Includes the role of aggregates in processing commands
- [Aggregate Structure and Responsibility](./AGGREGATE_STRUCTURE.md)
    - Illustrates what an aggregate is and how it functions
    - Shows internal state management via events
    - Explains event application and state reconstitution
- [Projection System](./PROJECTION_SYSTEM.md)
    - Demonstrates how projections create read models from events
    - Shows practical database interaction examples
    - Explains how projections optimize queries
- [Reactor System](./REACTOR_SYSTEM.md)
    - Shows how reactors listen to events and perform side effects
    - Provides examples of practical game mechanics (achievements, notifications)
    - Illustrates async processing capabilities
- [Event Store and Persistence](./EVENT_STORE.md)
    - Details event storage mechanisms
    - Shows event retrieval for aggregate reconstitution
    - Explains versioning and conflict detection
- [Scaling and Performance Considerations](./SCALING_AND_PERFORMANCE.md)
    - Shows how different components scale
    - Demonstrates read/write separation benefits
    - Explains event sourcing performance characteristics

## 1. Core Event Sourcing Concepts

### 1.1. Key Components to Implement

- **Events**: Immutable records of state changes (e.g., PlayerMovedEvent, ItemCollectedEvent)
- **Commands**: Represent user intentions (e.g., MovePlayerCommand, CollectItemCommand)
- **Aggregates**: Domain objects that maintain consistency (e.g., Player, Room)
- **Event Store**: Persistent storage for event sequences
- **Projections**: Read models built from events
- **Reactors**: Components that react to events with side effects

## 2. Project Structure Refactoring

```
go-design-patterns/
├── cmd/
│   └── adventure-quest/
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── command/              # Command handlers
│   │   ├── query/                # Query handlers
│   │   └── usecase/              # Use case implementations
│   ├── domain/
│   │   ├── aggregate/            # Aggregate roots
│   │   │   ├── enemy/
│   │   │   ├── game/
│   │   │   ├── internal/
│   │   │   ├── item/
│   │   │   ├── player/
│   │   │   ├── room/
│   │   │   └── game/
│   │   ├── command/              # Command definitions
│   │   ├── event/                # Event definitions
│   │   └── valueobject/          # Value objects
│   ├── eventsourcing/            # Generic event sourcing toolkit
│   ├── infrastructure/
│   │   ├── eventstore/           # Event storage implementation
│   │   ├── projection/           # Projection implementations
│   │   ├── reactor/              # Event reactors
│   │   └── repository/           # Repository implementations
│   └── interfaces/
│       ├── api/                  # API handlers
└──     └── console/              # Console UI
```

## 3. Implementation Steps

### 3.1. Create Event Sourcing Base Components

1. Define base event interface and structure:

  ```go
  // internal/eventsourcing/event.go
package eventsourcing

import (
	"github.com/google/uuid"
	"time"
)

type Event interface {
	AggregateID() uuid.UUID
	AggregateType() string
	EventType() string
	Version() int
	Timestamp() time.Time
	Payload() interface{}
}

type BaseEvent struct {
	ID           uuid.UUID
	AggrID       uuid.UUID
	AggrType     string
	EventType    string
	EventVersion int
	EventTime    time.Time
	EventPayload interface{}
}

// Implement Event interface methods
```

2. Define aggregate interface:

  ```go
  // internal/eventsourcing/aggregate.go
package eventsourcing

type Aggregate interface {
	ID() uuid.UUID
	Type() string
	Version() int
	ApplyEvent(event Event)
	UncommittedEvents() []Event
	ClearUncommittedEvents()
}

type BaseAggregate struct {
	AggregateID      uuid.UUID
	AggregateType    string
	AggregateVersion int
	Changes          []Event
}

```

3. Create command and handler interfaces:

```go
// internal/eventsourcing/command.go
package eventsourcing

type Command interface {
	CommandType() string
	AggregateID() uuid.UUID
	AggregateType() string
}

type CommandHandler interface {
	Handle(cmd Command) error
}
```

4. Implement event store interface:

```go
// internal/eventsourcing/eventstore.go
package eventsourcing

type EventStore interface {
	SaveEvents(aggregateID uuid.UUID, events []Event, expectedVersion int) error
	GetEvents(aggregateID uuid.UUID) ([]Event, error)
}
```

### 3.2. Define Domain Events

```go
// internal/domain/event/player_events.go
package event

type PlayerCreatedEvent struct {
	eventsourcing.BaseEvent
	Name   string
	Health int
}

type PlayerMovedEvent struct {
	eventsourcing.BaseEvent
	FromRoomID uuid.UUID
	ToRoomID   uuid.UUID
}

type ItemCollectedEvent struct {
	eventsourcing.BaseEvent
	ItemID   uuid.UUID
	ItemName string
}
```

### 3.3. Define Commands

```go
// internal/domain/command/player_commands.go
package command

type CreatePlayerCommand struct {
	PlayerID uuid.UUID
	Name     string
}

type MovePlayerCommand struct {
	PlayerID     uuid.UUID
	TargetRoomID uuid.UUID
}

type CollectItemCommand struct {
	PlayerID uuid.UUID
	ItemID   uuid.UUID
}
```

### 3.4. Implement Aggregates

```go
// internal/domain/aggregate/player/player.go
package player

type Player struct {
	eventsourcing.BaseAggregate
	name      string
	health    int
	inventory []uuid.UUID
	roomID    uuid.UUID
}

func (p *Player) ApplyEvent(event eventsourcing.Event) {
	switch e := event.Payload().(type) {
	case *event.PlayerCreatedEvent:
		p.applyPlayerCreated(e)
	case *event.PlayerMovedEvent:
		p.applyPlayerMoved(e)
	case *event.ItemCollectedEvent:
		p.applyItemCollected(e)
	}

	p.AggregateVersion++
}

// Apply specific event handlers
```

### 3.5. Implement Command Handlers

```go
// internal/app/command/player_command_handler.go
package command

type PlayerCommandHandler struct {
	repository player.Repository
	eventStore eventsourcing.EventStore
}

func (h *PlayerCommandHandler) Handle(cmd eventsourcing.Command) error {
	switch c := cmd.(type) {
	case *command.CreatePlayerCommand:
		return h.handleCreatePlayer(c)
	case *command.MovePlayerCommand:
		return h.handleMovePlayer(c)
		// Other handlers
	}
	return fmt.Errorf("unknown command type: %T", cmd)
}
```

### 3.6. Implement Repositories

```go
// internal/infrastructure/repository/player_repository.go
package repository

type PlayerRepository struct {
	eventStore eventsourcing.EventStore
}

func (r *PlayerRepository) Save(player *player.Player) error {
	return r.eventStore.SaveEvents(
		player.ID(),
		player.UncommittedEvents(),
		player.Version(),
	)
}

func (r *PlayerRepository) Get(id uuid.UUID) (*player.Player, error) {
	events, err := r.eventStore.GetEvents(id)
	if err != nil {
		return nil, err
	}

	p := player.NewPlayer()
	for _, event := range events {
		p.ApplyEvent(event)
	}

	return p, nil
}
```

### 3.7. Implement Projections

```go
// internal/infrastructure/projection/player_projection.go
package projection

type PlayerProjection struct {
	db repository.ReadDB
}

func (p *PlayerProjection) Apply(event eventsourcing.Event) error {
	switch e := event.Payload().(type) {
	case *event.PlayerCreatedEvent:
		return p.applyPlayerCreated(event, e)
	case *event.PlayerMovedEvent:
		return p.applyPlayerMoved(event, e)
		// Other event handlers
	}
	return nil
}
```

### 3.8. Implement Reactors

```go
// internal/infrastructure/reactor/achievement_reactor.go
package reactor

type AchievementReactor struct {
	achievementService service.AchievementService
}

func (r *AchievementReactor) Handle(event eventsourcing.Event) error {
	switch e := event.Payload().(type) {
	case *event.ItemCollectedEvent:
		return r.checkItemCollectionAchievements(event, e)
		// Other event reactions
	}
	return nil
}
```

## 4. Use Case Implementation Examples

### 4.1. Player Movement Use Case

```go
// internal/app/usecase/move_player.go
package usecase

type MovePlayerUseCase struct {
	commandHandler command.PlayerCommandHandler
	roomRepository repository.RoomRepository
}

func (uc *MovePlayerUseCase) Execute(playerID, targetRoomID uuid.UUID) error {
	// Validate move is possible
	room, err := uc.roomRepository.Get(targetRoomID)
	if err != nil {
		return err
	}

	// Create and handle command
	cmd := &command.MovePlayerCommand{
		PlayerID:     playerID,
		TargetRoomID: targetRoomID,
	}

	return uc.commandHandler.Handle(cmd)
}
```

## 5. Testing Strategy

1. **Unit Testing**: Test aggregates, commands, events in isolation
2. **Integration Testing**: Test command handlers with event store
3. **End-to-End Testing**: Test use cases with all components

## 6. Next Steps

1. Implement the basic event sourcing infrastructure
2. Define the core domain events and commands
3. Implement the player aggregate and repository
4. Create initial projections for the player state
5. Implement basic use cases (create player, move, collect items)
6. Add reactors for game mechanics
7. Develop the user interface

By following this step-by-step guide, you'll have a robust event sourcing implementation for your adventure game that
follows clean architecture principles and leverages the benefits of event sourcing.

```