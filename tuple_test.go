package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"math"
)

var tuples map[string]*Tuple
var colors map[string]*Color

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

func setColor(c string, arg1, arg2, arg3 float64) error {
	colors[c] = NewColor(arg1, arg2, arg3)
	return nil
}

func setNormalizeVector(n, v string ) error {
	tuples[n] = tuples[v].normalize()
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

func magnitudeEqualsFloat(t string, scalar float64) error {
	tuple := tuples[t]
	return expectTrue(tuple.magnitude() == scalar, fmt.Sprintf("magnitude should be %v, but is %g", scalar, tuple.magnitude()))
}

func magnitudeEqualsSquareRoot(t string, number float64) error {
	tuple := tuples[t]
	return expectTrue(tuple.magnitude() == math.Sqrt(number), fmt.Sprintf("magnitude should be √%v, but is %g", number, tuple.magnitude()))
}

func normalizeEqualsVector(t string, arg1, arg2, arg3 float64) error {
	tuple := tuples[t]
	return expectVector(tuple.normalize(), arg1, arg2, arg3)
}

func dotEqualsFloat(a, b string, scalar float64) error {
	tupleA := tuples[a]
	tupleB := tuples[b]
	dot := tupleA.dot(tupleB)
	return expectFloatEquals(dot, scalar)
}

func crossEqualsVector(a, b string, arg1, arg2, arg3 float64) error {
	tupleA := tuples[a]
	tupleB := tuples[b]
	cross := tupleA.cross(tupleB)
	return expectVector(cross, arg1, arg2, arg3)
}

func colorRedEqualsFloat(a string, scalar float64) error {
	return expectFloatEquals(colors[a].red(), scalar)
}

func colorBlueEqualsFloat(a string, scalar float64) error {
	return expectFloatEquals(colors[a].blue(), scalar)
}

func colorGreenEqualsFloat(a string, scalar float64) error {
	return expectFloatEquals(colors[a].green(), scalar)
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^`+varName+` ← `+tuple+`$`, setTuple)
	s.Step(`^`+varName+` ← `+point+`$`, setPoint)
	s.Step(`^`+varName+` ← `+vector+`$`, setVector)
	s.Step(`^`+varName+` ← `+color+`$`, setColor)
	s.Step(`^`+varName+` = `+point+`$`, equalsPoint)
	s.Step(`^`+varName+` = `+vector+`$`, equalsVector)
	s.Step(`^`+varName+` = `+tuple+`$`, equalsTuple)
	s.Step(`^-`+varName+` = `+tuple+`$`, negativeEqualsTuple)
	s.Step(`^`+varName+` \+ `+varName+` = `+tuple+`$`, addTuples)
	s.Step(`^`+varName+` \- `+varName+` = `+point+`$`, subPoints)
	s.Step(`^`+varName+` \- `+varName+` = `+vector+`$`, subVectors)
	s.Step(`^`+varName+` \* `+float+` = `+tuple+`$`, multipleEqualsTuple)
	s.Step(`^`+varName+` \/ `+float+` = `+tuple+`$`, divideEqualsTuple)
	s.Step(`^magnitude\(`+varName+`\) = `+float+`$`, magnitudeEqualsFloat)
	s.Step(`^magnitude\(`+varName+`\) = √`+number+`$`, magnitudeEqualsSquareRoot)
	s.Step(`^`+varName+` ← normalize\(`+varName+`\)$`, setNormalizeVector)
	s.Step(`^normalize\(`+varName+`\) = `+vector+`$`, normalizeEqualsVector)
	s.Step(`^dot\(`+varName+`, `+varName+`\) = `+float+`$`, dotEqualsFloat)
	s.Step(`^cross\(`+varName+`, `+varName+`\) = `+vector+`$`, crossEqualsVector)
	s.Step(`^`+varName+`.red = `+float+`$`, colorRedEqualsFloat)
	s.Step(`^`+varName+`.blue = `+float+`$`, colorBlueEqualsFloat)
	s.Step(`^`+varName+`.green = `+float+`$`, colorGreenEqualsFloat)

	s.BeforeScenario(func(sc *godog.Scenario) {
		tuples = make(map[string]*Tuple)
		colors = make(map[string]*Color)
	})
}

const (
	varName = `([A-Za-z0-9]*)`
	float   = `(\-*\d+\.\d+)`
	number  = `(\d+)`
	color   = `color\(` + float + `, ` + float + `, ` + float + `\)`
	point   = `point\(` + float + `, ` + float + `, ` + float + `\)`
	vector  = `vector\(` + float + `, ` + float + `, ` + float + `\)`
	tuple   = `tuple\(` + float + `, ` + float + `, ` + float + `, ` + float + `\)`
)
