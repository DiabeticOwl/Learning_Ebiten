package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game implements the ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World!")

	g.drawBlueImage(screen)
	g.drawRedImage(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) drawBlueImage(screen *ebiten.Image) {
	img := ebiten.NewImage(100, 100)
	img.Fill(color.RGBA{0, 0, 255, 1})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(50, 100)
	opts.GeoM.Rotate(0.5)

	screen.DrawImage(img, opts)
}

func (g *Game) drawRedImage(screen *ebiten.Image) {
	img := ebiten.NewImage(100, 100)
	img.Fill(color.RGBA{255, 0, 0, 1})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(150, 100)
	opts.GeoM.Rotate(0.5)

	screen.DrawImage(img, opts)
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World!")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
