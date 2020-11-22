package matrix

import (
	"github/lewismetcalf/goRayChallenge/tuple"
)

type Matrix struct {
	size int
	cells []float64
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
		if m.cells[i] != other.cells[i] {
			return false
		}
	}
	return true
}

func (m *Matrix) Get(x, y int) float64 {
	return m.cells[m.index(x, y)]
}

func (m *Matrix) index(x, y int) int {
	return y + (x * m.size)
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
	return m.Get(0, 0) * m.Get(1, 1) -
		m.Get(0, 1) * m.Get(1, 0)
}

func (m *Matrix) SubMatrix(row int, col int) Matrix {
	rows := m.Rows()
	subRows := splice2d(row, rows)
	for i, subRow := range subRows {
		subRows[i] = splice(col, subRow)
	}
	return NewMatrix(subRows)
}

func splice2d(i int, arr [][]float64) [][]float64 {
	return append(arr[0:i], arr[i+1:]...)
}

func splice(i int, arr []float64) []float64 {
	return append(arr[0:i], arr[i+1:]...)
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
	if row+col % 2 == 1 {
		return minor * -1
	} else {
		return minor
	}
}
