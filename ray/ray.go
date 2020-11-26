package ray

import (
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/tuple"
)

type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

func NewRay(o *tuple.Tuple, d *tuple.Tuple) Ray {
	return Ray{origin: *o, direction: *d}
}

func (r *Ray) Origin() tuple.Tuple {
	return r.origin
}

func (r *Ray) Direction() tuple.Tuple {
	return r.direction
}

func (r *Ray) Position(t float64) tuple.Tuple {
	return *r.origin.Add(r.direction.Multiply(t))
}

func (r *Ray) Transform(m matrix.Matrix) Ray {
	newDir := m.MultiplyTuple(&r.direction)
	newOrigin := m.MultiplyTuple(&r.origin)
	return NewRay(&newOrigin, &newDir)
}
