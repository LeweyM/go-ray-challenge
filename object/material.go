package object

import (
	"github/lewismetcalf/goRayChallenge/tuple"
)

type Material struct {
	ambient float64
	diffuse float64
	shininess float64
	specular float64
}

func NewMaterial() *Material {
	return &Material{0.1, 0.9, 200, 0.9}
}

func (m *Material) Color() *tuple.Color {
	return tuple.NewColor(1, 1, 1)
}

func (m *Material) Ambient() float64 {
	return m.ambient
}

func (m *Material) Diffuse() float64 {
	return m.diffuse
}

func (m *Material) Shininess() float64 {
	return m.shininess
}

func (m *Material) Specular() float64 {
	return m.specular
}

func (m *Material) Equals(other Material) bool {
	return m.Color().Equals(other.Color()) && *m == other
}

func (m *Material) SetAmbient(ambient float64) {
	m.ambient = ambient
}
