package object

import (
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/ray"
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
)

type Sphere struct {
	transform matrix.Matrix
	material  *Material
}

func NewSphere() *Sphere {
	return &Sphere{
		transform: matrix.NewIdentityMatrix(),
		material:  NewMaterial(),
	}
}

func (s *Sphere) Equals(other Sphere) bool {
	return s.transform.Equals(other.transform) && s.material.Equals(*other.material)
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

func (s *Sphere) Transform() matrix.Matrix {
	return s.transform
}

func (s *Sphere) SetTransform(t matrix.Matrix) {
	s.transform = t
}

func (s *Sphere) NormalAt(point *tuple.Tuple) tuple.Tuple {
	inverseTransform := s.transform.Invert()
	objectPoint := inverseTransform.MultiplyTuple(point)
	objectNormal := objectPoint.Subtract(tuple.NewPoint(0, 0, 0))

	transposedInverseTransform := inverseTransform.Transpose()
	worldNormal := transposedInverseTransform.MultiplyTuple(objectNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}

func (s *Sphere) Material() *Material {
	return s.material
}

func (s *Sphere) SetMaterial(mat *Material) {
	s.material = mat
}
