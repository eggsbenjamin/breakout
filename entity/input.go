package entity

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type PlayerInputComponent struct{}

func NewPlayerInputComponent() *PlayerInputComponent {
	return nil
}

func (p *PlayerInputComponent) Update(entity Entity) {
	for {
		e := sdl.PollEvent()
		if e == nil {
			return
		}

		switch e.(type) {
		case *sdl.MouseMotionEvent:
			//		player.MoveDelta = e.(*sdl.MouseMotionEvent).YRel
		case *sdl.QuitEvent:
			os.Exit(0) // TODO: remove
		}
	}
}
