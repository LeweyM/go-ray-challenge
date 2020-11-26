package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/object"
	"github/lewismetcalf/goRayChallenge/ray"
	"github/lewismetcalf/goRayChallenge/tuple"
)

var r ray.Ray
var r2 ray.Ray

var s object.Sphere

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
	intersection := xs.Get(arg1)
	return ExpectFloatEquals(intersection.Time(), arg2)
}

func xsIntersectsR() error {
	_, xs = s.Intersects(r)
	return nil
}

func xscount(arg1 int) error {
	return ExpectTrue(xs.Count() == arg1, fmt.Sprintf("expected %v, to equal %d", xs.Count(), arg1))
}

func xsObjectS(arg1 int) error {
	intersection := xs.Get(arg1)
	return ExpectObjectEquals(intersection.Object(), s)
}

func r2DirectionVector(arg1, arg2, arg3 float64) error {
	direction := r2.Direction()
	return ExpectVector(&direction, arg1, arg2, arg3)
}

func r2OriginPoint(arg1, arg2, arg3 float64) error {
	origin := r2.Origin()
	return ExpectPoint(&origin, arg1, arg2, arg3)
}

func rTransformrM() error {
	m := matrices["m"]
	r2 = r.Transform(m)
	return nil
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
	s.Step(`^xs\[(\d+)\]\.object = s$`, xsObjectS)
	// transform
	s.Step(`^r2 ← transform\(r, m\)$`, rTransformrM)
	s.Step(`^r2\.direction = vector\((\d+), (\d+), (\d+)\)$`, r2DirectionVector)
	s.Step(`^r2\.origin = point\((\d+), (\d+), (\d+)\)$`, r2OriginPoint)

}

func ExpectEqualsTuple(origin tuple.Tuple, o *tuple.Tuple) error {
	return ExpectTrue(origin.Equals(o), fmt.Sprintf("Expected %v to equal %p", origin, o))
}

const complexNum = `(\-?\d*\.?\d)`
