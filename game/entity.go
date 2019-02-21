package game

import (
	"log"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Entity interface {
	Subject
	Update(*Game, *sdl.Renderer)
	X() int32
	Y() int32
	Width() int32
	Height() int32
	SetX(int32)
	SetY(int32)
	HandleCollision(Entity)
}

// base data structure for common entity methods
type BaseEntity struct {
	x, y, width, height int32
	observers           []Observer
}

func NewBaseEntity(x, y, width, height int32) *BaseEntity {
	return &BaseEntity{
		x:         x,
		y:         y,
		width:     width,
		height:    height,
		observers: []Observer{},
	}
}

// not implemented as specific to entity
func (b *BaseEntity) Update(*Game, *sdl.Renderer) { log.Fatal("Update: not implemented") }

// not implemented as specific to entity
func (b *BaseEntity) HandleCollision(Entity) { log.Fatal("HandleCollision: not implemented") }

func (b *BaseEntity) AddObserver(observer Observer) {
	b.observers = append(b.observers, observer)
}

func (b *BaseEntity) RemoveObserver(observer Observer) {
	for i, existingObserver := range b.observers {
		if reflect.DeepEqual(observer, existingObserver) {
			b.observers = append(b.observers[:i], b.observers[i+1:]...)
		}
	}
}

func (b *BaseEntity) Notify(entity Entity, event Event) {
	for _, observer := range b.observers {
		observer.OnNotify(entity, event)
	}
}

func (b *BaseEntity) X() int32 {
	return b.x
}

func (b *BaseEntity) Y() int32 {
	return b.y
}

func (b *BaseEntity) SetX(x int32) {
	b.x = x
}

func (b *BaseEntity) SetY(y int32) {
	b.y = y
}

func (b *BaseEntity) Width() int32 {
	return b.width
}

func (b *BaseEntity) Height() int32 {
	return b.height
}
