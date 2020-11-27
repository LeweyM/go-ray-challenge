package tuple

type Color struct {
	T Tuple
}

func NewColor(r, g, b float64) *Color {
	return &Color{T: Tuple{r, g, b, 0.0}}
}

func (c *Color) Red() float64 {
	return c.T.X
}

func (c *Color) Green() float64 {
	return c.T.Y
}

func (c *Color) Blue() float64 {
	return c.T.Z
}

func (c *Color) Add(c2 *Color) *Color {
	return &Color{T: *c.T.Add(&c2.T)}
}

func (c *Color) Subtract(c2 *Color) *Color {
	return &Color{T: *c.T.Subtract(&c2.T)}
}

func (c *Color) MultiplyScalar(s float64) *Color {
	return &Color{T: *c.T.Multiply(s)}
}

func (c *Color) Multiply(c2 *Color) *Color {
	return NewColor(c.Red() * c2.Red(), c.Green() * c2.Green(), c.Blue() * c2.Blue())
}

func (c *Color) Equals(other *Color) bool {
	return c.T.Equals(&other.T)
}
