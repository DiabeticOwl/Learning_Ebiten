// draw-image runs a standard application with ebiten and draws two images
// from the same source.
package main

import (
	"bytes"
	"image"
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements the ebiten.Game interface.
type Game struct{}

var owl *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(owlImg))
	if err != nil {
		panic(err)
	}
	owl = ebiten.NewImageFromImage(img)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawOwlImages(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) drawOwlImages(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Rotate(-.5)
	opts.GeoM.Scale(0.1, 0.1)
	opts.GeoM.Translate(45, 200)

	screen.DrawImage(owl, opts)

	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Rotate(.5)
	opts.GeoM.Scale(0.1, 0.1)
	opts.GeoM.Translate(410, 105)

	screen.DrawImage(owl, opts)
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World!")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
