package main

import (
	"fmt"
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

func IsAPoint(t string) error {
	return expectTrue(tuples[t].isPoint(), fmt.Sprintf("%v should be a point", *tuples["a"]))
}

func IsNotAVector(t string) error {
	return expectFalse(tuples[t].isVector(), "should not be a vector")
}

func IsNotAPoint(t string) error {
	return expectFalse(tuples[t].isPoint(), fmt.Sprintf("%v should be a point", *tuples["a"]))
}

func IsAVector(t string) error {
	return expectTrue(tuples[t].isVector(), "should be a vector")
}

func addTuples(t1, t2 string, arg1, arg2, arg3, arg4 float64) error {
	tuple1 := tuples[t1]
	tuple2 := tuples[t2]
	tupleSum := tuple1.add(tuple2)
	return expectTuple(tupleSum, arg1, arg2, arg3, arg4)
}

func equalsTuple(t string, arg1, arg2, arg3, arg4 float64) error {
	tuple := tuples[t]
	return expectTuple(tuple, arg1, arg2, arg3, arg4)
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^([A-Za-z0-9]*) ← tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, setTuple)
	s.Step(`^([A-Za-z0-9]*) ← point\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, setPoint)
	s.Step(`^([A-Za-z0-9]*) ← vector\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, setVector)
	s.Step(`^([A-Za-z0-9]*) is a point$`, IsAPoint)
	s.Step(`^([A-Za-z0-9]*) is not a vector$`, IsNotAVector)
	s.Step(`^([A-Za-z0-9]*) is a vector$`, IsAVector)
	s.Step(`^([A-Za-z0-9]*) is not a point$`, IsNotAPoint)
	s.Step(`^([A-Za-z0-9]*) = tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, equalsTuple)
	s.Step(`^([A-Za-z0-9]*) \+ ([A-Za-z0-9]*) = tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, addTuples)
	s.BeforeScenario(func(sc *godog.Scenario) {
		tuples = make(map[string]*Tuple)
	})
}

func expectTuple(tuple *Tuple, arg1 float64, arg2 float64, arg3 float64, arg4 float64) error {
	if tuple.x == arg1 && tuple.y == arg2 && tuple.z == arg3 && tuple.w == arg4 {
		return nil
	}
	return fmt.Errorf("%v should have values (%g, %g, %g, %g)", *tuple, arg1, arg2, arg3, arg4)
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
