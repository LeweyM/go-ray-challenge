package world

import (
	"github/lewismetcalf/goRayChallenge/canvas"
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/ray"
	"github/lewismetcalf/goRayChallenge/tuple"
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

func (c *Camera) RayForPixel(px, py int) ray.Ray {
	xOffset := (float64(px) + 0.5) * c.pixelSize
	yOffset := (float64(py) + 0.5) * c.pixelSize

	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	transformInverse := c.Transform.Invert()
	pixel := transformInverse.MultiplyTuple(tuple.NewPoint(worldX, worldY, -1))
	origin := transformInverse.MultiplyTuple(tuple.NewPoint(0,0,0))
	direction := pixel.Subtract(&origin).Normalize()

	return ray.NewRay(&origin, &direction)
}

func (c *Camera) Render(w *World) *canvas.Canvas {
	image := canvas.NewCanvas(c.HSize, c.VSize)

	for y := 0; y < c.VSize; y++ {
		for x := 0; x < c.HSize; x++ {
			pixelRay := c.RayForPixel(x, y)
			color := w.ColorAt(pixelRay)
			//println(fmt.Sprintf("%v", color))
			image.WritePixel(x, y, color)
		}
	}

	return image
}
