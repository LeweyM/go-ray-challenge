package main

type Color struct {
	t Tuple
}

func NewColor(r, g, b float64) *Color {
	return &Color{t: Tuple{r, g, b, 0.0}}
}

func (c *Color) red() float64 {
	return c.t.x
}

func (c *Color) green() float64 {
	return c.t.y
}

func (c *Color) blue() float64 {
	return c.t.z
}
