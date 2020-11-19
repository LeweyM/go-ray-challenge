package main

import "math"

type Tuple struct {
	x, y, z, w float64
}

func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

func (t Tuple) isVector() bool {
	return t.w == 0.0
}

func (t Tuple) isPoint() bool {
	return t.w == 1.0
}

func (t Tuple) add(t2 *Tuple) *Tuple {
	return &Tuple{
		x: t.x + t2.x,
		y: t.y + t2.y,
		z: t.z + t2.z,
		w: t.w + t2.w,
	}
}

func (t Tuple) subtract(t2 *Tuple) *Tuple {
	return &Tuple{
		x: t.x - t2.x,
		y: t.y - t2.y,
		z: t.z - t2.z,
		w: t.w - t2.w,
	}
}

func (t Tuple) negate() *Tuple {
	return &Tuple{
		x: -t.x,
		y: -t.y,
		z: -t.z,
		w: -t.w,
	}
}

func (t Tuple) multiply(scalar float64) *Tuple {
	return &Tuple{
		x: t.x * scalar,
		y: t.y * scalar,
		z: t.z * scalar,
		w: t.w * scalar,
	}
}

func (t Tuple) divide(scalar float64) *Tuple {
	return &Tuple{
		x: t.x / scalar,
		y: t.y / scalar,
		z: t.z / scalar,
		w: t.w / scalar,
	}
}

func (t Tuple) magnitude() float64 {
	return math.Sqrt((t.x * t.x) + (t.y * t.y) + (t.z * t.z) + (t.w * t.w))
}

func (t Tuple) normalize() *Tuple {
	mag := t.magnitude()
	return &Tuple{
		x: t.x / mag,
		y: t.y / mag,
		z: t.z / mag,
		w: t.w / mag,
	}
}

func (t Tuple) dot(t2 *Tuple) float64 {
	return t.x * t2.x +
		t.y * t2.y +
		t.z * t2.z +
		t.w * t2.w
}

func (t Tuple) cross(b *Tuple) *Tuple {
	return NewVector(t.y * b.z - t.z * b.y, t.z * b.x - t.x * b.z, t.x * b.y - t.y * b.x)
}
