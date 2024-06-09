package classes

import (
	"fmt"
)

type FrameObject struct {
	X, Y, Width, Height int
	Ascii_str           []string
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

func (obj FrameObject) DrawFrameObject() {
	for _, line := range obj.Ascii_str {
		fmt.Println(line)
	}
}

func (frame Frame) DrawFrame() {
	objs := frame.Items
	frame_ascii := make([][]byte, frame.Height)
	for i := range frame_ascii {
		frame_ascii[i] = make([]byte, frame.Width)
	}

	for i := range frame_ascii {
		for j := range frame_ascii[i] {
			frame_ascii[i][j] = ' '

			for _, obj := range objs {
				fmt.Print(obj)
			}
		}
	}
	//padding according to Width and Height

	// for i, obj := range objs {
	// 	lbound := obj.X - obj.Width/2
	// 	rbound := obj.X + obj.Width/2
	// 	new_line := []byte(strings.Repeat(" ", frame.Width))
	// 	fmt.Print(new_line)

	// 	if i <= obj.Height {
	// 		copy(new_line[lbound:rbound+1], obj.Ascii_str[obj.Height-i-1])
	// 	}
	// 	lines = append(lines, string(new_line))
	// }

	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
}

func (g GIF) Animate() []string {
	panic("unimplemented")
}
