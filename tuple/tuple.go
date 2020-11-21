package tuple

import "math"

type Tuple struct {
	X, Y, Z, W float64
}

func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) Add(t2 *Tuple) *Tuple {
	return &Tuple{
		X: t.X + t2.X,
		Y: t.Y + t2.Y,
		Z: t.Z + t2.Z,
		W: t.W + t2.W,
	}
}

func (t Tuple) Subtract(t2 *Tuple) *Tuple {
	return &Tuple{
		X: t.X - t2.X,
		Y: t.Y - t2.Y,
		Z: t.Z - t2.Z,
		W: t.W - t2.W,
	}
}

func (t Tuple) Negate() *Tuple {
	return &Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}
}

func (t Tuple) Multiply(scalar float64) *Tuple {
	return &Tuple{
		X: t.X * scalar,
		Y: t.Y * scalar,
		Z: t.Z * scalar,
		W: t.W * scalar,
	}
}

func (t Tuple) Divide(scalar float64) *Tuple {
	return &Tuple{
		X: t.X / scalar,
		Y: t.Y / scalar,
		Z: t.Z / scalar,
		W: t.W / scalar,
	}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt((t.X * t.X) + (t.Y * t.Y) + (t.Z * t.Z) + (t.W * t.W))
}

func (t Tuple) Normalize() *Tuple {
	mag := t.Magnitude()
	return &Tuple{
		X: t.X / mag,
		Y: t.Y / mag,
		Z: t.Z / mag,
		W: t.W / mag,
	}
}

func (t Tuple) Dot(t2 *Tuple) float64 {
	return t.X* t2.X +
		t.Y* t2.Y +
		t.Z* t2.Z +
		t.W* t2.W
}

func (t Tuple) Cross(b *Tuple) *Tuple {
	return NewVector(t.Y* b.Z- t.Z* b.Y, t.Z* b.X- t.X* b.Z, t.X* b.Y- t.Y* b.X)
}

func (t *Tuple) Equals(other *Tuple) bool {
	return FloatEquals(t.X, other.X) && FloatEquals(t.Y, other.Y) && FloatEquals(t.Z, other.Z) && FloatEquals(t.W, other.W)
}

func FloatEquals(f1 float64, f2 float64) bool {
	EPSILON := 0.0001
	return math.Abs(f1 - f2) < EPSILON
}
