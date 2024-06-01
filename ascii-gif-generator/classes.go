package classes

type FrameObject struct {
    x, y, width, height int
    ascii_representation string
}

type Frame struct {
    width, height int
    items []FrameObject
}

type GIF struct {
    frames []Frame
    time int
}
