package world

import (
	"github/lewismetcalf/goRayChallenge/matrix"
	"math"
)

type Camera struct {
	HSize       int
	Transform   matrix.Matrix
	VSize       int
	FieldOfView float64
	halfWidth   float64
	halfHeight  float64
	pixelSize   float64
}

func NewCamera(hsize, vsize int, fieldOfView float64) *Camera {
	camera := Camera{
		HSize:       hsize,
		Transform:   matrix.NewIdentityMatrix(),
		VSize:       vsize,
		FieldOfView: fieldOfView,
	}
	camera.calculate()
	return &camera

}

func (c *Camera) PixelSize() float64 {
	return c.pixelSize
}

func (c *Camera) calculate() {
	halfView := math.Tan(c.FieldOfView / 2)
	aspect := float64(c.HSize) / float64(c.VSize)
	if aspect >= 1 {
		c.halfWidth = halfView
		c.halfHeight = halfView / aspect
	} else {
		c.halfWidth = halfView * aspect
		c.halfHeight = halfView
	}
	c.pixelSize = (c.halfWidth * 2) / float64(c.HSize)
}
