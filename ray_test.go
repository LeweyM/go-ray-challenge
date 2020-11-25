package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/object"
	"github/lewismetcalf/goRayChallenge/ray"
	"github/lewismetcalf/goRayChallenge/tuple"
)

var r ray.Ray
var s object.Sphere
var xs object.Intersection

func rRayoriginDirection() error {
	o := tuples["origin"]
	d := tuples["direction"]
	r = ray.NewRay(o, d)
	return nil
}

func rdirectionDirection() error {
	d := tuples["direction"]
	direction := r.Direction()
	return ExpectEqualsTuple(direction, d)
}

func roriginOrigin() error {
	o := tuples["origin"]
	origin := r.Origin()
	return ExpectEqualsTuple(origin, o)
}

func positionrPoint(t, arg2, arg3, arg4 float64) error {
	pos := r.Position(t)
	return ExpectPoint(&pos, arg2, arg3, arg4)
}

func rRaypointVector(arg1, arg2, arg3, arg4, arg5, arg6 float64) error {
	r = ray.NewRay(tuple.NewPoint(arg1, arg2, arg3), tuple.NewVector(arg4, arg5, arg6))
	return nil
}

func sSphere() error {
	s = *object.NewSphere()
	return nil
}

func xsEquals(arg1 int, arg2 float64) error {
	_, t1, t2 := xs.Get()
	if arg1 == 0 {
		return ExpectFloatEquals(t1, arg2)
	} else {
		return ExpectFloatEquals(t2, arg2)
	}
}

func xsIntersectsR() error {
	xs = s.Intersects(r)
	return nil
}

func xscount(arg1 int) error {
	return ExpectTrue(xs.Count() == arg1, fmt.Sprintf("expected %v, to equal %d", xs.Count(), arg1))
}

func RayContext(s *godog.ScenarioContext) {
	s.Step(`^r ← ray\(origin, direction\)$`, rRayoriginDirection)
	s.Step(`^r\.direction = direction$`, rdirectionDirection)
	s.Step(`^r\.origin = origin$`, roriginOrigin)
	s.Step(`^position\(r, (\-?\d*\.?\d)\) = point\((\-?\d*\.?\d), (\-?\d*\.?\d), (\-?\d*\.?\d)\)$`, positionrPoint)
	s.Step(`^r ← ray\(point\((\d+), (\d+), (\-?\d+)\), vector\((\d+), (\d+), (\d+)\)\)$`, rRaypointVector)
	// sphere
	s.Step(`^s ← sphere\(\)$`, sSphere)
	s.Step(`^xs\[(\d+)\] = (\-?\d*\.?\d)$`, xsEquals)
	s.Step(`^xs ← intersect\(s, r\)$`, xsIntersectsR)
	s.Step(`^xs\.count = (\d+)$`, xscount)
}

func ExpectEqualsTuple(origin tuple.Tuple, o *tuple.Tuple) error {
	return ExpectTrue(origin.Equals(o), fmt.Sprintf("Expected %v to equal %p", origin, o))
}

