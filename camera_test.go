package main

import (
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/canvas"
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


func ctransformRotation_yTranslation(arg1, arg2, arg3, arg4 float64) error {
	newRotationY := matrix.NewRotationY(math.Pi / arg1)
	c.Transform = newRotationY.Multiply(matrix.NewTranslation(arg2, arg3, arg4))
	return nil
}

func rRay_for_pixelc(arg1, arg2 int) error {
	r = c.RayForPixel(arg1, arg2)
	return nil
}

func rdirectionVector(arg1, arg2, arg3 float64) error {
	direction := r.Direction()
	return ExpectVector(&direction, arg1, arg2, arg3)
}

func rdirectionRootMinusRootVector(arg1, arg2, arg3 float64) error {
	direction := r.Direction()
	return ExpectVector(&direction, math.Sqrt(2)/arg1, arg2, -math.Sqrt(2)/arg3)
}

func roriginPoint(arg1, arg2, arg3 float64) error {
	origin := r.Origin()
	return ExpectPoint(&origin, arg1, arg2, arg3)
}

func ctransformView_transformfromToUp() error {
	from := tuples["from"]
	up := tuples["up"]
	to := tuples["to"]
	c.Transform = matrix.ViewTransform(*from, *to, *up)
	return nil
}

var image canvas.Canvas

func imageRendercW() error {
	image = *c.Render(w)
	return nil
}

func pixel_atimageColor(arg1, arg2 int, arg3, arg4, arg5 float64) error {
	return ExpectColor(image.PixelAt(arg1, arg2), arg3, arg4, arg5)
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

	s.Step(`^c\.transform ← rotation_y\(π\/(\d+)\) \* translation\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, ctransformRotation_yTranslation)
	s.Step(`^r ← ray_for_pixel\(c, (\d+), (\d+)\)$`, rRay_for_pixelc)
	s.Step(`^r\.direction = vector\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, rdirectionVector)
	s.Step(`^r\.direction = vector\(√2\/(\d+), `+complexNum+`, -√2\/(\d+)\)$`, rdirectionRootMinusRootVector)
	s.Step(`^r\.origin = point\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, roriginPoint)

	s.Step(`^c\.transform ← view_transform\(from, to, up\)$`, ctransformView_transformfromToUp)
	s.Step(`^image ← render\(c, w\)$`, imageRendercW)
	s.Step(`^pixel_at\(image, `+complexNum+`, `+complexNum+`\) = color\(`+complexNum+`, `+complexNum+`, `+complexNum+`\)$`, pixel_atimageColor)

}
