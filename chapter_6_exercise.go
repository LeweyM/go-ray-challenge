package main

import (
	"fmt"
	"github/lewismetcalf/goRayChallenge/light"
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/object"
	"github/lewismetcalf/goRayChallenge/tuple"
	"github/lewismetcalf/goRayChallenge/world"
	"math"
	"time"
)

func run6(h, v int) {
	floor := object.NewSphere()
	floor.SetTransform(matrix.NewScale(10, 0.01, 10))
	floorMaterial := object.NewMaterial()
	floorMaterial.SetColor(tuple.NewColor(1, 0.9, 0.9))
	floorMaterial.SetSpecular(0)
	floor.SetMaterial(floorMaterial)

	leftWall := object.NewSphere()
	leftWall.SetTransform(matrix.NewTranslation(0, 0, 5).
		Multiply(matrix.NewRotationY(-math.Pi / 4)).
		Multiply(matrix.NewRotationX(math.Pi / 2)).
		Multiply(matrix.NewScale(10, 0.01, 10)))
	leftWall.SetMaterial(floorMaterial)

	rightWall := object.NewSphere()
	rightWall.SetTransform(matrix.NewTranslation(0, 0, 5).
		Multiply(matrix.NewRotationY(math.Pi / 4)).
		Multiply(matrix.NewRotationX(math.Pi / 2)).
		Multiply(matrix.NewScale(10, 0.01, 10)))
	rightWall.SetMaterial(floorMaterial)

	middle := object.NewSphere()
	middle.SetTransform(matrix.NewTranslation(-0.5, 1, -1.5))
	middleMaterial := object.NewMaterial()
	middleMaterial.SetColor(tuple.NewColor(0.1, 1, 0.5))
	middleMaterial.SetDiffuse(0.7)
	middleMaterial.SetSpecular(0.3)
	middle.SetMaterial(middleMaterial)

	right := object.NewSphere()
	right.SetTransform(matrix.NewTranslation(1.5, 0.5, -0.5).
		Multiply(matrix.NewScale(0.5, 0.5, 0.9)))
	rightMaterial := object.NewMaterial()
	rightMaterial.SetColor(tuple.NewColor(0.5, 1, 0.1))
	rightMaterial.SetDiffuse(0.7)
	rightMaterial.SetSpecular(0.3)
	right.SetMaterial(rightMaterial)

	left := object.NewSphere()
	left.SetTransform(matrix.NewTranslation(-2.5, 0.33, -0.75).
		Multiply(matrix.NewScale(0.33, 0.33, 0.33)))
	leftMaterial := object.NewMaterial()
	leftMaterial.SetColor(tuple.NewColor(1, 0.8, 0.1))
	leftMaterial.SetDiffuse(0.7)
	leftMaterial.SetSpecular(0.3)
	left.SetMaterial(leftMaterial)

	newWorld := world.NewWorld()
	newWorld.SetLight(*light.NewPointLight(*tuple.NewPoint(-10, 10, -10), *tuple.NewColor(1,1,1)))
	newWorld.SetObjects(*floor, *leftWall, *rightWall, *left, *right, *middle)
	camera := world.NewCamera(h, v, math.Pi/3)
	camera.Transform = matrix.ViewTransform(
		*tuple.NewPoint(0, 1.5, -8),
		*tuple.NewPoint(0, 1, 0),
		*tuple.NewVector(0, 1, 0),
	)

	start := time.Now()
	canvas := camera.Render(newWorld)
	finish := time.Now()
	fmt.Printf("\nProcessing time: %v\n", finish.Sub(start))

	writeToOutput(canvas.ToPPM())


}
