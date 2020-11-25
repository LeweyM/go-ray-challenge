package object

import (
	"github/lewismetcalf/goRayChallenge/ray"
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
)

type Sphere struct {

}

func (s Sphere) Intersects(r ray.Ray) Intersection {

	sphereToRay := r.Origin().Subtract(tuple.NewPoint(0, 0, 0))

	direction := r.Direction()
	a := direction.Dot(&direction)
	b := r.Direction().Dot(sphereToRay) * 2.0
	c := sphereToRay.Dot(sphereToRay) - 1.0

	descriminant := b*b - 4*a*c

	if descriminant < 0 {
		return Intersection{intersections: []float64{}}
	}

	sqrtDescriminant := math.Sqrt(descriminant)

	t1 := (-b - sqrtDescriminant) / (2 * a)
	t2 := (-b + sqrtDescriminant) / (2 * a)

	return Intersection{intersections: []float64{t1, t2}}
}

func NewSphere() *Sphere {
	return &Sphere{}
}


