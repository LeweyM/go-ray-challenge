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
	return expectEqualMatrices(m1, m2)
}

func notEqualMatrices(a, b string) error {
	m1 := matrices[a]
	m2 := matrices[b]
	return ExpectFalse(m1.Equals(m2), fmt.Sprintf("expected %v and %v to be equal", m1, m2))
}

func theFollowingMatrixM(v string, m *messages.PickleStepArgument_PickleTable) error {
	matrices[v] = matrixFromPickleTable(m)
	return nil
}

func theFollowingXMatrixM(x, y int, v string, m *messages.PickleStepArgument_PickleTable) error {
	return theFollowingMatrixM(v, m)
}

func aBIsTheFollowingXMatrix(a, b string, x, y int, m *messages.PickleStepArgument_PickleTable) error {
	m1 := matrices[a]
	m2 := matrices[b]
	multiple := m1.Multiply(m2)
	expected := matrixFromPickleTable(m)

	return expectEqualMatrices(multiple, expected)
}

func aBTuple(a, b string, arg1, arg2, arg3, arg4 float64) error {
	m := matrices[a]
	t := tuples[b]
	res := m.MultiplyTuple(t)
	return ExpectTuple(&res, arg1, arg2, arg3, arg4)
}

func aEqualsIdentity_matrixA(a, unneededArg string) error {
	m := matrices[a]
	identityMatrix := matrix.NewIdentityMatrix()
	multiple := m.Multiply(identityMatrix)
	return ExpectTrue(multiple.Equals(m),
		fmt.Sprintf("expected %v and %v to be equal", multiple, m))
}

func identity_matrixAA(a, unneededArg string) error {
	t := tuples[a]
	identityMatrix := matrix.NewIdentityMatrix()
	multiple := identityMatrix.MultiplyTuple(t)
	return ExpectTrue(multiple.Equals(t),
		fmt.Sprintf("expected %v and %v to be equal", multiple, t))
}

func transposeAIsTheFollowingMatrix(a string, ma *messages.PickleStepArgument_PickleTable) error {
	m := matrices[a]
	expected := matrixFromPickleTable(ma)
	return expectEqualMatrices(m.Transpose(), expected)
}

func aIdentity_matrix(a string) error {
	m := matrices[a]
	return expectEqualMatrices(m, matrix.NewIdentityMatrix())
}

func aTransposeidentity_matrix(a string) error {
	identityMatrix := matrix.NewIdentityMatrix()
	matrices[a] = identityMatrix.Transpose()
	return nil
}

func determinantA(a string, arg1 float64) error {
	m := matrices[a]
	return ExpectFloatEquals(m.Determinant(), arg1)
}

func submatrixAIsTheFollowingXMatrix(a string, row, col, arg3, arg4 int, mm *messages.PickleStepArgument_PickleTable) error {
	m := matrices[a]
	return expectEqualMatrices(m.SubMatrix(row, col), matrixFromPickleTable(mm))
}

func bSubmatrixA(b, a string, arg1, arg2 int) error {
	m := matrices[a]
	matrices[b] = m.SubMatrix(arg1, arg2)
	return nil
}

func determinantB(a string, arg1 float64) error {
	m := matrices[a]
	return ExpectFloatEquals(m.Determinant(), arg1)
}

func minorA(a string, arg1, arg2 int, arg3 float64) error {
	m := matrices[a]
	return ExpectFloatEquals(m.Minor(arg1, arg2), arg3)
}

func cofactorA(a string, arg1, arg2 int, arg3 float64) error {
	m := matrices[a]
	return ExpectFloatEquals(m.Cofactor(arg1, arg2), arg3)
}

func InitializeMatrixScenario(s *godog.ScenarioContext) {
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = `+Float+`$`, m)
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = `+Number+`$`, m)
	s.Step(`^the following `+Number+`x`+Number+` matrix `+VarName+`:$`, theFollowingXMatrixM)
	s.Step(`^the following matrix `+VarName+`:$`, theFollowingMatrixM)
	s.Step(`^`+VarName+` = `+VarName+`$`, equalMatrices)
	s.Step(`^`+VarName+` != `+VarName+`$`, notEqualMatrices)
	s.Step(`^`+VarName+` \* `+VarName+` is the following (\d+)x(\d+) matrix:$`, aBIsTheFollowingXMatrix)
	s.Step(`^`+VarName+` \* `+VarName+` = tuple\((\d+), (\d+), (\d+), (\d+)\)$`, aBTuple)
	s.Step(`^`+VarName+` \* identity_matrix = `+VarName+``, aEqualsIdentity_matrixA)
	s.Step(`^identity_matrix \* `+VarName+` = `+VarName+`$`, identity_matrixAA)
	s.Step(`^transpose\(`+VarName+`\) is the following matrix:$`, transposeAIsTheFollowingMatrix)
	s.Step(`^`+VarName+` = identity_matrix$`, aIdentity_matrix)
	s.Step(`^`+VarName+` ← transpose\(identity_matrix\)$`, aTransposeidentity_matrix)
	s.Step(`^determinant\(`+VarName+`\) = `+Number+`\.$`, determinantA)
	s.Step(`^submatrix\(`+VarName+`, `+Number+`, `+Number+`\) is the following `+Number+`x`+Number+` matrix:$`, submatrixAIsTheFollowingXMatrix)
	s.Step(`^`+VarName+` ← submatrix\(`+VarName+`, `+Number+`, `+Number+`\)$`, bSubmatrixA)
	s.Step(`^determinant\(`+VarName+`\) = `+Number+`$`, determinantB)
	s.Step(`^minor\(`+VarName+`, `+Number+`, `+Number+`\) = `+Number+`$`, minorA)
	s.Step(`^cofactor\(`+VarName+`, `+Number+`, `+Number+`\) = `+Number+`$`, cofactorA)

	s.BeforeScenario(func(sc *godog.Scenario) {
		matrices = make(map[string]matrix.Matrix)
	})
}

func expectEqualMatrices(m1 matrix.Matrix, m2 matrix.Matrix) error {
	return ExpectTrue(m1.Equals(m2),
		fmt.Sprintf("Expected %v and %v to be equal", m1, m2))
}

func matrixFromPickleTable(m *messages.PickleStepArgument_PickleTable) matrix.Matrix {
	var mat [][]float64
	for _, row := range m.GetRows() {
		var r []float64
		for _, cell := range row.GetCells() {
			float, err := strconv.ParseFloat(cell.GetValue(), 64)
			if err != nil {
				return matrix.Matrix{}
			}
			r = append(r, float)
		}
		mat = append(mat, r)
	}
	return matrix.NewMatrix(mat)
}
