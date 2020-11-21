package canvas

import (
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
	"strconv"
	"strings"
)

type Canvas struct {
	width, height int
	pixels []*tuple.Color
}

func NewCanvas(width int, height int) *Canvas {
	pixels := make([]*tuple.Color, height*width)
	black := tuple.Color{T: tuple.Tuple{
		X: 0,
		Y: 0,
		Z: 0,
		W: 0,
	}}
	for i := 0; i < width*height; i++ {
		pixels[i] = &black
	}
	return &Canvas{width: width, height: height, pixels: pixels}
}

func (c *Canvas) Height() int {
	return c.height
}

func (c *Canvas) Width() int {
	return c.width
}

func (c *Canvas) PixelAt(x int, y int) *tuple.Color {
	return c.pixels[c.index(x, y)]
}

func (c *Canvas) WritePixel(x, y int, color *tuple.Color) {
	c.pixels[c.index(x, y)] = color
}

func (c *Canvas) index(x int, y int) int {
	return x + (y * c.width)
}

func (c *Canvas) ToPPM() string {
	response := "P3\n5 3\n255"
	for y := 0; y < c.height; y++ {
		response += "\n"
		var row []string
		for x := 0; x < c.width; x++ {
			r := toPrintedPixel(c.PixelAt(x, y).Red())
			g := toPrintedPixel(c.PixelAt(x, y).Green())
			b := toPrintedPixel(c.PixelAt(x, y).Blue())
			row = append(row, r)
			row = append(row, g)
			row = append(row, b)

		}
		response += splitRow([]byte(strings.Join(row, " ")))
	}
	return response
}

func toPrintedPixel(p float64) string {
	clamped := math.Max(math.Min(p, 1.0), 0.0)
	v := int(math.Round(clamped * 255))
	return strconv.Itoa(v)
}

func splitRow(row []byte) string {
	if len(row) <= 70 {
		return string(row)
	}
	i := 70
	for row[i] != ' ' {
		i--
	}
	row[i] = '\n'
	return string(row)
}
