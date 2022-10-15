package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type CrossHair struct {
	img         *ebiten.Image
	x           int
	y           int
	clicked     bool
	lastClickAt time.Time
}

func newCrosshair() *CrossHair {
	return &CrossHair{
		img: crosshairImage,
	}
}
