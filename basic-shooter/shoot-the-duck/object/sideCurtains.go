package object

import (
	"shoot-the-duck/utils"

	pngResources "shoot-the-duck/resources/PNG/stall"

	"github.com/hajimehoshi/ebiten/v2"
)

type sideCurtain struct{}

var (
	sidImg           = utils.DecodeImage(pngResources.SideCurtain_png)
	sidCurW, sidCurH = sidImg.Size()
)

func NewSideCurtain() Object {
	return &sideCurtain{}
}

func (sidCur *sideCurtain) Update() {}

func (sidCur *sideCurtain) Draw(screen *ebiten.Image) {
	sW, sH := screen.Size()

	// Setting the maximum height of the curtain by 85% of the screen.
	// The new ratio will determine how much the curtain must belittle or
	// enlarge itself. If the screen is bigger than the curtain image, the scale
	// on the Y axis will be positive (enlarging the image).
	sHLimit := float64(sH) * .85
	sidCurHDelta := float64(sidCurH) - sHLimit
	newSidCurHRatio := sidCurHDelta / float64(sidCurH)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1, 1-newSidCurHRatio)

	screen.DrawImage(sidImg, opts)

	opts = &ebiten.DrawImageOptions{}
	// -1 in the X axis will draw from right to left.
	opts.GeoM.Scale(-1, 1-newSidCurHRatio)
	opts.GeoM.Translate(float64(sW), 0)

	screen.DrawImage(sidImg, opts)
}
