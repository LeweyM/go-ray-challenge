package matrix

type Matrix struct {
	size int
	cells []float64
}

func NewMatrix4(r1, r2, r3, r4 []float64) Matrix {
	var c []float64
	c = append(r1, r2...)
	c = append(c, r3...)
	c = append(c, r4...)
	return Matrix{size: 4, cells: c}
}

func (m *Matrix) Get(x, y int) float64 {
	return m.cells[m.index(x, y)]
}

func (m *Matrix) index(x, y int) int {
	return y + (x * m.size)
}
