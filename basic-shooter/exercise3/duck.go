package main

import "github.com/hajimehoshi/ebiten/v2"

type Duck struct {
	id      int
	img     *ebiten.Image
	w       int
	h       int
	offsetX float64
	offsetY float64
	// Ducks will always go in the right direction horizontally.
	yDirection int
}

func newDuck(id int) *Duck {
	dW, dH := duckOutlineTargetWhite.Size()

	return &Duck{
		id:         id,
		img:        duckOutlineTargetWhite,
		w:          dW,
		h:          dH,
		yDirection: 1,
	}
}
