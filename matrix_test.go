package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github/lewismetcalf/goRayChallenge/matrix"
	"strconv"
)

var matrices map[string]matrix.Matrix

func setM(v string, x, y int, f float64) error {
	m := matrices[v]
	return ExpectFloatEquals(m.Get(x, y), f)
}

func mFraction(v string, x, y int, d, n float64) error {
	m := matrices[v]
	return ExpectFloatEquals(m.Get(x, y), d/n)
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
	return ExpectTrue(multiple.Equals(*t),
		fmt.Sprintf("expected %v and %v to be equal", multiple, t))
}

func transposeAIsTheFollowingMatrix(a string, ma *messages.PickleStepArgument_PickleTable) error {
	m := matrices[a]
	expected := matrixFromPickleTable(ma)
	return expectEqualMatrices(m.Transpose(), expected)
}

func aTransposeidentity_matrix(a, i string) error {
	identityMatrix := matrix.NewIdentityMatrix()
	matrices[i] = identityMatrix
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

func aIsInvertible(a string) error {
	m := matrices[a]
	return ExpectTrue(m.IsInvertable(), fmt.Sprintf("Expected %v to be invertable", m))
}

func aIsNotInvertible(a string) error {
	m := matrices[a]
	return ExpectFalse(m.IsInvertable(), fmt.Sprintf("Expected %v not be invertable", m))
}

func bInverseA(b, a string) error {
	m := matrices[a]
	matrices[b] = m.Invert()
	return nil
}

func bIsTheFollowingXMatrix(mm *messages.PickleStepArgument_PickleTable) error {
	return expectEqualMatrices(matrices["B"], matrixFromPickleTable(mm))
}

func tIsTheFollowingXMatrix(mm *messages.PickleStepArgument_PickleTable) error {
	return expectEqualMatrices(t, matrixFromPickleTable(mm))
}

func inverseAIsTheFollowingXMatrix(a string, arg1, arg2 int, mm *messages.PickleStepArgument_PickleTable) error {
	inverse := matrices[a]
	return expectEqualMatrices(inverse.Invert(), matrixFromPickleTable(mm))
}

func cATimesB(c, a, b string) error {
	ma := matrices[a]
	mb := matrices[b]
	matrices[c] = ma.Multiply(mb)
	return nil
}

func cInverseBA(c, b, a string) error {
	ma := matrices[a]
	mb := matrices[b]
	mc := matrices[c]
	return expectEqualMatrices(mc.Multiply(mb.Invert()), ma)
}

func InitializeMatrixScenario(s *godog.ScenarioContext) {
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = `+Float+`$`, setM)
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = `+Number+`$`, setM)
	s.Step(`^the following `+Number+`x`+Number+` matrix `+VarName+`:$`, theFollowingXMatrixM)
	s.Step(`^the following matrix `+VarName+`:$`, theFollowingMatrixM)
	s.Step(`^`+VarName+` = `+VarName+`$`, equalMatrices)
	s.Step(`^`+VarName+` != `+VarName+`$`, notEqualMatrices)
	s.Step(`^`+VarName+` \* `+VarName+` is the following (\d+)x(\d+) matrix:$`, aBIsTheFollowingXMatrix)
	s.Step(`^`+VarName+` \* `+VarName+` = tuple\((\d+), (\d+), (\d+), (\d+)\)$`, aBTuple)
	s.Step(`^`+VarName+` \* identity_matrix = `+VarName+``, aEqualsIdentity_matrixA)
	s.Step(`^identity_matrix \* `+VarName+` = `+VarName+`$`, identity_matrixAA)
	s.Step(`^transpose\(`+VarName+`\) is the following matrix:$`, transposeAIsTheFollowingMatrix)
	s.Step(`^`+VarName+` ← transpose\(`+VarName+`\)$`, aTransposeidentity_matrix)
	s.Step(`^determinant\(`+VarName+`\) = `+Number+`\.$`, determinantA)
	s.Step(`^submatrix\(`+VarName+`, `+Number+`, `+Number+`\) is the following `+Number+`x`+Number+` matrix:$`, submatrixAIsTheFollowingXMatrix)
	s.Step(`^`+VarName+` ← submatrix\(`+VarName+`, `+Number+`, `+Number+`\)$`, bSubmatrixA)
	s.Step(`^determinant\(`+VarName+`\) = `+Number+`$`, determinantB)
	s.Step(`^minor\(`+VarName+`, `+Number+`, `+Number+`\) = `+Number+`$`, minorA)
	s.Step(`^cofactor\(`+VarName+`, `+Number+`, `+Number+`\) = `+Number+`$`, cofactorA)
	s.Step(`^`+VarName+` is invertible$`, aIsInvertible)
	s.Step(`^`+VarName+` is not invertible$`, aIsNotInvertible)
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = `+Float+`$`, setM)
	s.Step(`^`+VarName+`\[`+Number+`,`+Number+`\] = (\-*\d+)\/(\d+)$`, mFraction)
	s.Step(`^`+VarName+` ← inverse\(`+VarName+`\)$`, bInverseA)
	s.Step(`^B is the following 4x4 matrix:$`, bIsTheFollowingXMatrix)
	s.Step(`^t is the following 4x4 matrix:$`, tIsTheFollowingXMatrix)
	s.Step(`^inverse\(`+VarName+`\) is the following (\d+)x(\d+) matrix:$`, inverseAIsTheFollowingXMatrix)
	s.Step(`^`+VarName+` ← `+VarName+` \* `+VarName+`$`, cATimesB)
	s.Step(`^`+VarName+` \* inverse\(`+VarName+`\) = `+VarName+`$`, cInverseBA)

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
