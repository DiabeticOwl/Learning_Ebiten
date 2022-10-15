//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/stall/curtain_straight.png -output ./PNG/stall/curtain_straight.go -package stall -var TopCurtain_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/stall/curtain.png -output ./PNG/stall/curtain.go -package stall -var SideCurtain_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/stall/bg_green.png -output ./PNG/stall/bg_green.go -package stall -var BgGreen_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/stall/bg_wood.png -output ./PNG/stall/bg_wood.go -package stall -var BgWood_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/stall/water1.png -output ./PNG/stall/water1.go -package stall -var Water1_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/objects/duck_outline_target_white.png -output ./PNG/objects/duck_outline_target_white.go -package objects -var DuckOutlineTargetWhite_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/hud/crosshair_white_large.png -output ./PNG/hud/crosshair_white_large.go -package hud -var Crosshair_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./fonts/PenguinAttack/PenguinAttack.ttf -output ./fonts/PenguinAttack.go -package font -var PenguinAttack_font

package resources
