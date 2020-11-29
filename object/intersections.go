package object

import (
	"sort"
)

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
	index := indexWithFirstNonNegativeTime(i.intersections)

	if i.Count() == index {
		return false, Intersection{}
	}

	closest := i.Get(index)
	if closest.Time() >= 0 {
		return true, closest
	} else {
		return false, Intersection{}
	}
}

func indexWithFirstNonNegativeTime(intersections []Intersection) int {
	var index = 0
	for range intersections {
		xs := intersections[index]
		if xs.Time() < 0 {
			index++
		}
	}
	return index
}
