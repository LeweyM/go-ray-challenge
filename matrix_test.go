package main

import (
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

func theFollowingXMatrixM(x, y int, v string, m *messages.PickleStepArgument_PickleTable) error {
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
	matrices[v] = matrix.NewMatrix4(mat[0], mat[1], mat[2], mat[3])
	return nil
}


func InitializeMatrixScenario(s *godog.ScenarioContext) {
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = `+Float+`$`, m)
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = `+Number+`$`, m)
	s.Step(`^the following `+Number+`x`+Number+` matrix `+VarName+`:$`, theFollowingXMatrixM)

	s.BeforeScenario(func(sc *godog.Scenario) {
		matrices = make(map[string]matrix.Matrix)
	})
}
