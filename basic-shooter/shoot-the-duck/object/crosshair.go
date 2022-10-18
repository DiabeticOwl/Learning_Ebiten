package object

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"shoot-the-duck/resources/PNG/hud"
	"shoot-the-duck/utils"
)

const (
	// Used for delimiting how long a click can take. If the user presses for
	// more than 200 milliseconds then the game will see it as an unique click.
	debouncer = 200 * time.Millisecond
)

type crossHair struct {
	img         *ebiten.Image
	x           int
	y           int
	lastClickAt time.Time
}

func NewCrosshair() Object {
	return &crossHair{
		img: utils.DecodeImage(hud.Crosshair_png),
	}
}

func (c *crossHair) Update() {
	c.x, c.y = ebiten.CursorPosition()

	now := time.Now()
	leftButtonPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if leftButtonPressed && now.Sub(c.lastClickAt) > debouncer {
		oldScore := *GameScore
		c.lastClickAt = now

		for duck := range Ducks {
			dXPos := duck.offsetX + float64(duck.w)
			// The greater the position in Y is the more it will be towards
			// the bottom of the screen.
			dYPos := DuckInitialHeightPos - float64(duck.h) - duck.offsetY

			xDelta := math.Abs(float64(c.x) - dXPos)
			yDelta := math.Abs(float64(c.y) - dYPos)

			// 40, 10 and 100 is an arbitrary number for how close the crosshair
			// needs to be in order to take down the duck.
			if xDelta <= 40 && (yDelta >= 10 && yDelta <= 100) {
				delete(Ducks, duck)

				*GameScore += 10
				break
			}
		}

		if oldScore == *GameScore {
			*GameScore -= 5
		}
	}
}

func (c *crossHair) Draw(screen *ebiten.Image) {
	croHaiW, croHaiH := c.img.Size()

	// Coordinates must be adjusted to halve of the image size in order to focus
	// the image towards the mouse.
	// Must be subtracted because of the direction in which the drawing cursor's
	// direction.
	croHaiXPos := float64(c.x) - float64(croHaiW)/2
	croHaiYPos := float64(c.y) - float64(croHaiH)/2

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(croHaiXPos, croHaiYPos)

	screen.DrawImage(c.img, opts)
}
