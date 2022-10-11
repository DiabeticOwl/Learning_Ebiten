package main

import (
	"bytes"
	"image"
	_ "image/png"
	"math"

	resources "basic-shooter/resources/PNG/stall"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements the ebiten.Game interface.
type Game struct{}

var (
	topCurtain  *ebiten.Image
	sideCurtain *ebiten.Image
	bgGreen     *ebiten.Image
	bgWood      *ebiten.Image
)

func decodeImage(imgSlice []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imgSlice))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func init() {
	topCurtain = decodeImage(resources.TopCurtain_png)
	sideCurtain = decodeImage(resources.SideCurtain_png)
	bgGreen = decodeImage(resources.BgGreen_png)
	bgWood = decodeImage(resources.BgWood_png)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawTopCurtain(screen)
	g.drawSideCurtains(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) drawTopCurtain(screen *ebiten.Image) {
	sW, _ := screen.Size()
	topCurW, _ := topCurtain.Size()

	nDraws := int(math.Ceil(float64(sW) / float64(topCurW)))

	for i := 0; i < nDraws; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(topCurW*i), 0)

		screen.DrawImage(topCurtain, opts)
	}
}

func (g *Game) drawSideCurtains(screen *ebiten.Image) {
	// TODO: Make the drawings be behind of the top curtain.
	sW, _ := screen.Size()
	// sidCurW, _ := sideCurtain.Size()
	_, topCurH := topCurtain.Size()

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, float64(topCurH-20))

	screen.DrawImage(sideCurtain, opts)

	opts = &ebiten.DrawImageOptions{}
	// Will draw from right to left.
	opts.GeoM.Scale(-1, 1)
	opts.GeoM.Translate(float64(sW), float64(topCurH-20))

	screen.DrawImage(sideCurtain, opts)
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Simple Shooter Game")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
