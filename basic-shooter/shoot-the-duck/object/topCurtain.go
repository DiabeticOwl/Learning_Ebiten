package object

import (
	"math"
	"shoot-the-duck/utils"

	pngResources "shoot-the-duck/resources/PNG/stall"

	"github.com/hajimehoshi/ebiten/v2"
)

type topCurtain struct{}

var (
	topCurImg  = utils.DecodeImage(pngResources.TopCurtain_png)
	topCurW, _ = topCurImg.Size()
)

func NewTopCurtain() Object {
	return &topCurtain{}
}

func (topCur *topCurtain) Update() {}

func (topCur *topCurtain) Draw(screen *ebiten.Image) {
	sW, _ := screen.Size()

	nDraws := int(math.Ceil(float64(sW) / float64(topCurW)))

	for i := 0; i < nDraws; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(topCurW*i), 0)

		screen.DrawImage(topCurImg, opts)
	}
}
