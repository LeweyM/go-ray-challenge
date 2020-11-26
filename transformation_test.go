package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
	"strconv"
)

var sqrt2 = math.Sqrt(2)

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

func transformTranslation(a string, arg1, arg2, arg3 float64) error {
	t = matrix.NewTranslation(arg1, arg2, arg3)
	matrices[a] = matrix.NewTranslation(arg1, arg2, arg3)
	return nil
}

func transformTimesVEqualsV(a, v, unneeded string) error {
	vector := tuples[v]
	multiplyTuple := t.MultiplyTuple(vector)
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

func rotationY(r string, rads float64) error {
	matrices[r] = matrix.NewRotationY(math.Pi / rads)
	return nil
}

func halfQuarterTimesPEqualsPoint(m, p string, p1 float64, p2root string, p2Denom float64, p3root string, p3Denom float64) error {
	rotation := matrices[m]
	point := tuples[p]
	p2 := parseRoot(p2root)
	p3 := parseRoot(p3root)
	multiply := rotation.MultiplyTuple(point)
	return ExpectPoint(&multiply, p1, p2/p2Denom, p3/p3Denom)
}

func half_quarterPPointb(a, b string, d1 string, n1, d2 float64, d3 string, n3 float64) error {
	rotation := matrices[a]
	point := tuples[b]
	p1 := parseRoot(d1)
	p3 := parseRoot(d3)
	multiply := rotation.MultiplyTuple(point)
	return ExpectPoint(&multiply, p1/n1, d2, p3/n3)
}

func half_quarterPPointc(a, p string) error {
	point := tuples[p]
	halfQuarter := matrices[a]
	multiply := halfQuarter.MultiplyTuple(point)
	return ExpectPoint(&multiply, sqrt2/-2, sqrt2/2, 0)
}

func full_quarterRotation_z(a string, arg1 float64) error {
	matrices[a] = matrix.NewRotationZ(math.Pi / arg1)
	return nil
}

func half_quarterRotation_z(a string, arg1 float64) error {
	matrices[a] = matrix.NewRotationZ(math.Pi / arg1)
	return nil
}

func transformShearing(s string, arg1, arg2, arg3, arg4, arg5, arg6 float64) error {
	matrices[s] = matrix.NewShearing(arg1, arg2, arg3, arg4, arg5, arg6)
	return nil
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
	s.Step(`^`+VarName+` ← rotation_y\(π \/ `+Number+`\)$`, rotationY)
	s.Step(`^`+VarName+` \* `+VarName+` = point\(`+Number+`, `+rootNumber+`\/`+Number+`, `+rootNumber+`\/`+Number+`\)$`, halfQuarterTimesPEqualsPoint)
	s.Step(`^`+VarName+` \* `+VarName+` = point\(`+rootNumber+`\/`+Number+`, `+Number+`, `+rootNumber+`\/`+Number+`\)$`, half_quarterPPointb)
	s.Step(`^`+VarName+` ← rotation_z\(π \/ (\d+)\)$`, full_quarterRotation_z)
	s.Step(`^`+VarName+` \* `+VarName+` = point\(-√2\/2, √2\/2, 0\)$`, half_quarterPPointc)
	s.Step(`^`+VarName+` ← rotation_z\(π \/ (\d+)\)$`, half_quarterRotation_z)
	//shearing
	s.Step(`^`+VarName+` ← shearing\((\d+), (\d+), (\d+), (\d+), (\d+), (\d+)\)$`, transformShearing)

	s.BeforeScenario(func(sc *godog.Scenario) {
		matrices = make(map[string]matrix.Matrix)
	})
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

func ExpectVectorEqual(multiplyTuple tuple.Tuple, vector *tuple.Tuple) error {
	return ExpectTrue(multiplyTuple.Equals(vector),
		fmt.Sprintf("Expected %v and %p to be equal.", multiplyTuple, vector))
}

const rootNumber = `(\-?√\d+)`
