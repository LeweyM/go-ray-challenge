package object

import (
	"github.com/gofrs/uuid"
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/ray"
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
)

type Sphere struct {
	id        uuid.UUID
	transform matrix.Matrix
}

func NewSphere() *Sphere {
	id, _ := uuid.NewV4()
	return &Sphere{transform: matrix.NewIdentityMatrix(), id: id}
}

func (s Sphere) Equals(other Sphere) bool {
	return s.id == other.id
}

func (s Sphere) Intersects(r ray.Ray) (bool, Intersections) {
	transform := s.Transform()
	ray := r.Transform(transform.Invert())

	sphereToRay := ray.Origin().Subtract(tuple.NewPoint(0, 0, 0))

	direction := ray.Direction()
	a := direction.Dot(&direction)
	b := ray.Direction().Dot(sphereToRay) * 2.0
	c := sphereToRay.Dot(sphereToRay) - 1.0

	descriminant := b*b - 4*a*c

	if descriminant < 0 {
		return false, *NewIntersections()
	}

	sqrtDescriminant := math.Sqrt(descriminant)

	t1 := (-b - sqrtDescriminant) / (2 * a)
	t2 := (-b + sqrtDescriminant) / (2 * a)

	return true, *NewIntersections(*NewIntersection(s, math.Min(t1, t2)), *NewIntersection(s, math.Max(t1, t2)))
}

func (s Sphere) Transform() matrix.Matrix {
	return s.transform
}

func (s *Sphere) SetTransform(t matrix.Matrix) {
	s.transform = t
}

func (s Sphere) NormalAt(point *tuple.Tuple) tuple.Tuple {
	inverseTransform := s.transform.Invert()
	objectPoint := inverseTransform.MultiplyTuple(point)
	objectNormal := objectPoint.Subtract(tuple.NewPoint(0, 0, 0))

	transposedInverseTransform := inverseTransform.Transpose()
	worldNormal := transposedInverseTransform.MultiplyTuple(objectNormal)
	worldNormal.W = 0
	return *worldNormal.Normalize()
}
