package object

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

func NewIntersection(object Sphere, t float64) *Intersection {
	return &Intersection{object: object, t: t}
}
