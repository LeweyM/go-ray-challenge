package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/object"
)

var xs object.Intersections
var i object.Intersection
var iIsSomething bool
var i1 object.Intersection
var i2 object.Intersection
var i3 object.Intersection
var i4 object.Intersection

func iIntersectionS(arg1 float64) error {
	i = *object.NewIntersection(s, arg1)
	return nil
}

func i1IntersectionS(arg1 float64) error {
	i1 = *object.NewIntersection(s, arg1)
	return nil
}

func i2IntersectionS(arg1 float64) error {
	i2 = *object.NewIntersection(s, arg1)
	return nil
}

func i3IntersectionS(arg1 float64) error {
	i3 = *object.NewIntersection(s, arg1)
	return nil
}

func i4IntersectionS(arg1 float64) error {
	i4 = *object.NewIntersection(s, arg1)
	return nil
}

func iobjectS() error {
	return ExpectObjectEquals(i.Object(), s)
}

func it(arg1 float64) error {
	return ExpectFloatEquals(i.Time(), arg1)
}

func ixsIntersectionsI1I2() error {
	xs = *object.NewIntersections(i1, i2)
	return nil
}

func ixsIntersectionsI1I2I3I4() error {
	xs = *object.NewIntersections(i1, i2, i3, i4)
	return nil
}

func ixsT(i int, t float64) error {
	intersection := xs.Get(i)
	return ExpectFloatEquals(intersection.Time(), t)
}

func ixscount(arg1 int) error {
	return ExpectIntEquals(arg1, xs.Count())
}

func iHitxs() error {
	iIsSomething, i = xs.Hit()
	return nil
}

func xsIntersectionsiI() error {
	xs = *object.NewIntersections(i1, i2)
	return nil
}

func iIsNothing() error {
	return ExpectFalse(iIsSomething, "Expected i to be nothing")
}

func IntersectionFeatureContext(s *godog.ScenarioContext) {
	s.Step(`^i ← intersection\(`+complexNum+`, s\)$`, iIntersectionS)
	s.Step(`^i\.object = s$`, iobjectS)
	s.Step(`^i\.t = `+complexNum+`$`, it)
	s.Step(`^i1 ← intersection\(`+complexNum+`, s\)$`, i1IntersectionS)
	s.Step(`^i2 ← intersection\(`+complexNum+`, s\)$`, i2IntersectionS)
	s.Step(`^i3 ← intersection\(`+complexNum+`, s\)$`, i3IntersectionS)
	s.Step(`^i4 ← intersection\(`+complexNum+`, s\)$`, i4IntersectionS)
	s.Step(`^xs ← intersections\(i1, i2\)$`, ixsIntersectionsI1I2)
	s.Step(`^xs ← intersections\(i1, i2, i3, i4\)$`, ixsIntersectionsI1I2I3I4)
	s.Step(`^xs\.count = (\d+)$`, ixscount)
	// hit
	s.Step(`^i ← hit\(xs\)$`, iHitxs)
	s.Step(`^xs ← intersections\(i2, i1\)$`, xsIntersectionsiI)
	s.Step(`^i is nothing$`, iIsNothing)

	s.BeforeScenario(func(sc *godog.Scenario) {
		xs = *object.NewIntersections()
	})
}

func ExpectIntEquals(arg1, arg2 int) error {
	return ExpectTrue(arg1 == arg2, fmt.Sprintf("Expected %d, got %d", arg1, arg2))
}

func ExpectObjectEquals(a, b object.Sphere) error {
	return ExpectTrue(a.Equals(b), fmt.Sprintf("Expected %v, got %v.", a, b))
}
