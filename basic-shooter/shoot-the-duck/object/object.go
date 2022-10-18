package object

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	Draw(screen *ebiten.Image)
	Update()
}

var (
	Ducks            = make(map[*Duck]struct{})
	GameScore        *int
	GameScreenHeight int
	GameScreenWidth  int
	GameTick         *int
)

const (
	right int = 1
	left  int = -1
	down  int = 1
	up    int = -1
)
