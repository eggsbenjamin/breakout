package entity

import (
	"github.com/eggsbenjamin/breakout/game"
	"github.com/veandco/go-sdl2/sdl"
)

type Entity interface {
	GetPosition() game.Point
	GetY() int32
	SetX(int32)
	SetY(int32)
	GetWidth() int32
	GetHeight() int32
	Update(*game.Game, *sdl.Renderer)
}
