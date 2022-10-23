package main

import (
	_ "image/png"
	"math/rand"
	"time"

	"shoot-the-duck/object"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements the ebiten.Game interface.
type Game struct {
	tick    *int
	speed   float64
	objects []object.Object

	score *int
}

const (
	screenWidth  = 700
	screenHeight = 500
	speed        = 60 / 30
)

// Update proceeds the game state by executing the "Update" method in each
// passed object.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	*g.tick++

	for _, o := range g.objects {
		o.Update()
	}

	return nil
}

// Draw executes the "Draw" method in each passed object.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	for _, o := range g.objects {
		o.Draw(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the
// (logical) screen size. It runs every time the outside size changes.
// If you don't have to adjust the screen size with the outside size, just
// return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// If the screen size value in the Object value is not changed then
	// when the application is rescaled towards the size of the web browser each
	// object that depends on GameScreenWidth and GameScreenHeight will have
	// outdated values.
	object.GameScreenWidth = outsideWidth
	object.GameScreenHeight = outsideHeight

	return outsideWidth, outsideHeight

}

func main() {
	rand.Seed(time.Now().UnixMicro())

	game := &Game{
		tick:  new(int),
		score: new(int),
		speed: speed,
		objects: []object.Object{
			object.NewBackground(),
			object.NewWater(),
			object.NewBackgroundWood(),
			object.NewDuck(),
			object.NewCrosshair(),
			object.NewSideCurtain(),
			object.NewTopCurtain(),
			object.NewScore(),
		},
	}

	object.GameScore = game.score
	object.GameScreenWidth = screenWidth
	object.GameScreenHeight = screenHeight
	object.GameTick = game.tick

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Shoot the Duck")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
