package canvas

import "goRayChallenge/tuple"

type Canvas struct {
	width, height int
}

func (c Canvas) Height() int {
	return c.height
}

func (c Canvas) Width() int {
	return c.width
}

func (c Canvas) PixelAt(x int, y int) tuple.Color {
	return tuple.Color{T: tuple.Tuple{
		X: 0,
		Y: 0,
		Z: 0,
		W: 0,
	}}
}

func NewCanvas(width int, height int) *Canvas {
	return &Canvas{width: width, height: height}
}


