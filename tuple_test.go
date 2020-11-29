package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
)

var tuples map[string]*tuple.Tuple
var colors map[string]*tuple.Color

func ExpectColor(c *tuple.Color, arg1 float64, arg2 float64, arg3 float64) error {
	return ExpectTuple(&c.T, arg1, arg2, arg3, 0.0)
}

func ExpectVector(t *tuple.Tuple, arg1 float64, arg2 float64, arg3 float64) error {
	if !t.IsVector() {
		return fmt.Errorf("%v should be a vector", *t)
	}
	return ExpectTuple(t, arg1, arg2, arg3, 0.0)
}

func ExpectPoint(t *tuple.Tuple, arg1 float64, arg2 float64, arg3 float64) error {
	if !t.IsPoint() {
		return fmt.Errorf("%v should be a point", *t)
	}
	return ExpectTuple(t, arg1, arg2, arg3, 1.0)
}

func ExpectTuple(tuple *tuple.Tuple, arg1 float64, arg2 float64, arg3 float64, arg4 float64) error {
	if FloatEquals(tuple.X, arg1) && FloatEquals(tuple.Y, arg2) && FloatEquals(tuple.Z, arg3) && FloatEquals(tuple.W, arg4) {
		return nil
	}
	return fmt.Errorf("%v should have values (%g, %g, %g, %g)", *tuple, arg1, arg2, arg3, arg4)
}

func setTuple(t string, arg1, arg2, arg3, arg4 float64) error {
	tuples[t] = &tuple.Tuple{X: arg1, Y: arg2, Z: arg3, W: arg4}
	return nil
}

func setPoint(p string, arg1, arg2, arg3 float64) error {
	tuples[p] = tuple.NewPoint(arg1, arg2, arg3)
	return nil
}

func setVector(v string, arg1, arg2, arg3 float64) error {
	tuples[v] = tuple.NewVector(arg1, arg2, arg3)
	return nil
}

func setColor(c string, arg1, arg2, arg3 float64) error {
	colors[c] = tuple.NewColor(arg1, arg2, arg3)
	return nil
}

func equalsColor(c string, arg1, arg2, arg3 float64) error {
	color := colors[c]
	return ExpectColor(color, arg1, arg2, arg3)
}

func setNormalizeVector(n, v string ) error {
	normalize := tuples[v].Normalize()
	tuples[n] = &normalize
	return nil
}

func addTuples(t1, t2 string, arg1, arg2, arg3, arg4 float64) error {
	tuple1 := tuples[t1]
	tuple2 := tuples[t2]
	tupleSum := tuple1.Add(tuple2)
	return ExpectTuple(tupleSum, arg1, arg2, arg3, arg4)
}

func subPoints(t1, t2 string, arg1, arg2, arg3 float64) error {
	tuple1 := tuples[t1]
	tuple2 := tuples[t2]
	tupleSum := tuple1.Subtract(tuple2)
	return ExpectPoint(tupleSum, arg1, arg2, arg3)
}

func subVectors(t1, t2 string, arg1, arg2, arg3 float64) error {
	tuple1 := tuples[t1]
	tuple2 := tuples[t2]
	tupleSum := tuple1.Subtract(tuple2)
	return ExpectVector(tupleSum, arg1, arg2, arg3)
}

func equalsPoint(t1 string, arg1, arg2, arg3 float64) error {
	tuple := tuples[t1]
	return ExpectPoint(tuple, arg1, arg2, arg3)
}

func equalsVector(t1 string, arg1, arg2, arg3 float64) error {
	tuple := tuples[t1]
	return ExpectVector(tuple, arg1, arg2, arg3)
}

func equalsTuple(t string, arg1, arg2, arg3, arg4 float64) error {
	tuple := tuples[t]
	return ExpectTuple(tuple, arg1, arg2, arg3, arg4)
}

