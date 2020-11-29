package main

import (
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/matrix"
)

func tView_transformfromToUp() error {
	from := tuples["from"]
	to := tuples["to"]
	up := tuples["up"]
	t = matrix.ViewTransform(*from, *to, *up)
	return nil
}

func tScaling(arg1, arg2, arg3 float64) error {
	scale := matrix.NewScale(arg1, arg2, arg3)
	return expectEqualMatrices(t, scale)
}

func tTranslation(arg1, arg2, arg3 float64) error {
	translation := matrix.NewTranslation(arg1, arg2, arg3)
	return expectEqualMatrices(t, translation)
}

func ViewTransformationContext(s *godog.ScenarioContext) {
	s.Step(`^t ← view_transform\(from, to, up\)$`, tView_transformfromToUp)
	s.Step(`^t = scaling\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, tScaling)
	s.Step(`^t = translation\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, tTranslation)

	s.BeforeScenario(func(sc *godog.Scenario) {
		t = matrices["t"]
	})
}
