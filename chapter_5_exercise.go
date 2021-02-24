package main

import (
	"fmt"
	canvas2 "github/lewismetcalf/goRayChallenge/canvas"
	light2 "github/lewismetcalf/goRayChallenge/light"
	"github/lewismetcalf/goRayChallenge/object"
	"github/lewismetcalf/goRayChallenge/ray"
	"github/lewismetcalf/goRayChallenge/tuple"
	"os"
	"time"
)

func run5() {
	rayOrigin := tuple.NewPoint(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0

	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := canvas2.NewCanvas(canvasPixels, canvasPixels)
	shape := object.NewSphere()
	material := object.NewMaterial()
	material.SetColor(tuple.NewColor(1, 0.2, 1))
	shape.SetMaterial(material)

	lightPosition := tuple.NewPoint(-10, 10, -10)
	lightColor := tuple.NewColor(1, 1, 1)
	light := light2.NewPointLight(*lightPosition, *lightColor)

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)

			position := tuple.NewPoint(worldX, worldY, wallZ)

			normalize := position.Subtract(rayOrigin).Normalize()
			r := ray.NewRay(rayOrigin, &normalize)
			hit, xs := shape.Intersects(r)
			if hit {
				ok, intersection := xs.Hit()
				if ok {
					point := r.Position(intersection.Time())
					obj := intersection.Object()
					normal := obj.NormalAt(&point)
					eye := r.Direction().Negate()
					m := obj.Material()
					color := m.Lighting(light, &point, eye, &normal, false)
					canvas.WritePixel(x, y, &color)
				}
			}
		}
	}

	writeToOutput(canvas.ToPPM())
}

func writeToOutput(ppm string) {
	f, err := os.Create("output/" + time.Now().String() + ".ppm")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(ppm + "\n")

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("done")
}
