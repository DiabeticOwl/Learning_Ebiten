package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"image"
	_ "image/png"

	resources "animation-with-different-sizes/resources/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	tick      float64
	speed     float64
	frames    []frameSpec
	numFrames int
}

type framesSpec struct {
	Frames []frameSpec `json:"frames"`
}

type frameSpec struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

var (
	coins *ebiten.Image
	//go:embed resources/coins.json
	jsonFile embed.FS
)

const (
	// Scale of the image.
	imgScale = 4
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.CoinsPNG))
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

// Layout takes the outside size (e.g., the window size) and returns the
// (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just
// return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	sW, sH := screen.Size()

	// By calculating the moment in which the application is each update of
	// the game, the coordinates and details stored on the configuration file
	// for the frame to draw are instantiated in the variable called "frame".
	frameIndex := int(g.tick/g.speed) % g.numFrames
	frame := g.frames[frameIndex]
	rect := image.Rect(frame.X, frame.Y, frame.X+frame.W, frame.Y+frame.H)
	subImg := coins.SubImage(rect).(*ebiten.Image)

	// With sW/2 being the center of the X axis and sH/2 being the center of the
	// Y axis in the screen, tX and tY will represent a coordinate that allows
	// half of the width and height of the image to be drawn.
	tX := sW/2 - frame.W/2*imgScale
	tY := sH/2 - frame.H/2*imgScale

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(imgScale, imgScale)
	ops.GeoM.Translate(float64(tX), float64(tY))

	screen.DrawImage(subImg, ops)
}

// buildFrames takes the path of the configuration file for the frames of the
// sprite sheet and parses it on the game instance attached to this method.
func (g *Game) buildFrames() {
	j, err := jsonFile.ReadFile("resources/coins.json")
	if err != nil {
		panic(err)
	}

	fSpec := &framesSpec{}

	json.Unmarshal(j, fSpec)

	g.frames = fSpec.Frames
	g.numFrames = len(g.frames)
}

func main() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Animation with various sizes")

	g := &Game{
		// 20 frames per second.
		speed: 60 / 20,
	}

	g.buildFrames()

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
