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
	w := strconv.Itoa(c.width)
	h := strconv.Itoa(c.height)
	response := "P3\n"+w+" "+h+"\n255"
	for y := 0; y < c.height; y++ {
		response += "\n"
		var row []string
		for x := 0; x < c.width; x++ {
			row = append(row, toPrintedPixel(c.PixelAt(x, y).Red()))
			row = append(row, toPrintedPixel(c.PixelAt(x, y).Green()))
			row = append(row, toPrintedPixel(c.PixelAt(x, y).Blue()))
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
	j := 70
	for j < len(row) {
		for row[j] != ' ' {
			j--
		}
		row[j] = '\n'
		j += 70
	}
	return string(row)
}
