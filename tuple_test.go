package main

import (
	"fmt"
	"github.com/cucumber/godog"
)

var a *Tuple

func aTuple(arg1, arg2, arg3, arg4 float64) error {
	a = &Tuple{arg1, arg2, arg3, arg4}
	return nil
}

func ax(arg1 float64) error {
	if a.x != arg1 {
		return fmt.Errorf("x should be %f", arg1)
	}
	return nil
}

func ay(arg1 float64) error {
	if a.y != arg1 {
		return fmt.Errorf("y should be %g", arg1)
	}
	return nil
}

func az(arg1 float64) error {
	if a.z != arg1 {
		return fmt.Errorf("z should be %g", arg1)
	}
	return nil
}

func aw(arg1 float64) error {
	if a.w != arg1 {
		return fmt.Errorf("w should be %g", arg1)
	}
	return nil
}

func aIsAPoint() error {
	if !a.isPoint() {
		return fmt.Errorf("%v should be a point", *a)
	}
	return nil
}

func aIsNotAVector() error {
	if a.isVector() {
		return fmt.Errorf("should not be a vector")
	}
	return nil
}

func aIsNotAPoint() error {
	if a.isPoint() {
		return fmt.Errorf("%v should not be a point", *a)
	}
	return nil
}

func aIsAVector() error {
	if !a.isVector() {
		return fmt.Errorf("should be a vector")
	}
	return nil
}

func aPoint(arg1, arg2, arg3 float64) error {
	a = NewPoint(arg1, arg2, arg3)
	return nil
}

func aVector(arg1, arg2, arg3 float64) error {
	a = NewVector(arg1, arg2, arg3)
	return nil
}

func aEqualsTuple(arg1, arg2, arg3, arg4 float64) error {
	if a.x == arg1 && a.y == arg2 && a.z == arg3 && a.w == arg4 {
		return nil
	}
	return fmt.Errorf("%v should have values (%g, %g, %g, %g)", *a, arg1, arg2, arg3, arg4)
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^a ← tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, aTuple)
	s.Step(`^a ← point\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, aPoint)
	s.Step(`^a ← vector\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, aVector)
	s.Step(`^a\.x = (\-*\d+\.\d+)$`, ax)
	s.Step(`^a\.y = (\-*\d+\.\d+)$`, ay)
	s.Step(`^a\.z = (\-*\d+\.\d+)$`, az)
	s.Step(`^a\.w = (\-*\d+\.\d+)$`, aw)
	s.Step(`^a is a point$`, aIsAPoint)
	s.Step(`^a is not a vector$`, aIsNotAVector)
	s.Step(`^a is a vector$`, aIsAVector)
	s.Step(`^a is not a point$`, aIsNotAPoint)
	s.Step(`^a = tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, aEqualsTuple)
	s.BeforeScenario(func(sc *godog.Scenario) {
		a = nil
	})
}
