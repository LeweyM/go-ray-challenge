package main

import (
	"fmt"
	"math"
)

func expectPoint(t *Tuple, arg1 float64, arg2 float64, arg3 float64) error {
	if !t.isPoint() {
		return fmt.Errorf("%v should be a point", *t)
	}
	return expectTuple(t, arg1, arg2, arg3, 1.0)
}

func expectVector(t *Tuple, arg1 float64, arg2 float64, arg3 float64) error {
	if !t.isVector() {
		return fmt.Errorf("%v should be a vector", *t)
	}
	return expectTuple(t, arg1, arg2, arg3, 0.0)
}

func expectTuple(tuple *Tuple, arg1 float64, arg2 float64, arg3 float64, arg4 float64) error {
	if floatEquals(tuple.x, arg1) && floatEquals(tuple.y, arg2) && floatEquals(tuple.z, arg3) && floatEquals(tuple.w, arg4) {
		return nil
	}
	return fmt.Errorf("%v should have values (%g, %g, %g, %g)", *tuple, arg1, arg2, arg3, arg4)
}

func floatEquals(f1 float64, f2 float64) bool {
	EPSILON := 0.0001
	return math.Abs(f1 - f2) < EPSILON
}

func expectTrue(cond bool, errMsg string) error {
	if !cond {
		return fmt.Errorf(errMsg)
	}
	return nil
}

func expectFalse(cond bool, errMsg string) error {
	return expectTrue(!cond, errMsg)
}
