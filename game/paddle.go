package game

import (
	"log"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Paddle struct {
	*BaseEntity
	xDelta   int32
	graphics *PaddleGraphicsComponent
	input    *PaddleInputComponent
	colour   Colour
}

func NewPaddle(x, y, width, height int32, colour Colour) *Paddle {
	return &Paddle{
		BaseEntity: NewBaseEntity(x, y, width, height),
		graphics:   NewPaddleGraphicsComponent(),
		input:      NewPaddleInputComponent(),
		colour:     colour,
	}
}

func (p *Paddle) Update(game *Game, renderer *sdl.Renderer) {
	p.graphics.Update(p, renderer)
	p.input.Update(p, game)
}

func (p *Paddle) HandleCollision(entity Entity) {
	p.Notify(p, NewPaddleHitEvent())
}

type PaddleGraphicsComponent struct{}

func NewPaddleGraphicsComponent() *PaddleGraphicsComponent {
	return &PaddleGraphicsComponent{}
}

func (b *PaddleGraphicsComponent) Update(paddle *Paddle, renderer *sdl.Renderer) {
	if err := renderer.FillRect(
		&sdl.Rect{
			X: paddle.x,
			Y: paddle.y,
			W: paddle.width,
			H: paddle.height,
		},
	); err != nil {
		log.Fatalf("error rendering paddle: %q", err)
	}
}

type PaddleInputComponent struct{}

func NewPaddleInputComponent() *PaddleInputComponent {
	return &PaddleInputComponent{}
}

func (p *PaddleInputComponent) Update(paddle *Paddle, game *Game) {
	for {
		e := sdl.PollEvent()
		if e == nil {
			return
		}

		switch e.(type) {
		case *sdl.MouseMotionEvent:
			delta := paddle.X() + e.(*sdl.MouseMotionEvent).XRel
			if delta < 0 {
				paddle.SetX(0)
			} else if delta+paddle.Width() > game.Width {
				paddle.SetX(game.Width - paddle.Width())
			} else {
				paddle.SetX(delta)
			}
		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}
