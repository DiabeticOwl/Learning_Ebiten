package object

import (
	"math"
	pngResources "shoot-the-duck/resources/PNG/stall"
	"shoot-the-duck/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	waterOffsetXSpeed = 1
	maxWaterOffsetX   = 100
	waterOffsetYSpeed = .3
	maxWaterOffsetY   = 20
)

var (
	DuckInitialHeightPos float64

	waterDirectionX int
	waterDirectionY int
	waterOffsetX    float64
	waterOffsetY    float64

	water1       = utils.DecodeImage(pngResources.Water1_png)
	wat1W, wat1H = water1.Size()
)

type water struct{}

func NewWater() Object {
	return &water{}
}

func (bg *water) Update() {
	// Water Logic
	if waterOffsetX >= maxWaterOffsetX {
		waterDirectionX = -waterDirectionX
	} else if waterOffsetX <= 0 {
		waterDirectionX = right
	}
	waterOffsetX += float64(waterDirectionX) * waterOffsetXSpeed

	if waterOffsetY >= maxWaterOffsetY {
		waterDirectionY = -waterDirectionY
	} else if waterOffsetY <= 0 {
		waterDirectionY = down
	}
	waterOffsetY += float64(waterDirectionY) * waterOffsetYSpeed
}

func (bg *water) Draw(screen *ebiten.Image) {
	sW, sH := screen.Size()

	waterScale := .4

	sHPos := float64(sH) * .82
	sHPos -= float64(wat1H) * waterScale

	// The greater this number is the more height is taking from the screen i.e.,
	// the drawing will go more towards the bottom.
	DuckInitialHeightPos = sHPos + sHPos*.10

	nDraws := int(math.Ceil(float64(sW) / (float64(wat1W) * waterScale)))

	for i := -2; i < nDraws; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(waterScale, waterScale)

		imgW := float64(wat1W*i) * waterScale

		opts.GeoM.Translate(imgW, sHPos)
		opts.GeoM.Translate(waterOffsetX, waterOffsetY)

		screen.DrawImage(water1, opts)
	}
}
