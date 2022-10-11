//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/stall/curtain_straight.png -output ./PNG/Stall/curtain_straight.go -package stall -var TopCurtain_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/stall/curtain.png -output ./PNG/Stall/curtain.go -package stall -var SideCurtain_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/stall/bg_green.png -output ./PNG/Stall/bg_green.go -package stall -var BgGreen_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ./PNG/stall/bg_wood.png -output ./PNG/Stall/bg_wood.go -package stall -var BgWood_png

package resources
