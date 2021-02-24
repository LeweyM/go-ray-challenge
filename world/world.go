package world

import (
	"github/lewismetcalf/goRayChallenge/light"
	"github/lewismetcalf/goRayChallenge/matrix"
	object "github/lewismetcalf/goRayChallenge/object"
	"github/lewismetcalf/goRayChallenge/ray"
	"github/lewismetcalf/goRayChallenge/tuple"
	"sort"
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

func (w *World) Intersect(r ray.Ray) object.Intersections {
	xss := []object.Intersection{}
	for _, o := range w.Objects() {
		ok, intersections := o.Intersects(r)
		if ok {
			xss = append(xss, intersections.Get(0))
			xss = append(xss, intersections.Get(1))
		}
	}
	sort.Slice(xss, func(i, j int) bool {
		return xss[i].Time() < xss[j].Time()
	})
	return *object.NewIntersections(xss...)
}

func (w *World) ShadeHit(comps object.Computations) *tuple.Color {
	sphere := comps.Object()
	material := sphere.Material()
	eyeVector := comps.EyeVector()
	point := comps.Point()
	normalVector := comps.NormalVector()
	overPoint := comps.OverPoint()
	lighting := material.Lighting(&w.light, &point, &eyeVector, &normalVector, w.IsShadowed(&overPoint))
	return &lighting
}

func (w *World) ColorAt(r ray.Ray) *tuple.Color {
	intersections := w.Intersect(r)
	hasHit, intersection := intersections.Hit()
	if !hasHit {
		return tuple.NewColor(0,0,0)
	} else {
		computations := intersection.PrepareComputations(r)
		return w.ShadeHit(computations)
	}
}

func (w *World) IsShadowed(p *tuple.Tuple) bool {
	v := w.light.Position().Subtract(p)
	direction := v.Normalize()
	distance := v.Magnitude()
	r := ray.NewRay(p, &direction)
	intersections := w.Intersect(r)
	hit, xs := intersections.Hit()
	return hit && xs.Time() < distance
}

func (w *World) AddObject(s2 object.Sphere) {
	w.objects = append(w.objects, s2)
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
	s1.SetMaterial(m1)

	s2 := object.NewSphere()
	s2.SetTransform(matrix.NewScale(0.5, 0.5, 0.5))

	w.SetObjects(*s1, *s2)

	light := light.NewPointLight(*tuple.NewPoint(-10, 10, -10), *tuple.NewColor(1, 1, 1))
	w.SetLight(*light)

	return w
}
