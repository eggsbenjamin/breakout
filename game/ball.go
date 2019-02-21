package game

import (
	"log"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Ball struct {
	*BaseEntity
	physics                      *BallPhysicsComponent
	graphics                     *BallGraphicsComponent
	radius, xVelocity, yVelocity int32
	colour                       Colour
}

func NewBall(x, y, radius, xVelocity, yVelocity int32, colour Colour) *Ball {
	return &Ball{
		BaseEntity: NewBaseEntity(x, y, radius*2, radius*2),
		physics:    NewBallPhysicsComponent(),
		graphics:   NewBallGraphicsComponent(),
		radius:     radius,
		xVelocity:  xVelocity,
		yVelocity:  yVelocity,
		colour:     colour,
	}
}

func (b *Ball) Width() int32 {
	return b.radius * 2
}

func (b *Ball) Height() int32 {
	return b.radius * 2
}

func (b *Ball) NextPos() (int32, int32) {
	return b.x + b.xVelocity, b.y + b.yVelocity
}

func (b *Ball) Update(game *Game, renderer *sdl.Renderer) {
	b.physics.Update(b, game)
	b.graphics.Update(b, renderer)
}

func (b *Ball) HandleCollision(entity Entity) {
	if b.collidedFromLeft(entity) || b.collidedFromRight(entity) {
		b.xVelocity *= -1
	}

	if b.collidedFromTop(entity) || b.collidedFromBottom(entity) {
		b.yVelocity *= -1
	}
}

func (b *Ball) collidedFromLeft(entity Entity) bool {
	nextX, _ := b.NextPos()

	return b.x+b.width < entity.X() && // was not colliding
		nextX+b.width >= entity.X()
}

func (b *Ball) collidedFromRight(entity Entity) bool {
	nextX, _ := b.NextPos()

	return b.x >= entity.X()+entity.Width() && // was not colliding
		nextX <= entity.X()+entity.Width()
}

func (b *Ball) collidedFromTop(entity Entity) bool {
	_, nextY := b.NextPos()

	return b.y+b.height < entity.Y() && // was not colliding
		nextY+b.height >= entity.Y()
}

func (b *Ball) collidedFromBottom(entity Entity) bool {
	_, nextY := b.NextPos()

	return b.y >= entity.Y()+entity.Height() && // was not colliding
		nextY <= entity.Y()+entity.Height()
}

type BallGraphicsComponent struct{}

func NewBallGraphicsComponent() *BallGraphicsComponent {
	return &BallGraphicsComponent{}
}

func (b *BallGraphicsComponent) Update(ball *Ball, renderer *sdl.Renderer) {
	if ok := gfx.FilledCircleRGBA(
		renderer,
		ball.x+ball.radius,
		ball.y+ball.radius,
		ball.radius,
		ball.colour.R,
		ball.colour.G,
		ball.colour.B,
		ball.colour.A,
	); !ok {
		log.Fatal("error rendering ball")
	}
}

type BallPhysicsComponent struct{}

func NewBallPhysicsComponent() *BallPhysicsComponent {
	return &BallPhysicsComponent{}
}

func (b *BallPhysicsComponent) Update(ball *Ball, game *Game) {
	nextX, nextY := ball.NextPos()

	if nextX+ball.radius >= game.Width || nextX-ball.radius <= 0 {
		ball.xVelocity = -ball.xVelocity
	}

	if nextY+ball.radius >= game.Height || nextY-ball.radius <= 0 {
		ball.yVelocity = -ball.yVelocity
	}

	for _, entity := range game.Entities {
		if entity == ball {
			continue
		}

		if Intersects(
			Rectangle{X: nextX, Y: nextY, Width: ball.Width(), Height: ball.Height()},
			Rectangle{X: entity.X(), Y: entity.Y(), Width: entity.Width(), Height: entity.Height()},
		) {
			ball.HandleCollision(entity)
			entity.HandleCollision(ball)
			break
		}
	}

	ball.SetX(ball.x + ball.xVelocity)
	ball.SetY(ball.y + ball.yVelocity)
}
