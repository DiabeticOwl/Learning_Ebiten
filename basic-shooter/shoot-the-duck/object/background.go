package object

import (
	"shoot-the-duck/utils"

	pngResources "shoot-the-duck/resources/PNG/stall"

	"github.com/hajimehoshi/ebiten/v2"
)

type background struct{}

var (
	bgImg    = utils.DecodeImage(pngResources.BgGreen_png)
	bgW, bgH = bgImg.Size()
)

func NewBackground() Object {
	return &background{}
}

func (bg *background) Update() {}

func (bg *background) Draw(screen *ebiten.Image) {
	sW, sH := screen.Size()

	bgWDelta := float64(bgW - sW)
	newbgWRatio := bgWDelta / float64(bgW)
	bgHDelta := float64(bgH - sH)
	newbgHRatio := bgHDelta / float64(bgH)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1-newbgWRatio, 1-newbgHRatio)

	screen.DrawImage(bgImg, opts)
}