func negativeEqualsTuple(t string, arg1, arg2, arg3, arg4 float64) error {
	tuple := tuples[t]
	return ExpectTuple(tuple.Negate(), arg1, arg2, arg3, arg4)
}

func multipleEqualsTuple(t string, scalar, arg1, arg2, arg3, arg4 float64) error {
	tuple := tuples[t]
	return ExpectTuple(tuple.Multiply(scalar), arg1, arg2, arg3, arg4)
}

func divideEqualsTuple(t string, scalar, arg1, arg2, arg3, arg4 float64) error {
	tuple := tuples[t]
	return ExpectTuple(tuple.Divide(scalar), arg1, arg2, arg3, arg4)
}

func magnitudeEqualsFloat(t string, scalar float64) error {
	tuple := tuples[t]
	return ExpectTrue(tuple.Magnitude() == scalar, fmt.Sprintf("magnitude should be %v, but is %g", scalar, tuple.Magnitude()))
}

func magnitudeEqualsSquareRoot(t string, number float64) error {
	tuple := tuples[t]
	return ExpectTrue(tuple.Magnitude() == math.Sqrt(number), fmt.Sprintf("magnitude should be √%v, but is %g", number, tuple.Magnitude()))
}

func normalizeEqualsVector(t string, arg1, arg2, arg3 float64) error {
	tuple := tuples[t]
	normalize := tuple.Normalize()
	return ExpectVector(&normalize, arg1, arg2, arg3)
}

func dotEqualsFloat(a, b string, scalar float64) error {
	tupleA := tuples[a]
	tupleB := tuples[b]
	dot := tupleA.Dot(tupleB)
	return ExpectFloatEquals(dot, scalar)
}

func crossEqualsVector(a, b string, arg1, arg2, arg3 float64) error {
	tupleA := tuples[a]
	tupleB := tuples[b]
	cross := tupleA.Cross(*tupleB)
	return ExpectVector(cross, arg1, arg2, arg3)
}

func colorRedEqualsFloat(a string, scalar float64) error {
	return ExpectFloatEquals(colors[a].Red(), scalar)
}

func colorBlueEqualsFloat(a string, scalar float64) error {
	return ExpectFloatEquals(colors[a].Blue(), scalar)
}

func colorGreenEqualsFloat(a string, scalar float64) error {
	return ExpectFloatEquals(colors[a].Green(), scalar)
}

func addColors(c1, c2 string, arg1, arg2, arg3 float64) error {
	color1 := colors[c1]
	color2 := colors[c2]
	colorSum := color1.Add(color2)
	return ExpectColor(colorSum, arg1, arg2, arg3)
}

func subColors(c1, c2 string, arg1, arg2, arg3 float64) error {
	color1 := colors[c1]
	color2 := colors[c2]
	colorSum := color1.Subtract(color2)
	return ExpectColor(colorSum, arg1, arg2, arg3)
}

func multiplyColorByScalar(c string, scalar, arg1, arg2, arg3 float64) error {
	color := colors[c]
	colorScaled := color.MultiplyScalar(scalar)
	return ExpectColor(colorScaled, arg1, arg2, arg3)
}

func multiplyColors(c1, c2 string, arg1, arg2, arg3 float64) error {
	color1 := colors[c1]
	color2 := colors[c2]
	colorScaled := color1.Multiply(color2)
	return ExpectColor(colorScaled, arg1, arg2, arg3)
}

func nRootRootVector(arg1, arg2, arg3, arg4, arg5 float64) error {
	n = *tuple.NewVector(math.Sqrt(arg1)/arg2, math.Sqrt(arg3)/arg4, arg5)
	tuples["n"] = tuple.NewVector(math.Sqrt(arg1)/arg2, math.Sqrt(arg3)/arg4, arg5)
	return nil
}

var reflection tuple.Tuple

func rReflectvN() error {
	v := tuples["v"]
	n := tuples["n"]

	reflection = v.Reflect(n)
	return nil
}

