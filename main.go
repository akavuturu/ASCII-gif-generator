package main

import (
	classes "ASCII-gif-generator/packages"
)

var rocket []string = []string{
	"     /\\   ",
	"     ||    ",
	"     ||    ",
	"    /||\\  ",
	"   /:||:\\ ",
	"   |:||:|  ",
	"   |/||\\| ",
	"     **    ",
	"     **    "}

func main() {
	init_rocket := classes.FrameObject{X: 64, Y: 0, Width: 6, Height: 9, Ascii_str: rocket}
	// init_rocket.DrawFrameObject()

	init_frame := classes.Frame{Width: 128, Height: 18, Items: []classes.FrameObject{init_rocket}}

	init_frame.DrawFrame()
}
