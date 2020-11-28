package object

import (
	"github/lewismetcalf/goRayChallenge/tuple"
)

type Computations struct {
	time float64
	object Sphere
	point tuple.Tuple
	eyeV tuple.Tuple
	normalv tuple.Tuple
}

func (c *Computations) Time() float64 {
	return c.time
}

func (c *Computations) Object() Sphere {
	return c.object
}

func (c *Computations) Point() tuple.Tuple {
	return c.point
}

func (c *Computations) EyeVector() tuple.Tuple {
	return c.eyeV
}

func (c *Computations) NormalVector() tuple.Tuple {
	return c.normalv
}