func rVector(arg1, arg2, arg3 float64) error {
	return ExpectVector(&reflection, arg1, arg2, arg3)
}

func InitializeTupleScenario(s *godog.ScenarioContext) {
	s.Step(`^`+VarName+` ← `+TupleRex+`$`, setTuple)
	s.Step(`^`+VarName+` ← tuple\(` + Number + `, ` + Number + `, ` + Number + `, ` + Number + `\)$`, setTuple)
	s.Step(`^`+VarName+` ← `+Point+`$`, setPoint)
	s.Step(`^`+VarName+` ← point\(` + Number + `, ` + Number + `, ` + Number + `\)$`, setPoint)
	s.Step(`^`+VarName+` ← `+Vector+`$`, setVector)
	s.Step(`^`+VarName+` ← vector\(` + Number + `, ` + Number + `, ` + Number + `\)$`, setVector)
	s.Step(`^`+VarName+` ← `+Color+`$`, setColor)
	s.Step(`^`+VarName+` = `+Color+`$`, equalsColor)
	s.Step(`^`+VarName+` = `+Point+`$`, equalsPoint)
	s.Step(`^`+VarName+` = `+Vector+`$`, equalsVector)
	s.Step(`^`+VarName+` = `+TupleRex+`$`, equalsTuple)
	s.Step(`^-`+VarName+` = `+TupleRex+`$`, negativeEqualsTuple)
	s.Step(`^`+VarName+` \+ `+VarName+` = `+TupleRex+`$`, addTuples)
	s.Step(`^`+VarName+` \- `+VarName+` = `+Point+`$`, subPoints)
	s.Step(`^`+VarName+` \- `+VarName+` = `+Vector+`$`, subVectors)
	s.Step(`^`+VarName+` \* `+Float+` = `+TupleRex+`$`, multipleEqualsTuple)
	s.Step(`^`+VarName+` \/ `+Float+` = `+TupleRex+`$`, divideEqualsTuple)
	s.Step(`^magnitude\(`+VarName+`\) = `+Float+`$`, magnitudeEqualsFloat)
	s.Step(`^magnitude\(`+VarName+`\) = √`+Number+`$`, magnitudeEqualsSquareRoot)
	s.Step(`^`+VarName+` ← normalize\(`+VarName+`\)$`, setNormalizeVector)
	s.Step(`^normalize\(`+VarName+`\) = `+Vector+`$`, normalizeEqualsVector)
	s.Step(`^dot\(`+VarName+`, `+VarName+`\) = `+Float+`$`, dotEqualsFloat)
	s.Step(`^cross\(`+VarName+`, `+VarName+`\) = `+Vector+`$`, crossEqualsVector)
	s.Step(`^`+VarName+`.red = `+Float+`$`, colorRedEqualsFloat)
	s.Step(`^`+VarName+`.blue = `+Float+`$`, colorBlueEqualsFloat)
	s.Step(`^`+VarName+`.green = `+Float+`$`, colorGreenEqualsFloat)
	s.Step(`^`+VarName+` \+ `+VarName+` = `+Color+`$`, addColors)
	s.Step(`^`+VarName+` \- `+VarName+` = `+Color+`$`, subColors)
	s.Step(`^`+VarName+` \* `+VarName+` = `+Color+`$`, multiplyColors)
	s.Step(`^`+VarName+` \* `+Float+` = `+Color+`$`, multiplyColorByScalar)
	// reflection
	s.Step(`^n ← vector\(√(\d+)\/(\d+), √(\d+)\/(\d+), (\d+)\)$`, nRootRootVector)
	s.Step(`^r ← reflect\(v, n\)$`, rReflectvN)
	s.Step(`^r = vector\((\d+), (\d+), (\d+)\)$`, rVector)

	s.BeforeScenario(func(sc *godog.Scenario) {
		tuples = make(map[string]*tuple.Tuple)
		colors = make(map[string]*tuple.Color)
	})
}
