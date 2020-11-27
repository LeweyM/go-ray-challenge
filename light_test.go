package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/light"
	"github/lewismetcalf/goRayChallenge/tuple"
)

var intensity tuple.Color
var position tuple.Tuple
var l light.PointLight

func intensityColor(arg1, arg2, arg3 float64) error {
	intensity = *tuple.NewColor(arg1, arg2, arg3)
	return nil
}

func lightPoint_lightpositionIntensity() error {
	l = *light.NewPointLight(position, intensity)
	return nil
}

func lightintensityIntensity() error {
	return ExpectColorEquals(intensity, l.Intensity())
}

func lightpositionPosition() error {
	return ExpectEqualsTuple(*l.Position(), &position)
}

func LightsContext(s *godog.ScenarioContext) {
	s.Step(`^intensity ← color\((\d+), (\d+), (\d+)\)$`, intensityColor)
	s.Step(`^light ← point_light\(position, intensity\)$`, lightPoint_lightpositionIntensity)
	s.Step(`^light\.intensity = intensity$`, lightintensityIntensity)
	s.Step(`^light\.position = position$`, lightpositionPosition)
}

func ExpectColorEquals(color tuple.Color, t2 *tuple.Color) error {
	return ExpectTrue(color.Equals(t2), fmt.Sprintf("Expected %v, got %p", intensity, l.Intensity()))
}
