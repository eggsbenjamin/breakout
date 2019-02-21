package game

import "github.com/veandco/go-sdl2/sdl"

type Game struct {
	Width, Height, FPS int32
	Entities           []Entity
	Player             *Player
}

func New(width, height, fps int32, entities []Entity, player *Player) *Game {
	return &Game{
		Width:    width,
		Height:   height,
		FPS:      fps,
		Entities: entities,
		Player:   player,
	}
}

func (g *Game) Update(renderer *sdl.Renderer) {
	for _, entity := range g.Entities {
		entity.Update(g, renderer)
	}

	g.Player.Update(g, renderer)
}

func (g *Game) OnNotify(entity Entity, event Event) {
	switch event.Type() {
	case TILE_DESTROYED_EVENT:
		for i, existingEntity := range g.Entities {
			if entity == existingEntity {
				g.Entities = append(g.Entities[:i], g.Entities[i+1:]...)
			}
		}
	}
}
