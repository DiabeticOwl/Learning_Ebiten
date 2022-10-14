package main

import "github.com/hajimehoshi/ebiten/v2"

type CrossHair struct {
	img     *ebiten.Image
	x       int
	y       int
	clicked bool
}

func newCrosshair() *CrossHair {
	return &CrossHair{
		img: crosshairImage,
	}
}
