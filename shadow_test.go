package main

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github/lewismetcalf/goRayChallenge/object"
)

var inShadow bool

func in_shadowTrue() error {
	inShadow = true
	return nil
}

func resultLightingmLightPositionEyevNormalvIn_shadow() error {
	eyev := tuples["eyev"]
	normalv := tuples["normalv"]
	result = mat.Lighting(&l, &position, eyev, normalv, inShadow)
	return nil
}

func is_shadowedwPIsFalse() error {
	p := tuples["p"]
	return ExpectFalse(w.IsShadowed(p), "Expected p NOT to be in shadow")
}

func is_shadowedwPIsTrue() error {
	p := tuples["p"]
	return ExpectTrue(w.IsShadowed(p), "Expected p to be in shadow")
}

func s1IsAddedToW() error {
	w.AddObject(*s1)
	return nil
}

func s2IsAddedToW() error {
	w.AddObject(*s2)
	return nil
}

func iIntersectionS2(arg1 float64) error {
	i = *object.NewIntersection(*s2, arg1)
	return nil
}

func s1Sphere() error {
	s1 = object.NewSphere()
	return nil
}

func compsover_pointzEPSILON(arg1 int) error {
	return ExpectTrue(comps.OverPoint().Z < -0.0001/2, "Expected overpoint.z to be less")
}

func compspointzCompsover_pointz() error {
	return ExpectTrue(comps.Point().Z > comps.OverPoint().Z, "Expected overpoint.z to be greater than")
}

func shapeSphereWith(arg1 *messages.PickleStepArgument_PickleTable) error {
	shape = *parseSphereFromTable(arg1)
	return nil
}

func ShadowFeatureContext(s *godog.ScenarioContext) {
	s.Step(`^in_shadow ← true$`, in_shadowTrue)
	s.Step(`^result ← lighting\(m, light, position, eyev, normalv, in_shadow\)$`, resultLightingmLightPositionEyevNormalvIn_shadow)
	s.Step(`^is_shadowed\(w, p\) is false$`, is_shadowedwPIsFalse)
	s.Step(`^is_shadowed\(w, p\) is true$`, is_shadowedwPIsTrue)

	s.Step(`^i ← intersection\((\d+), s2\)$`, iIntersectionS2)
	s.Step(`^s1 is added to w$`, s1IsAddedToW)
	s.Step(`^s2 is added to w$`, s2IsAddedToW)
	s.Step(`^s1 ← sphere\(\)$`, s1Sphere)

	s.Step(`^comps\.over_point\.z < -EPSILON\/(\d+)$`, compsover_pointzEPSILON)
	s.Step(`^comps\.point\.z > comps\.over_point\.z$`, compspointzCompsover_pointz)
	s.Step(`^shape ← sphere\(\) with:$`, shapeSphereWith)
}
