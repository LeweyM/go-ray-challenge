package main

import (
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
)

var t matrix.Matrix
var m matrix.Matrix
var n tuple.Tuple

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

func nNormal_atsPoint(arg1, arg2, arg3 float64) error {
	n = s.NormalAt(tuple.NewPoint(arg1, arg2, arg3))
	return nil
}

func nNormal_atsRootPoint(arg1, arg2, arg3, arg4, arg5, arg6 float64) error {
	n = s.NormalAt(tuple.NewPoint(math.Sqrt(arg1)/arg2, math.Sqrt(arg3)/arg4, math.Sqrt(arg5)/arg6))
	return nil
}

func nNormalizen() error {
	return ExpectEqualsTuple(n, n.Normalize())
}

func nVector(arg1, arg2, arg3 float64) error {
	return ExpectVector(&n, arg1, arg2, arg3)
}

func nRootVector(arg1, arg2, arg3, arg4, arg5, arg6 float64) error {
	return ExpectVector(&n, math.Sqrt(arg1)/arg2, math.Sqrt(arg3)/arg4, math.Sqrt(arg5)/arg6)
}

func mScalingRotation_z(arg1, arg2, arg3, arg4 float64) error {
	scale := matrix.NewScale(arg1, arg2, arg3)
	m = scale.Multiply(matrix.NewRotationZ(math.Pi/arg4))
	return nil
}

func nNormal_atsRootTwoPoint(arg1 float64) error {
	n = s.NormalAt(tuple.NewPoint(arg1, math.Sqrt(2)/2, -math.Sqrt(2)/2))
	return nil
}

func set_transformsM() error {
	s.SetTransform(m)
	return nil
}


func SphereContext(s *godog.ScenarioContext) {
	s.Step(`^s\.transform = identity_matrix$`, stransformIdentity_matrix)
	s.Step(`^set_transform\(s, t\)$`, set_transformsT)
	s.Step(`^s\.transform = t$`, stransformT)
	s.Step(`^set_transform\(s, scaling\((\d+), (\d+), (\d+)\)\)$`, set_transformsScaling)
	s.Step(`^set_transform\(s, translation\((\d+), (\d+), (\d+)\)\)$`, set_transformsTranslation)
	// normalAt
	s.Step(`^n ← normal_at\(s, point\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)\)$`, nNormal_atsPoint)
	s.Step(`^n ← normal_at\(s, point\(√(\d+)\/(\d+), √(\d+)\/(\d+), √(\d+)\/(\d+)\)\)$`, nNormal_atsRootPoint)
	s.Step(`^n = normalize\(n\)$`, nNormalizen)
	s.Step(`^n = vector\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, nVector)
	s.Step(`^n = vector\(√(\d+)\/(\d+), √(\d+)\/(\d+), √(\d+)\/(\d+)\)$`, nRootVector)
	s.Step(`^m ← scaling\(`+complexNum+`, `+complexNum+`, `+complexNum+`\) \* rotation_z\(π\/(\d+)\)$`, mScalingRotation_z)
	s.Step(`^n ← normal_at\(s, point\((\d+), √2\/2, -√2\/2\)\)$`, nNormal_atsRootTwoPoint)
	s.Step(`^set_transform\(s, m\)$`, set_transformsM)


}
