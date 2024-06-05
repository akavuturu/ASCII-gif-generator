package classes

import (
	"fmt"
	"strings"
)

type FrameObject struct {
	X, Y, Width, Height int
	Ascii_str           string
}

type Frame struct {
	Width, Height int
	Items         []FrameObject
}

type GIF struct {
	Frames []Frame
	Time   int
}

/* FUNCTIONS */
// func DrawFrameObject(item FrameObject) string {
// 	panic("unimplemented")
// }

func (frame Frame) DrawFrame() {
	objs := frame.Items
	lines := make([]string, frame.Height)

	//padding according to Width and Height

	for _, obj := range objs {
		// _ := obj.X - obj.Width/2
		// _ := obj.X + obj.Width/2
		new_line := strings.Repeat(" ", frame.Width)
		fmt.Print(obj.X)

		// if i <= obj.Height {
		// 	new_line[obj_start : obj_end+1] = obj.Ascii_str[obj.Height-i]
		// }
		lines = append(lines, new_line)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func (g GIF) Animate() []string {
	panic("unimplemented")
}
