package object

import (
	"fmt"
	"image/color"
	"shoot-the-duck/utils"

	fntResources "shoot-the-duck/resources/fonts"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type score struct{}

var (
	penguinAttackFnt = utils.ImportFont(fntResources.PenguinAttack_font,
		&truetype.Options{Size: 60})
)

func NewScore() Object {
	return &score{}
}

func (bg *score) Update() {}

func (bg *score) Draw(screen *ebiten.Image) {
	score := fmt.Sprintf("Score: %d", *GameScore)

	text.Draw(screen, score, penguinAttackFnt, 10,
		int(float64(GameScreenHeight)*.95), color.Black)
}
