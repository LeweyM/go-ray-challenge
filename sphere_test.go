package main

import (
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/matrix"
)

var t matrix.Matrix

func stransformIdentity_matrix() error {
	return expectEqualMatrices(s.Transform(), matrix.NewIdentityMatrix())
}

func set_transformsT() error {
	s.SetTransform(t)
	return nil
}

func stransformT() error {
	return expectEqualMatrices(s.Transform(), t)
}

func set_transformsScaling(arg1, arg2, arg3 float64) error {
	s.SetTransform(matrix.NewScale(arg1, arg2, arg3))
	return nil
}

func set_transformsTranslation(arg1, arg2, arg3 float64) error {
	s.SetTransform(matrix.NewTranslation(arg1, arg2, arg3))
	return nil
}

func SphereContext(s *godog.ScenarioContext) {
	s.Step(`^s\.transform = identity_matrix$`, stransformIdentity_matrix)
	s.Step(`^set_transform\(s, t\)$`, set_transformsT)
	s.Step(`^s\.transform = t$`, stransformT)
	s.Step(`^set_transform\(s, scaling\((\d+), (\d+), (\d+)\)\)$`, set_transformsScaling)
	s.Step(`^set_transform\(s, translation\((\d+), (\d+), (\d+)\)\)$`, set_transformsTranslation)
}
