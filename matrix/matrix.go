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
