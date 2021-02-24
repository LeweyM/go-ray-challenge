package main

import (
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/light"
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
)

func lightPoint_lightpointColor(arg1, arg2, arg3, arg4, arg5, arg6 float64) error {
	l = *light.NewPointLight(*tuple.NewPoint(arg1, arg2, arg3), *tuple.NewColor(arg4, arg5, arg6))
	return nil
}

var result tuple.Color

func resultColor(arg1, arg2, arg3 float64) error {
	return ExpectColor(&result, arg1, arg2, arg3)
}

func resultLightingmLightPositionEyevNormalv() error {
	eyev := tuples["eyev"]
	normalv := tuples["normalv"]
	position := tuples["position"]
	result = mat.Lighting(&l, position, eyev, normalv, false)
	return nil
}

func eyevNegativeVector(arg1, arg2, arg3, arg4, arg5 float64) error {
	tuples["eyev"] = tuple.NewVector(arg1, math.Sqrt(arg2)/arg3, -math.Sqrt(arg4)/arg5)
	return nil
}

func eyevNegativeNegativeVector(arg1, arg2, arg3, arg4, arg5 float64) error {
	tuples["eyev"] = tuple.NewVector(arg1, -math.Sqrt(arg2)/arg3, -math.Sqrt(arg4)/arg5)
	return nil
}

func LightingContext(s *godog.ScenarioContext) {
	s.Step(`^light ← point_light\(point\(`+complexNum+`, `+complexNum+`, `+complexNum+`\), color\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)\)$`, lightPoint_lightpointColor)
	s.Step(`^result = color\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, resultColor)
	s.Step(`^result ← lighting\(m, light, position, eyev, normalv\)$`, resultLightingmLightPositionEyevNormalv)
	s.Step(`^eyev ← vector\((\d+), √(\d+)\/(\d+), -√(\d+)\/(\d+)\)$`, eyevNegativeVector)
	s.Step(`^eyev ← vector\((\d+), -√(\d+)\/(\d+), -√(\d+)\/(\d+)\)$`, eyevNegativeNegativeVector)
}
