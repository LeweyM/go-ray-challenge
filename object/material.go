package object

import (
	"github/lewismetcalf/goRayChallenge/light"
	"github/lewismetcalf/goRayChallenge/tuple"
	"math"
)

type Material struct {
	ambient   float64
	diffuse   float64
	shininess float64
	specular  float64
	color     tuple.Color
}

func NewMaterial() *Material {
	return &Material{0.1, 0.9, 200, 0.9, *tuple.NewColor(1, 1, 1)}
}

func (m *Material) Color() *tuple.Color {
	return &m.color
}

func (m *Material) SetColor(color *tuple.Color) {
	m.color = *color
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
	return m.Color().Equals(other.color) &&
		m.diffuse == other.diffuse &&
		m.ambient == other.ambient &&
		m.specular == other.specular &&
		m.shininess == other.shininess
}

func (m *Material) SetAmbient(ambient float64) {
	m.ambient = ambient
}

func (m *Material) Lighting(l *light.PointLight, point, eye, normal *tuple.Tuple) tuple.Color {
	effectiveColor := m.Color().Multiply(l.Intensity())

	lightV := l.Position().Subtract(point).Normalize()

	ambient := effectiveColor.MultiplyScalar(m.ambient)
	var diffuse *tuple.Color
	var specular *tuple.Color

	lightDotNormal := lightV.Dot(normal)
	if lightDotNormal < 0 {
		diffuse = tuple.NewColor(0, 0, 0)
		specular = tuple.NewColor(0, 0, 0)
	} else {
		diffuse = effectiveColor.MultiplyScalar(m.diffuse).MultiplyScalar(lightDotNormal)
		reflectV := lightV.Negate().Reflect(normal)
		reflectDotEye := reflectV.Dot(eye)
		if reflectDotEye <= 0 {
			specular = tuple.NewColor(0, 0, 0)
		} else {
			factor := math.Pow(reflectDotEye, m.shininess)
			specular = l.Intensity().MultiplyScalar(m.specular).MultiplyScalar(factor)
		}
	}
	return *ambient.Add(diffuse).Add(specular)
}

func (m *Material) SetDiffuse(d float64) {
	m.diffuse = d
}

func (m *Material) SetSpecular(s float64) {
	m.specular = s
}
