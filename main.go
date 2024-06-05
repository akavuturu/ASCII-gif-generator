package main

import (
	"ASCII-gif-generator/classes"
)

const rocket = `     /\
     ||
     ||
    /||\
   /:||:\
   |:||:|
   |/||\|
     **
     **`

func main() {
	init_rocket := classes.FrameObject{X: 64, Y: 0, Width: 6, Height: 9, Ascii_str: rocket}

	init_frame := classes.Frame{Width: 128, Height: 24, Items: []classes.FrameObject{init_rocket}}
	// center rocket in frame and print all whitespace around
	init_frame.DrawFrame()
}
