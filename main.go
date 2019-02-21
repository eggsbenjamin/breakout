package main

import (
	"log"
	"os"
	"strconv"

	"github.com/eggsbenjamin/breakout/game"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func init() {
	if err := ttf.Init(); err != nil {
		log.Fatalf("error initialising sdl ttf pkg: %q", err)
	}
}

func main() {
	var (
		width  = MustGetInt32Env("WIDTH")
		height = MustGetInt32Env("HEIGHT")
		fps    = MustGetInt32Env("FPS")
	)

	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("breakout", 0, 0, width, height, sdl.WINDOW_FULLSCREEN_DESKTOP)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Destroy()

	sdl.SetHint(sdl.HINT_RENDER_VSYNC, "1")
	sdl.SetHint(sdl.HINT_VIDEO_DOUBLE_BUFFER, "1")

	displayMode, err := sdl.GetDesktopDisplayMode(0)
	if err != nil {
		log.Fatal(err)
	}

	if err = renderer.SetScale(float32(displayMode.W)/float32(width), float32(displayMode.H)/float32(height)); err != nil {
		log.Fatal(err)
	}

	sdl.ShowCursor(0)

	renderTexture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, width, height)
	if err != nil {
		log.Fatal(err)
	}

	if err := renderer.SetRenderTarget(renderTexture); err != nil {
		log.Fatal(err)
	}

	levelImg, err := os.Open("./assets/levels/16x14_level.ppm")
	if err != nil {
		log.Fatal(err)
	}

	tileMapCfg := game.TileMapConfig{
		TileWidth:  width / 16,
		TileHeight: height / 14,
	}

	tileMap, err := game.NewTileMapFromPPM(tileMapCfg, levelImg)
	if err != nil {
		log.Fatal(err)
	}

	entities := []game.Entity{}

	ball := game.NewBall(200, 200, height/48, 1, 1, game.Colour{255, 255, 255, 255})
	paddle := game.NewPaddle(width/2, height-2*(height/14), width/16, height/14, game.Colour{255, 255, 255, 255})
	tiles := tileMap.Tiles()

	entities = append(entities, ball)
	entities = append(entities, paddle)
	entities = append(entities, tiles...)

	player := game.NewPlayer()

	game := game.New(
		width,
		height,
		fps,
		entities,
		player,
	)

	for _, tile := range tiles {
		tile.AddObserver(game)
		tile.AddObserver(player)
	}

	run(game, renderTexture, renderer, window)
}

func run(game *game.Game, renderTexture *sdl.Texture, renderer *sdl.Renderer, window *sdl.Window) {
	for {
		renderer.SetDrawColor(0, 0, 0, 0)

		if err := renderer.Clear(); err != nil {
			log.Fatal(err)
		}

		game.Update(renderer)

		if err := renderer.SetRenderTarget(nil); err != nil {
			log.Fatal(err)
		}

		if err := renderer.Copy(renderTexture, nil, nil); err != nil {
			log.Fatal(err)
		}

		renderer.Present()

		if err := renderer.SetRenderTarget(renderTexture); err != nil {
			log.Fatal(err)
		}

		sdl.Delay(1000 / uint32(game.FPS))
	}
}

func MustGetEnv(k string) string {
	v, ok := os.LookupEnv(k)
	if !ok {
		log.Fatalf("env var %s not found", k)
	}
	return v
}

func MustGetInt32Env(k string) int32 {
	stringVal := MustGetEnv(k)
	v, err := strconv.ParseInt(stringVal, 10, 32)
	if err != nil {
		log.Fatalf("env var %s is not an integer: '%s'", k, stringVal)
	}

	return int32(v)
}
