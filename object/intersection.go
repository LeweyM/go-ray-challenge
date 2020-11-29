package object

import (
	"github/lewismetcalf/goRayChallenge/ray"
)

type Intersection struct {
	object Sphere
	t      float64
}

func (i *Intersection) Time() float64 {
	return i.t
}

func (i *Intersection) Object() Sphere {
	return i.object
}

func (i *Intersection) PrepareComputations(r ray.Ray) Computations {
	var inside bool
	point := r.Position(i.t)
	eyeV := r.Direction().Negate()
	normalV := i.object.NormalAt(&point)
	if normalV.Dot(eyeV) < 0 {
		inside = true
		normalV = *normalV.Negate()
	} else {
		inside = false
	}

	return Computations{
		time:    i.t,
		object:  i.object,
		point:   point,
		eyeV:    *eyeV,
		normalv: normalV,
		inside:  inside,
	}
}

func NewIntersection(object Sphere, t float64) *Intersection {
	return &Intersection{object: object, t: t}
}
