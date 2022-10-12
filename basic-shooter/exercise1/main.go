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
	g.drawBackground(screen)
	g.drawBackgroundWood(screen)
	g.drawSideCurtains(screen)
	g.drawTopCurtain(screen)
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

func (g *Game) drawBackgroundWood(screen *ebiten.Image) {
	sW, sH := screen.Size()
	topBgWW, _ := bgWood.Size()

	// The Y position where it draws the wood background.
	sHPos := float64(sH) * .80

	nDraws := int(math.Ceil(float64(sW) / float64(topBgWW)))

	for i := 0; i < nDraws; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(topBgWW*i), sHPos)

		screen.DrawImage(bgWood, opts)
	}
}

func (g *Game) drawSideCurtains(screen *ebiten.Image) {
	sW, sH := screen.Size()
	_, sidCurH := sideCurtain.Size()

	// Setting the maximum height of the curtain by 85% of the screen.
	// The new ratio will determine how much the curtain must belittle or
	// enlarge itself. If the screen is bigger than the curtain image, the scale
	// on the Y axis will be positive (enlarging the image).
	sHLimit := float64(sH) * .85
	sidCurHDelta := float64(sidCurH) - sHLimit
	newSidCurHRatio := sidCurHDelta / float64(sidCurH)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1, 1-newSidCurHRatio)

	screen.DrawImage(sideCurtain, opts)

	opts = &ebiten.DrawImageOptions{}
	// -1 in the X axis will draw from right to left.
	opts.GeoM.Scale(-1, 1-newSidCurHRatio)
	opts.GeoM.Translate(float64(sW), 0)

	screen.DrawImage(sideCurtain, opts)
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	sW, sH := screen.Size()
	bgW, bgH := bgGreen.Size()

	bgWDelta := float64(bgW - sW)
	newbgWRatio := bgWDelta / float64(bgW)
	bgHDelta := float64(bgH - sH)
	newbgHRatio := bgHDelta / float64(bgH)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1-newbgWRatio, 1-newbgHRatio)

	screen.DrawImage(bgGreen, opts)
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(700, 500)
	ebiten.SetWindowTitle("Simple Shooter Game")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
