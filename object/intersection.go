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
	negateRayDirection := r.Direction().Negate()
	position := r.Position(i.t)
	return Computations{
		time:    i.t,
		object:  i.object,
		point:   position,
		eyeV:    *negateRayDirection,
		normalv: i.object.NormalAt(&position),
	}
}

func NewIntersection(object Sphere, t float64) *Intersection {
	return &Intersection{object: object, t: t}
}
