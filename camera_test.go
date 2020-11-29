package main

import (
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/world"
	"math"
)

var c world.Camera
var hSize int
var vSize int
var fieldOfView float64

func cCamerahsizeVsizeField_of_view() error {
	c = *world.NewCamera(hSize, vSize, fieldOfView)
	return nil
}

func cfield_of_view(arg1 float64) error {
	return ExpectFloatEquals(c.FieldOfView, math.Pi/arg1)
}

func chsize(arg1 int) error {
	return ExpectIntEquals(c.HSize, arg1)
}

func ctransformIdentity_matrix() error {
	return expectEqualMatrices(c.Transform, matrix.NewIdentityMatrix())
}

func cvsize(arg1 int) error {
	return ExpectIntEquals(c.VSize, arg1)
}

func field_of_view(arg1 float64) error {
	fieldOfView = math.Pi/arg1
	return nil
}

func hsize(arg1 int) error {
	hSize = arg1
	return nil
}

func vsize(arg1 int) error {
	vSize = arg1
	return nil
}

func cCamera(arg1, arg2 int, arg3 float64) error {
	c = *world.NewCamera(arg1, arg2, math.Pi/arg3)
	return nil
}

func cpixel_size(arg1 float64) error {
	return ExpectFloatEquals(c.PixelSize(), arg1)
}


func CameraFeatureContext(s *godog.ScenarioContext) {
	s.Step(`^c ← camera\(hsize, vsize, field_of_view\)$`, cCamerahsizeVsizeField_of_view)
	s.Step(`^c\.field_of_view = π\/(\d+)$`, cfield_of_view)
	s.Step(`^c\.hsize = (\d+)$`, chsize)
	s.Step(`^c\.transform = identity_matrix$`, ctransformIdentity_matrix)
	s.Step(`^c\.vsize = (\d+)$`, cvsize)
	s.Step(`^field_of_view ← π\/(\d+)$`, field_of_view)
	s.Step(`^hsize ← (\d+)$`, hsize)
	s.Step(`^vsize ← (\d+)$`, vsize)
	s.Step(`^c ← camera\((\d+), (\d+), π\/(\d+)\)$`, cCamera)
	s.Step(`^c\.pixel_size = `+complexNum+`$`, cpixel_size)
}
