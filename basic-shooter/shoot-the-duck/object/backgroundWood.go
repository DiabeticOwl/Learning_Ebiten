package object

import (
	"math"
	"shoot-the-duck/utils"

	pngResources "shoot-the-duck/resources/PNG/stall"

	"github.com/hajimehoshi/ebiten/v2"
)

type backgroundWood struct{}

var (
	bgWImg  = utils.DecodeImage(pngResources.BgWood_png)
	bgWW, _ = bgWImg.Size()
)

func NewBackgroundWood() Object {
	return &backgroundWood{}
}

func (bg *backgroundWood) Update() {}

func (bg *backgroundWood) Draw(screen *ebiten.Image) {
	sW, sH := screen.Size()

	// The Y position where it draws the wood background.
	sHPos := float64(sH) * .80

	nDraws := int(math.Ceil(float64(sW) / float64(bgWW)))

	for i := 0; i < nDraws; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(bgWW*i), sHPos)

		screen.DrawImage(bgWImg, opts)
	}
}
