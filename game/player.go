package game

import (
	"log"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Player struct {
	score, multiplier int
	graphics          *PlayerGraphicsComponent
}

func NewPlayer() *Player {
	return &Player{
		graphics:   NewPlayerGraphicsComponent("./assets/fonts/short_xurkit.otf", Colour{255, 255, 255, 255}),
		multiplier: 1,
	}
}

func (p *Player) Update(game *Game, renderer *sdl.Renderer) {
	p.graphics.Update(p, renderer)
}

func (p *Player) Score() int {
	return p.score
}

func (p *Player) OnNotify(entity Entity, event Event) {
	switch event.Type() {
	case TILE_DESTROYED_EVENT:
		tile, ok := entity.(*Tile)
		if !ok {
			log.Fatal("error converting entity to *Tile")
		}
		p.score += p.multiplier * tile.Value()
		p.multiplier++
	case PADDLE_HIT_EVENT:
		p.multiplier = 1
	}
}

type PlayerGraphicsComponent struct {
	font       *ttf.Font
	fontColour Colour
}

func NewPlayerGraphicsComponent(fontPath string, fontColour Colour) *PlayerGraphicsComponent {
	font, err := ttf.OpenFont(fontPath, 32)
	if err != nil {
		log.Fatalf("error opening font: %q", err)
	}

	return &PlayerGraphicsComponent{
		font:       font,
		fontColour: fontColour,
	}
}

func (p *PlayerGraphicsComponent) Update(player *Player, renderer *sdl.Renderer) {
	scoreText, err := p.font.RenderUTF8Solid(
		strconv.Itoa(player.Score()),
		sdl.Color{
			R: p.fontColour.R,
			G: p.fontColour.G,
			B: p.fontColour.B,
			A: p.fontColour.A,
		},
	)
	if err != nil {
		log.Fatalf("error rendering score text: %q", err)
	}

	scoreTextTexture, err := renderer.CreateTextureFromSurface(scoreText)
	if err != nil {
		log.Fatalf("error rendering score text: %q", err)
	}

	if err := renderer.Copy(scoreTextTexture, nil, &sdl.Rect{220, 10, 30, 45}); err != nil {
		log.Fatalf("error rendering score text: %q", err)
	}
}
