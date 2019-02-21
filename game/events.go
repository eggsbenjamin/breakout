package game

type Event interface {
	Type() int
}

type BaseEvent struct {
	eventType int
}

func (e BaseEvent) Type() int {
	return e.eventType
}

func NewBaseEvent(eventType int) BaseEvent {
	return BaseEvent{
		eventType: eventType,
	}
}

type Observer interface {
	OnNotify(Entity, Event)
}

type Subject interface {
	AddObserver(Observer)
	RemoveObserver(Observer)
	Notify(Entity, Event)
}

const (
	TILE_DESTROYED_EVENT = 1
	PADDLE_HIT_EVENT     = 2
)

type PaddleHitEvent struct {
	BaseEvent
}

func NewPaddleHitEvent() PaddleHitEvent {
	return PaddleHitEvent{
		BaseEvent: NewBaseEvent(PADDLE_HIT_EVENT),
	}
}

type TileDestroyedEvent struct {
	BaseEvent
}

func NewTileDestroyedEvent() TileDestroyedEvent {
	return TileDestroyedEvent{
		BaseEvent: NewBaseEvent(TILE_DESTROYED_EVENT),
	}
}
