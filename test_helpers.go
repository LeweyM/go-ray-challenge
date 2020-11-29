package main

import (
	"fmt"
	"math"
)

const (
	VarName    = `([A-Za-z0-9_]*)`
	complexNum = `(\-?\d*\.?\d*)`
	Float      = `(\-*\d*\.\d*)`
	Number     = `(\-*\d+)`
	Color      = `color\(` + complexNum + `, ` + complexNum + `, ` + complexNum + `\)`
	Point      = `point\(` + Float + `, ` + Float + `, ` + Float + `\)`
	Vector     = `vector\(` + Float + `, ` + Float + `, ` + Float + `\)`
	TupleRex   = `tuple\(` + Float + `, ` + Float + `, ` + Float + `, ` + Float + `\)`
)

func ExpectFloatEquals(a float64, b float64) error {
	return ExpectTrue(FloatEquals(a, b), fmt.Sprintf("Expected %g, got %g", b, a))
}

func FloatEquals(f1 float64, f2 float64) bool {
	EPSILON := 0.0001
	return math.Abs(f1 - f2) < EPSILON
}

func ExpectTrue(cond bool, errMsg string) error {
	if !cond {
		return fmt.Errorf(errMsg)
	}
	return nil
}

func ExpectFalse(cond bool, errMsg string) error {
	return ExpectTrue(!cond, errMsg)
}
