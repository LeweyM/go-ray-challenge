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

var comps object.Computations
var shape object.Sphere

func compsPrepare_computationsiR() error {
	comps = i.PrepareComputations(r)
	return nil
}

func compstIt() error {
	return ExpectFloatEquals(comps.Time(), i.Time())
}

func iIntersectionShape(arg1 float64) error {
	i = *object.NewIntersection(shape, arg1)
	return nil
}

func compsobjectIobject() error {
	return ExpectEqualsSpheres(comps.Object(), i.Object())
}

func compspointPoint(arg1, arg2, arg3 float64) error {
	point := comps.Point()
	return ExpectPoint(&point, arg1, arg2, arg3)
}

func compseyevVector(arg1, arg2, arg3 float64) error {
	vector := comps.EyeVector()
	return ExpectVector(&vector, arg1, arg2, arg3)
}

func compsnormalvVector(arg1, arg2, arg3 float64) error {
	vector := comps.NormalVector()
	return ExpectVector(&vector, arg1, arg2, arg3)
}

func shapeSphere() error {
	shape = *object.NewSphere()
	return nil
}

func compsinsideTrue() error {
	return ExpectTrue(comps.IsInside(), fmt.Sprintf("Expected to be inside"))
}

func compsinsideFalse() error {
	return ExpectFalse(comps.IsInside(), fmt.Sprintf("Expected to be inside"))
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
	//precomputation
	s.Step(`^comps ← prepare_computations\(i, r\)$`, compsPrepare_computationsiR)
	s.Step(`^comps\.eyev = vector\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, compseyevVector)
	s.Step(`^comps\.normalv = vector\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, compsnormalvVector)
	s.Step(`^comps\.object = i\.object$`, compsobjectIobject)
	s.Step(`^comps\.point = point\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, compspointPoint)
	s.Step(`^comps\.t = i\.t$`, compstIt)
	s.Step(`^i ← intersection\((\d+), shape\)$`, iIntersectionShape)
	s.Step(`^shape ← sphere\(\)$`, shapeSphere)
	s.Step(`^comps\.inside = true$`, compsinsideTrue)
	s.Step(`^comps\.inside = false$`, compsinsideFalse)

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
