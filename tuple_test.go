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

func InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^([A-Za-z0-9]*) ← tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, setTuple)
	s.Step(`^([A-Za-z0-9]*) ← point\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, setPoint)
	s.Step(`^([A-Za-z0-9]*) ← vector\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, setVector)
	s.Step(`^([A-Za-z0-9]*) = point\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, equalsPoint)
	s.Step(`^([A-Za-z0-9]*) = vector\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, equalsVector)
	s.Step(`^([A-Za-z0-9]*) = tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, equalsTuple)
	s.Step(`^-([A-Za-z0-9]*) = tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, negativeEqualsTuple)
	s.Step(`^([A-Za-z0-9]*) \+ ([A-Za-z0-9]*) = tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, addTuples)
	s.Step(`^([A-Za-z0-9]*) \- ([A-Za-z0-9]*) = point\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, subPoints)
	s.Step(`^([A-Za-z0-9]*) \- ([A-Za-z0-9]*) = vector\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, subVectors)
	s.BeforeScenario(func(sc *godog.Scenario) {
		tuples = make(map[string]*Tuple)
	})
}
