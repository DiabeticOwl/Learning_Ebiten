package main

import (
	"bytes"
	"image"
	_ "image/png"

	resources "first-animation/resources/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	tick  float64
	speed float64
}

var (
	coins *ebiten.Image
)

const (
	// Size of each frame in the image.
	imgSize = 16
	// Amount of frames in the image.
	numFrames = 8
	// Scale of the image.
	imgScale = 6
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.Items_Coin))
	if err != nil {
		panic(err)
	}

	coins = ebiten.NewImageFromImage(img)
}

// Update proceeds the game state.
// Update is called every tick/frame (60 frames each sec by default).
func (g *Game) Update() error {
	g.tick++
	return nil
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	sW, sH := screen.Size()
	cW, cH := imgSize*imgScale, imgSize*imgScale

	// With sW/2 being the center of the X axis and sH/2 being the center of the
	// Y axis in the screen, tX and tY will represent a coordinate that allows
	// half of the width and height of the image to be drawn.
	tX := sW/2 - cW/2
	tY := sH/2 - cH/2

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(imgScale, imgScale)
	ops.GeoM.Translate(float64(tX), float64(tY))

	// Which frame to show.
	frameIndex := int(g.tick/g.speed) % numFrames

	frameX := frameIndex * imgSize
	subImg := coins.SubImage(image.Rect(frameX, 0, frameX+imgSize, imgSize)).(*ebiten.Image)

	screen.DrawImage(subImg, ops)
}

func main() {
	ebiten.SetWindowSize(200, 200)
	ebiten.SetWindowTitle("Draw tiles")

	g := &Game{
		// 60 represents the number of ticks that the Update method counts each second.
		// 10 represents the amount of frames to be shown each second.
		speed: 60 / 10,
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
