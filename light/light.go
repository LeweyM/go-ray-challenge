package light

import "github/lewismetcalf/goRayChallenge/tuple"

type PointLight struct {
	position  tuple.Tuple
	intensity tuple.Color
}

func (l *PointLight) Equals(other *PointLight) bool {
	return l.position.Equals(other.position) &&
		l.intensity.Equals(other.intensity)
}

func (l *PointLight) Intensity() *tuple.Color {
	return &l.intensity
}

func (l *PointLight) Position() *tuple.Tuple {
	return &l.position
}

func NewPointLight(position tuple.Tuple, intensity tuple.Color) *PointLight {
	return &PointLight{position, intensity}
}
