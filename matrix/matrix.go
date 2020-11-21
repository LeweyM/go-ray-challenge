package matrix

type Matrix struct {
	size int
	cells []float64
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
