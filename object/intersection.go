package object

import "math"

type Intersection struct {
	intersections []float64
}

func (i Intersection) Count() int {
	return len(i.intersections)
}

func (i *Intersection) Get() (bool, float64, float64) {
	t1 := math.Min(i.intersections[0], i.intersections[1])
	t2 := math.Max(i.intersections[0], i.intersections[1])
	return i.Count() > 0, t1, t2
}
