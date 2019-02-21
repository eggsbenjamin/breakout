package game

import (
	"io"
	"log"

	"github.com/lmittmann/ppm"
	"github.com/veandco/go-sdl2/sdl"
)

type Tile struct {
	*BaseEntity
	graphics      TileGraphicsComponent
	width, height int32
	value         int
	colour        Colour
}

func NewTile(x, y, width, height int32, value int, colour Colour) *Tile {
	return &Tile{
		BaseEntity: NewBaseEntity(x, y, width, height),
		graphics:   *NewTileGraphicsComponent(),
		width:      width,
		height:     height,
		value:      value,
		colour:     colour,
	}
}

func (t *Tile) Update(game *Game, renderer *sdl.Renderer) {
	t.graphics.Update(t, renderer)
}

func (t *Tile) HandleCollision(Entity) {
	t.Notify(t, NewTileDestroyedEvent())
}

func (t *Tile) Value() int {
	return t.value
}

type TileGraphicsComponent struct{}

func NewTileGraphicsComponent() *TileGraphicsComponent {
	return &TileGraphicsComponent{}
}

func (t *TileGraphicsComponent) Update(tile *Tile, renderer *sdl.Renderer) {
	if err := renderer.SetDrawColor(
		tile.colour.R,
		tile.colour.G,
		tile.colour.B,
		tile.colour.A,
	); err != nil {
		log.Fatalf("error rendering tile: %q", err)
	}

	if err := renderer.FillRect(
		&sdl.Rect{
			X: tile.x,
			Y: tile.y,
			W: tile.width,
			H: tile.height,
		},
	); err != nil {
		log.Fatalf("error rendering tile: %q", err)
	}
}

type TileMap struct {
	tileWidth, tileHeight int32
	tiles                 []Entity
}

type TileMapConfig struct {
	TileWidth, TileHeight int32
}

func (t *TileMap) Tiles() []Entity {
	return t.tiles
}

func NewTileMapFromPPM(cfg TileMapConfig, r io.Reader) (TileMap, error) {
	tileMap := TileMap{
		tileWidth:  cfg.TileWidth,
		tileHeight: cfg.TileHeight,
	}

	img, err := ppm.Decode(r)
	if err != nil {
		return tileMap, err
	}

	for i := 0; i < img.Bounds().Max.X; i++ {
		for j := 0; j < img.Bounds().Max.Y; j++ {
			r, g, b, a := img.At(i, j).RGBA()
			colour := NewColour(uint8(r), uint8(g), uint8(b), uint8(a))

			if isEmptySpace(colour) {
				continue
			}

			tile := NewTile(int32(i)*tileMap.tileWidth, int32(j)*tileMap.tileHeight, tileMap.tileWidth, tileMap.tileHeight, 10, colour)
			tileMap.tiles = append(tileMap.tiles, tile)
		}
	}

	return tileMap, nil
}

func isEmptySpace(colour Colour) bool {
	return colour.R == 0 && colour.G == 0 && colour.B == 0
}
