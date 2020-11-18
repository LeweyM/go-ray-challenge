package main

import (
	"github.com/cucumber/godog"
)

var tuples map[string]*Tuple

func setTuple(tuple string, arg1, arg2, arg3, arg4 float64) error {
	tuples[tuple] = &Tuple{arg1, arg2, arg3, arg4}
	return nil
}

func setPoint(p string, arg1, arg2, arg3 float64) error {
	tuples[p] = NewPoint(arg1, arg2, arg3)
	return nil
}

func setVector(v string, arg1, arg2, arg3 float64) error {
	tuples[v] = NewVector(arg1, arg2, arg3)
	return nil
}

func addTuples(t1, t2 string, arg1, arg2, arg3, arg4 float64) error {
	tuple1 := tuples[t1]
	tuple2 := tuples[t2]
	tupleSum := tuple1.add(tuple2)
	return expectTuple(tupleSum, arg1, arg2, arg3, arg4)
}

func subPoints(t1, t2 string, arg1, arg2, arg3 float64) error {
	tuple1 := tuples[t1]
	tuple2 := tuples[t2]
	tupleSum := tuple1.subtract(tuple2)
	return expectPoint(tupleSum, arg1, arg2, arg3)
}

func subVectors(t1, t2 string, arg1, arg2, arg3 float64) error {
	tuple1 := tuples[t1]
	tuple2 := tuples[t2]
	tupleSum := tuple1.subtract(tuple2)
	return expectVector(tupleSum, arg1, arg2, arg3)
}

func equalsPoint(t1 string, arg1, arg2, arg3 float64) error {
	tuple := tuples[t1]
	return expectPoint(tuple, arg1, arg2, arg3)
}

func equalsVector(t1 string, arg1, arg2, arg3 float64) error {
	tuple := tuples[t1]
	return expectVector(tuple, arg1, arg2, arg3)
}

func equalsTuple(t string, arg1, arg2, arg3, arg4 float64) error {
	tuple := tuples[t]
	return expectTuple(tuple, arg1, arg2, arg3, arg4)
}

func negativeEqualsTuple(t string, arg1, arg2, arg3, arg4 float64) error {
	tuple := tuples[t]
	return expectTuple(tuple.negate(), arg1, arg2, arg3, arg4)
}

func multipleEqualsTuple(t string, scalar, arg1, arg2, arg3, arg4 float64) error {
	tuple := tuples[t]
	return expectTuple(tuple.multiply(scalar), arg1, arg2, arg3, arg4)
}

func divideEqualsTuple(t string, scalar, arg1, arg2, arg3, arg4 float64) error {
	tuple := tuples[t]
	return expectTuple(tuple.divide(scalar), arg1, arg2, arg3, arg4)
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^`+varName+` ← `+tuple+`$`, setTuple)
	s.Step(`^`+varName+` ← `+point+`$`, setPoint)
	s.Step(`^`+varName+` ← `+vector+`$`, setVector)
	s.Step(`^`+varName+` = `+point+`$`, equalsPoint)
	s.Step(`^`+varName+` = `+vector+`$`, equalsVector)
	s.Step(`^`+varName+` = `+tuple+`$`, equalsTuple)
	s.Step(`^-`+varName+` = `+tuple+`$`, negativeEqualsTuple)
	s.Step(`^`+varName+` \+ `+varName+` = `+tuple+`$`, addTuples)
	s.Step(`^`+varName+` \- `+varName+` = `+point+`$`, subPoints)
	s.Step(`^`+varName+` \- `+varName+` = `+vector+`$`, subVectors)
	s.Step(`^`+varName+` \* `+floatingPoint+` = `+tuple+`$`, multipleEqualsTuple)
	s.Step(`^`+varName+` \/ `+floatingPoint+` = `+tuple+`$`, divideEqualsTuple)
	s.BeforeScenario(func(sc *godog.Scenario) {
		tuples = make(map[string]*Tuple)
	})
}

const (
	varName       = `([A-Za-z0-9]*)`
	floatingPoint = `(\-*\d+\.\d+)`
	point         = `point\(` + floatingPoint + `, ` + floatingPoint + `, ` + floatingPoint + `\)`
	vector        = `vector\(` + floatingPoint + `, ` + floatingPoint + `, ` + floatingPoint + `\)`
	tuple         = `tuple\(` + floatingPoint + `, ` + floatingPoint + `, ` + floatingPoint + `, ` + floatingPoint + `\)`
)
