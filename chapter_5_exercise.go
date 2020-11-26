package main

import (
	"fmt"
	canvas2 "github/lewismetcalf/goRayChallenge/canvas"
	"github/lewismetcalf/goRayChallenge/object"
	"github/lewismetcalf/goRayChallenge/ray"
	"github/lewismetcalf/goRayChallenge/tuple"
	"os"
)

func run() {
	rayOrigin := tuple.NewPoint(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0

	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := canvas2.NewCanvas(canvasPixels, canvasPixels)
	color := tuple.NewColor(1, 0 ,0)
	shape := object.NewSphere()

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)

			position := tuple.NewPoint(worldX, worldY, wallZ)

			r := ray.NewRay(rayOrigin, position.Subtract(rayOrigin).Normalize())
			hit, _ := shape.Intersects(r)
			if hit {
				canvas.WritePixel(x, y, color)
			}
		}
	}

	f, err := os.Create("output/output.ppm")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(canvas.ToPPM() + "\n")

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("done")
}
