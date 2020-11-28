package world

import (
	"github/lewismetcalf/goRayChallenge/light"
	"github/lewismetcalf/goRayChallenge/matrix"
	object "github/lewismetcalf/goRayChallenge/object"
	"github/lewismetcalf/goRayChallenge/tuple"
)

type World struct {
	light   light.PointLight
	objects []object.Sphere
}

func (w *World) Objects() []object.Sphere {
	return w.objects
}

func (w *World) Light() (bool, light.PointLight) {
	return false, w.light
}

func (w *World) SetObjects(spheres ...object.Sphere) {
	w.objects = spheres
}

func (w *World) SetLight(light light.PointLight) {
	w.light = light
}

func NewWorld() *World {
	return &World{}
}

func NewDefaultWorld() *World {
	w := NewWorld()

	s1 := object.NewSphere()
	m1 := object.NewMaterial()
	m1.SetColor(tuple.NewColor(0.8, 1.0, 0.6))
	m1.SetDiffuse(0.7)
	m1.SetSpecular(0.2)
	s1.SetMaterial(*m1)

	s2 := object.NewSphere()
	s2.SetTransform(matrix.NewScale(0.5, 0.5, 0.5))

	w.SetObjects(*s1, *s2)

	light := light.NewPointLight(*tuple.NewPoint(-10, 10, -10), *tuple.NewColor(1, 1, 1))
	w.SetLight(*light)

	return w
}
