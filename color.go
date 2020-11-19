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

func (c *Color) add(c2 *Color) *Color {
	return NewColor(c.red() + c2.red(), c.green() + c2.green(), c.blue() + c2.blue())
}

func (c *Color) subtract(c2 *Color) *Color {
	return NewColor(c.red() - c2.red(), c.green() - c2.green(), c.blue() - c2.blue())
}

func (c *Color) multiplyScalar(s float64) *Color {
	return NewColor(c.red() * s, c.green() * s, c.blue() * s)
}

func (c *Color) multiply(c2 *Color) *Color {
	return NewColor(c.red() * c2.red(), c.green() * c2.green(), c.blue() * c2.blue())
}
