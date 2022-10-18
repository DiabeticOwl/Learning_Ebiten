package object

import (
	"math/rand"
	"shoot-the-duck/utils"

	objResources "shoot-the-duck/resources/PNG/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

type Duck struct {
	img     *ebiten.Image
	w       int
	h       int
	offsetX float64
	offsetY float64
	// Ducks will always go in the right direction horizontally.
	yDirection int
}

const (
	maxDuckOffsetY = 10
)

var (
	dImg   = utils.DecodeImage(objResources.DuckOutlineTargetWhite_png)
	dW, dH = dImg.Size()
)

func NewDuck() *Duck {
	return &Duck{
		img:        dImg,
		w:          dW,
		h:          dH,
		yDirection: up,
	}
}

func (d *Duck) Update() {
	// Each second (where tick/60==0) there is a 60% chance of a duck appearing
	// into the screen.
	if *GameTick%60 == 0 && rand.Float64() < 0.6 {
		Ducks[NewDuck()] = struct{}{}
	}

	for duck := range Ducks {
		duck.offsetX += 1.5
		duck.offsetY += 1.5 * float64(duck.yDirection)

		if int(duck.offsetX) >= GameScreenWidth {
			delete(Ducks, duck)
			continue
		}

		if rand.Float64() < 0.4 {
			duck.yDirection *= -1
		}

		if duck.offsetY >= maxDuckOffsetY {
			duck.yDirection = up
		}
	}
}

func (d *Duck) Draw(screen *ebiten.Image) {
	for duck := range Ducks {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(sidCurW)-float64(duck.w)/2, DuckInitialHeightPos-float64(duck.h))
		opts.GeoM.Translate(duck.offsetX, duck.offsetY)

		screen.DrawImage(duck.img, opts)
	}
}
