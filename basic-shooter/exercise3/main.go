package main

import (
	"bytes"
	"image"
	_ "image/png"
	"math"
	"math/rand"
	"time"

	hudResources "basic-shooter-exercise-3/resources/PNG/hud"
	objResources "basic-shooter-exercise-3/resources/PNG/objects"
	pngResources "basic-shooter-exercise-3/resources/PNG/stall"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements the ebiten.Game interface.
type Game struct {
	tick  int
	speed float64

	duckInitialHeightPos float64
	ducks                map[*Duck]struct{}

	waterOffsetX float64
	waterOffsetY float64

	screenWidth int

	crossHair *CrossHair
}

var (
	waterDirectionX int
	waterDirectionY int

	topCurtain             *ebiten.Image
	sideCurtain            *ebiten.Image
	bgGreen                *ebiten.Image
	bgWood                 *ebiten.Image
	water1                 *ebiten.Image
	duckOutlineTargetWhite *ebiten.Image
	crosshairImage         *ebiten.Image
)

const (
	waterOffsetXSpeed = 1
	maxWaterOffsetX   = 100
	waterOffsetYSpeed = .3
	maxWaterOffsetY   = 20

	maxDuckOffsetY = 10

	debouncer = 200 * time.Millisecond
)

func decodeImage(imgSlice []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imgSlice))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func init() {
	topCurtain = decodeImage(pngResources.TopCurtain_png)
	sideCurtain = decodeImage(pngResources.SideCurtain_png)
	bgGreen = decodeImage(pngResources.BgGreen_png)
	water1 = decodeImage(pngResources.Water1_png)
	bgWood = decodeImage(pngResources.BgWood_png)
	duckOutlineTargetWhite = decodeImage(objResources.DuckOutlineTargetWhite_png)
	crosshairImage = decodeImage(hudResources.Crosshair_png)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.tick++

	// Water Logic
	if g.waterOffsetX >= maxWaterOffsetX {
		waterDirectionX = -waterDirectionX
	} else if g.waterOffsetX <= 0 {
		waterDirectionX = 1
	}
	g.waterOffsetX += float64(waterDirectionX) * waterOffsetXSpeed

	if g.waterOffsetY >= maxWaterOffsetY {
		waterDirectionY = -waterDirectionY
	} else if g.waterOffsetY <= 0 {
		waterDirectionY = 1
	}
	g.waterOffsetY += float64(waterDirectionY) * waterOffsetYSpeed

	// CrossHair Logic
	g.crossHair.x, g.crossHair.y = ebiten.CursorPosition()

	now := time.Now()
	leftButtonPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if leftButtonPressed && now.Sub(g.crossHair.lastClickAt) > debouncer {
		g.crossHair.lastClickAt = now
		g.crossHair.clicked = leftButtonPressed

		for duck := range g.ducks {
			dXPos := duck.offsetX + float64(duck.w)
			// The greater the position in Y is the more it will be towards
			// the bottom of the screen.
			dYPos := g.duckInitialHeightPos - float64(duck.h) - duck.offsetY

			xDelta := math.Abs(float64(g.crossHair.x) - dXPos)
			yDelta := math.Abs(float64(g.crossHair.y) - dYPos)

			// 40, 10 and 100 is an arbitrary number for how close the crosshair
			// needs to be in order to take down the duck.
			if xDelta <= 40 && (yDelta >= 10 && yDelta <= 100) {
				delete(g.ducks, duck)
				g.crossHair.clicked = false
				break
			}
		}
	}

	// Ducks logic
	if g.tick%60 == 0 && rand.Float64() < 0.6 {
		g.ducks[newDuck(len(g.ducks)+1)] = struct{}{}
	}

	for duck := range g.ducks {
		duck.offsetX += 1.5
		duck.offsetY += 1.5 * float64(duck.yDirection)

		if int(duck.offsetX) >= g.screenWidth {
			delete(g.ducks, duck)
			g.crossHair.clicked = false
			continue
		}

		if rand.Float64() < 0.4 {
			duck.yDirection *= -1
		}

		if duck.offsetY >= maxDuckOffsetY {
			duck.yDirection = -1
		}
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBackground(screen)
	g.drawMovingWater(screen)
	g.drawBackgroundWood(screen)
	g.drawMovingDucks(screen)
	g.drawCrossHair(screen)
	g.drawSideCurtains(screen)
	g.drawTopCurtain(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the
// (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just
// return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) drawCrossHair(screen *ebiten.Image) {
	croHaiW, croHaiH := g.crossHair.img.Size()

	// Coordinates must be adjusted to halve of the image size in order to focus
	// the image towards the mouse.
	// Must be subtracted because of the direction in which the drawing cursor's
	// direction.
	croHaiXPos := float64(g.crossHair.x) - float64(croHaiW)/2
	croHaiYPos := float64(g.crossHair.y) - float64(croHaiH)/2

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(croHaiXPos, croHaiYPos)

	screen.DrawImage(g.crossHair.img, opts)
}

func (g *Game) drawMovingDucks(screen *ebiten.Image) {
	sidCurW, _ := sideCurtain.Size()

	for duck := range g.ducks {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(sidCurW)-float64(duck.w)/2, g.duckInitialHeightPos-float64(duck.h))
		opts.GeoM.Translate(duck.offsetX, duck.offsetY)

		screen.DrawImage(duck.img, opts)
	}
}

func (g *Game) drawMovingWater(screen *ebiten.Image) {
	sW, sH := screen.Size()
	wat1W, wat1H := water1.Size()

	waterScale := .4

	sHPos := float64(sH) * .82
	sHPos -= float64(wat1H) * waterScale

	// The greater this number is the more height is taking from the screen i.e.,
	// the drawing will go more towards the bottom.
	g.duckInitialHeightPos = sHPos + sHPos*.10

	nDraws := int(math.Ceil(float64(sW) / (float64(wat1W) * waterScale)))

	for i := -1; i < nDraws; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(waterScale, waterScale)

		imgW := float64(wat1W*i) * waterScale

		opts.GeoM.Translate(imgW, sHPos)
		opts.GeoM.Translate(g.waterOffsetX, g.waterOffsetY)

		screen.DrawImage(water1, opts)
	}
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
	rand.Seed(time.Now().UnixMicro())

	game := &Game{
		speed:       60 / 30,
		screenWidth: 700,
		crossHair:   newCrosshair(),
		ducks:       make(map[*Duck]struct{}),
	}

	ebiten.SetWindowSize(700, 500)
	ebiten.SetWindowTitle("Simple Shooter Game")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
