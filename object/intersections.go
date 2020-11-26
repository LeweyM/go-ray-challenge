package object

import "sort"

type Intersections struct {
	intersections []Intersection
}

func NewIntersections(intersections ...Intersection) *Intersections {
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].Time() < intersections[j].Time()
	})
	return &Intersections{intersections: intersections}
}

func (i *Intersections) Get(index int) Intersection {
	return i.intersections[index]
}

func (i *Intersections) Count() int {
	return len(i.intersections)
}

func (i *Intersections) Hit() (bool, Intersection) {
	closest := i.Get(0)
	if closest.Time() >= 0 {
		return true, closest
	} else {
		return false, Intersection{}
	}
}
