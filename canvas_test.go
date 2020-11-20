package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"goRayChallenge/canvas"
)

var canvases map[string] *canvas.Canvas

func cCanvas(c string, arg1, arg2 int) error {
	canvases[c] = canvas.NewCanvas(arg1, arg2)
	return nil
}

func cheight(c string, arg1 int) error {
	return ExpectTrue(canvases[c].Height() == arg1,
		fmt.Sprintf("expected height to be %d, got %d", arg1, canvases[c].Height()))
}

func cwidth(c string, arg1 int) error {
	return ExpectTrue(canvases[c].Width() == arg1,
		fmt.Sprintf("expected width to be %d, got %d", arg1, canvases[c].Width()))
}

func everyPixelOfCIsColor(c string, arg1, arg2, arg3 float64) error {
	canvas := canvases[c]
	for i := 0; i < canvas.Width(); i++ {
		for j := 0; j < canvas.Height(); j++ {
			pixel := canvas.PixelAt(i, j)
			err := ExpectColor(&pixel, arg1, arg2, arg3)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


func InitializeCanvasScenario(s *godog.ScenarioContext) {
	s.Step(`^`+VarName+` ← canvas\(`+Number+`, `+Number+`\)$`, cCanvas)
	s.Step(`^`+VarName+`\.height = `+Number+`$`, cheight)
	s.Step(`^`+VarName+`\.width = `+Number+`$`, cwidth)
	s.Step(`^every pixel of `+VarName+` is `+Color+`$`, everyPixelOfCIsColor)

	s.BeforeScenario(func(sc *godog.Scenario) {
		canvases = make(map[string]*canvas.Canvas)
	})
}
