package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
	"strconv"
)

func transformPPoint(t, p string, arg1, arg2, arg3 float64) error {
	transform := matrices[t]
	point := tuples[p]
	multiplyTuple := transform.MultiplyTuple(point)
	return ExpectPoint(&multiplyTuple, arg1, arg2, arg3)
}

func transformVVector(t, v string, arg1, arg2, arg3 float64) error {
	transform := matrices[t]
	vector := tuples[v]
	multiplyTuple := transform.MultiplyTuple(vector)
	return ExpectVector(&multiplyTuple, arg1, arg2, arg3)
}

func transformTranslation(t string, arg1, arg2, arg3 float64) error {
	matrices[t] = matrix.NewTranslation(arg1, arg2, arg3)
	return nil
}

func transformTimesVEqualsV(t, v, unneeded string) error {
	transform := matrices[t]
	vector := tuples[v]
	multiplyTuple := transform.MultiplyTuple(vector)
	return ExpectVectorEqual(multiplyTuple, vector)
}

func transformScaling(t string, arg1, arg2, arg3 float64) error {
	matrices[t] = matrix.NewScale(arg1, arg2, arg3)
	return nil
}

func rotationX(r string, rads float64) error {
	matrices[r] = matrix.NewRotationX(math.Pi / rads)
	return nil
}

func halfQuarterTimesPEqualsPoint(m, p string, p1 float64, p2root string, p2Denom float64, p3root string, p3Denom float64) error {
	rotation := matrices[m]
	point := tuples[p]
	p2 := parseRoot(p2root)
	p3 := parseRoot(p3root)
	multiply := rotation.MultiplyTuple(point)
	return ExpectPoint(&multiply, p1, p2 / p2Denom, p3 / p3Denom)
}

func parseRoot(s string) float64 {
	sign := 1.0
	rest := s
	if s[0] == '-' {
		sign = -1.0
		rest = s[1:]
	}
	float, err := strconv.ParseFloat(rest[3:], 64)
	if err == nil {
		return math.Sqrt(float) * sign
	}
	panic("bad root parsing")
}

func InitializeTransformationScenario(s *godog.ScenarioContext) {
	// transformation
	s.Step(`^`+VarName+` \* `+VarName+` = point\(`+Number+`, `+Number+`, `+Number+`\)$`, transformPPoint)
	s.Step(`^`+VarName+` \* `+VarName+` = vector\(`+Number+`, `+Number+`, `+Number+`\)$`, transformVVector)
	s.Step(`^`+VarName+` ← translation\(`+Number+`, `+Number+`, `+Number+`\)$`, transformTranslation)
	s.Step(`^`+VarName+` \* `+VarName+` = `+VarName+`$`, transformTimesVEqualsV)
	s.Step(`^`+VarName+` ← scaling\(`+Number+`, `+Number+`, `+Number+`\)$`, transformScaling)
	// rotation
	s.Step(`^`+VarName+` ← rotation_x\(π \/ `+Number+`\)$`, rotationX)
	s.Step(`^`+VarName+` \* `+VarName+` = point\(`+Number+`, `+rootNumber+`\/`+Number+`, `+rootNumber+`\/`+Number+`\)$`, halfQuarterTimesPEqualsPoint)

	s.BeforeScenario(func(sc *godog.Scenario) {
		matrices = make(map[string]matrix.Matrix)
	})
}

func ExpectVectorEqual(multiplyTuple tuple.Tuple, vector *tuple.Tuple) error {
	return ExpectTrue(multiplyTuple.Equals(vector),
		fmt.Sprintf("Expected %v and %p to be equal.", multiplyTuple, vector))
}

const rootNumber = `(\-?√\d+)`
