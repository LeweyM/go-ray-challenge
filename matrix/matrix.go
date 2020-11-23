package matrix

import (
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
)

type Matrix struct {
	size int
	cells []float64
}

func NewTranslation(x, y, z float64) Matrix {
	t := NewIdentityMatrix()
	t.set(0, 3, x)
	t.set(1, 3, y)
	t.set(2, 3, z)
	return t
}

func NewScale(x, y, z float64) Matrix {
	t := NewIdentityMatrix()
	t.set(0, 0, x)
	t.set(1, 1, y)
	t.set(2, 2, z)
	return t
}

func NewRotationX(rads float64) Matrix {
	t := NewIdentityMatrix()
	t.set(1, 1, math.Cos(rads))
	t.set(2, 2, math.Cos(rads))
	t.set(1, 2, math.Sin(rads) * -1.0)
	t.set(2, 1, math.Sin(rads))
	return t
}

func NewIdentityMatrix() Matrix {
	return newMatrix([]float64{1,0,0,0, 0,1,0,0, 0,0,1,0, 0,0,0,1})
}

func NewMatrix(m [][]float64) Matrix {
	var cells []float64
	for _, r := range m {
		for _, c := range r {
			cells = append(cells, c)
		}
	}
	return Matrix{size: len(m), cells: cells}
}

func newMatrix(m []float64) Matrix {
	var cells []float64
	for _, c := range m {
		cells = append(cells, c)
	}
	return Matrix{size: 4, cells: cells}
}

func (m *Matrix) Equals(other Matrix) bool {
	if m.size != other.size {
		return false
	}
	for i := 0; i < m.size; i++ {
		FloatEquals(m.cells[i], other.cells[i])
		if !FloatEquals(m.cells[i], other.cells[i]) {
			return false
		}
	}
	return true
}

func (m *Matrix) Get(x, y int) float64 {
	return m.cells[m.index(x, y)]
}

func (m *Matrix) Multiply(other Matrix) Matrix {
	var newCells []float64
	for row := 0; row < m.size; row++ {
		for col := 0; col < m.size; col++ {
			cell := m.Get(row, 0) * other.Get(0, col) +
			m.Get(row, 1) * other.Get(1, col) +
			m.Get(row, 2) * other.Get(2, col) +
			m.Get(row, 3) * other.Get(3, col)
			newCells = append(newCells, cell)
		}
	}
	return newMatrix(newCells)
}

func (m *Matrix) MultiplyTuple(t *tuple.Tuple) tuple.Tuple {
	var r []float64
	for row := 0; row < m.size; row++ {
		c := m.Get(row, 0) * t.X +
			m.Get(row, 1) * t.Y +
			m.Get(row, 2) * t.Z +
			m.Get(row, 3) * t.W
		r = append(r, c)
	}
	return tuple.Tuple{X: r[0], Y: r[1], Z: r[2], W: r[3]}
}

func (m *Matrix) Transpose() Matrix {
	var d []float64
	for col := 0; col < m.size; col++ {
		for row := 0; row < m.size; row++ {
			d = append(d, m.Get(row, col))
		}
	}
	return Matrix{cells: d, size: m.size}
}

func (m *Matrix) Determinant() float64 {
	if m.size == 2 {
		return m.Get(0, 0) * m.Get(1, 1) -
			m.Get(0, 1) * m.Get(1, 0)
	} else {
		firstRow := m.Rows()[0]
		counter := 0.0
		for i, cell := range firstRow {
			counter += cell * m.Cofactor(0, i)
		}
		return counter
	}
}

func (m *Matrix) SubMatrix(row int, col int) Matrix {
	rows := m.Rows()
	subRows := splice2d(row, rows)
	for i, subRow := range subRows {
		subRows[i] = splice(col, subRow)
	}
	return NewMatrix(subRows)
}

func (m *Matrix) Rows() [][]float64 {
	cellsCpy := make([]float64, m.size*m.size)
	copy(cellsCpy, m.cells)
	var rows [][]float64
	for i := 0; i < m.size; i++ {
		offset := i * m.size
		rows = append(rows, cellsCpy[offset : offset+m.size])
	}
	return rows
}

func (m *Matrix) Minor(row, col int) float64 {
	subMatrix := m.SubMatrix(row, col)
	return subMatrix.Determinant()
}

func (m *Matrix) Cofactor(row, col int) float64 {
	minor := m.Minor(row, col)
	if (row+col) % 2 == 1 {
		return minor * -1
	} else {
		return minor
	}
}

func (m *Matrix) IsInvertable() bool {
	return m.Determinant() != 0
}

func (m *Matrix) Invert() Matrix {
	cells := make([]float64, m.size*m.size)
	copy(cells, m.cells)
	for row := 0; row < m.size; row++ {
		for col := 0; col < m.size; col++ {
			c := m.Cofactor(row, col)
			cells[m.index(col, row)] = c / m.Determinant()
		}
	}
	return newMatrix(cells)
}

func FloatEquals(f1 float64, f2 float64) bool {
	EPSILON := 0.0001
	return math.Abs(f1 - f2) < EPSILON
}

func splice2d(i int, arr [][]float64) [][]float64 {
	return append(arr[0:i], arr[i+1:]...)
}

func splice(i int, arr []float64) []float64 {
	return append(arr[0:i], arr[i+1:]...)
}

func (m *Matrix) index(x, y int) int {
	return y + (x * m.size)
}

func (m *Matrix) set(row, col int, value float64) {
	m.cells[m.index(row, col)] = value
}
