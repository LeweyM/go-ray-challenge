package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github/lewismetcalf/goRayChallenge/matrix"
	"strconv"
)

var matrices map[string]matrix.Matrix

func m(v string, x, y int, f float64) error {
	m := matrices[v]
	return ExpectFloatEquals(m.Get(x, y), f)
}

func equalMatrices(a, b string) error {
	m1 := matrices[a]
	m2 := matrices[b]
	return ExpectTrue(m1.Equals(m2), fmt.Sprintf("expected %v and %v to be equal", m1, m2))
}

func notEqualMatrices(a, b string) error {
	m1 := matrices[a]
	m2 := matrices[b]
	return ExpectFalse(m1.Equals(m2), fmt.Sprintf("expected %v and %v to be equal", m1, m2))
}

func theFollowingMatrixM(v string, m *messages.PickleStepArgument_PickleTable) error {
	var mat [][]float64
	for _, row := range m.GetRows() {
		var r []float64
		for _, cell := range row.GetCells() {
			float, err := strconv.ParseFloat(cell.GetValue(), 64)
			if err != nil {
				return err
			}
			r = append(r, float)
		}
		mat = append(mat, r)
	}
	matrices[v] = matrix.NewMatrix(mat)
	return nil
}

func theFollowingXMatrixM(x, y int, v string, m *messages.PickleStepArgument_PickleTable) error {
	return theFollowingMatrixM(v, m)
}

func InitializeMatrixScenario(s *godog.ScenarioContext) {
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = `+Float+`$`, m)
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = `+Number+`$`, m)
	s.Step(`^the following `+Number+`x`+Number+` matrix `+VarName+`:$`, theFollowingXMatrixM)
	s.Step(`^the following matrix `+VarName+`:$`, theFollowingMatrixM)
	s.Step(`^`+VarName+` = `+VarName+`$`, equalMatrices)
	s.Step(`^`+VarName+` != `+VarName+`$`, notEqualMatrices)

	s.BeforeScenario(func(sc *godog.Scenario) {
		matrices = make(map[string]matrix.Matrix)
	})
}
