package utils

import (
	"bytes"
	"image"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

func DecodeImage(imgSlice []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imgSlice))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func ImportFont(fntAsset []byte, fntOpts *truetype.Options) font.Face {
	tt, _ := truetype.Parse(fntAsset)

	return truetype.NewFace(tt, fntOpts)
}
