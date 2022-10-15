package main

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type customFnt struct {
	fnt     font.Face
	fntOpts *truetype.Options
}

func importFont(fntAsset []byte, fntOpts *truetype.Options) font.Face {
	tt, _ := truetype.Parse(fntAsset)

	return truetype.NewFace(tt, fntOpts)
}
