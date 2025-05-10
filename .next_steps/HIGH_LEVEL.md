# High-Level Event Sourcing Architecture Diagram

Here's the first diagram that provides a high-level overview of the event sourcing architecture for your Adventure Quest
game:

```mermaid
flowchart TB
    subgraph "Client Interactions"
        UI[UI/Console Interface]
        API[API Endpoints]
    end

    subgraph "Write Path"
        CMD[Commands]
        CMDH[Command Handlers]
        AGG[Aggregates]
    end

    subgraph "Event Storage"
        ES[Event Store]
    end

    subgraph "Read Path"
        PROJ[Projections]
        RM[Read Models/DB]
        QH[Query Handlers]
    end

    subgraph "Side Effects"
        R[Reactors]
        EXT[External Systems]
    end

%% Write flow
    UI -->|1 . Issues commands| CMD
    API -->|1 . Issues commands| CMD
    CMD -->|2 . Validated by| CMDH
    CMDH -->|3 . Updates| AGG
    AGG -->|4 . Produces events| ES
%% Read flow
    ES -->|5 . Feeds| PROJ
    PROJ -->|6 . Updates| RM
    QH -->|7 . Queries| RM
    RM -->|8 . Returns data| UI
    RM -->|8 . Returns data| API
%% Side effects
    ES -->|9 . Triggers| R
    R -->|10 . Affects| EXT
%% Aggregate reconstitution
    ES -.->|Loads history| AGG
```

## Architecture Overview Explanation

This diagram illustrates the core components of an event sourcing architecture:

1. **Client Interactions**: The entry points for user interactions (UI, API)

2. **Write Path**:
    - Commands represent user intentions (e.g., "MovePlayer")
    - Command Handlers validate commands and apply them to the appropriate aggregate
    - Aggregates enforce business rules and generate events when state changes

3. **Event Storage**:
    - The Event Store is the single source of truth, containing all events in sequence
    - Events are never modified once stored (immutable)

4. **Read Path**:
    - Projections subscribe to events and build optimized read models
    - Read Models are updated when relevant events occur
    - Query Handlers access these read models to fulfill client requests

5. **Side Effects**:
    - Reactors listen for events and perform side effects
    - Examples include notifications, achievements, or integration with external systems

6. **Aggregate Reconstitution**:
    - When handling commands, aggregates reload their state from events

This architecture provides several benefits:

- Complete audit history
- Separation of read and write concerns
- Ability to rebuild state from events
- Excellent support for eventual consistency

The next diagrams will dive deeper into each of these components with more specific examples from your Adventure Quest
game.