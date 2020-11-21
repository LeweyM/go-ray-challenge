package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/canvas"
	"github/lewismetcalf/goRayChallenge/tuple"
	"strings"
)

var canvases map[string]*canvas.Canvas
var ppms map[string]string

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
			err := ExpectColor(pixel, arg1, arg2, arg3)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func setEveryPixelOfCIsColor(c string, arg1, arg2, arg3 float64) error {
	canvas := canvases[c]
	color := tuple.NewColor(arg1, arg2, arg3)
	for i := 0; i < canvas.Width(); i++ {
		for j := 0; j < canvas.Height(); j++ {
			canvas.WritePixel(i, j, color)
		}
	}
	return nil
}

func pixel_atcRed(c string, arg1, arg2 int, r string) error {
	canvas := canvases[c]
	color := colors[r]
	pixel := canvas.PixelAt(arg1, arg2)
	return ExpectColor(pixel, color.Red(), color.Green(), color.Blue())
}

func write_pixelcRed(c string, arg1, arg2 int, r string) error {
	canvas := canvases[c]
	color := colors[r]
	canvas.WritePixel(arg1, arg2, color)
	return nil
}

func linesOfPpmAre(arg1, arg2 int, ppm string, document *godog.DocString) error {
	toPPM := ppms[ppm]
	section := getLines(arg1, arg2, toPPM)
	return ExpectTrue(document.Content == section, fmt.Sprintf("Expected:\n%s\nGot:\n%s", document.Content, section))
}

func getLines(arg1 int, arg2 int, str string) string {
	lines := strings.Split(str, "\n")
	return strings.Join(lines[arg1-1:arg2], "\n")
}

func ppmCanvas_to_ppmc(ppm, c string) error {
	canvas := canvases[c]
	ppms[ppm] = canvas.ToPPM()
	return nil
}

func InitializeCanvasScenario(s *godog.ScenarioContext) {
	s.Step(`^`+VarName+` ← canvas\(`+Number+`, `+Number+`\)$`, cCanvas)
	s.Step(`^`+VarName+`\.height = `+Number+`$`, cheight)
	s.Step(`^`+VarName+`\.width = `+Number+`$`, cwidth)
	s.Step(`^every pixel of `+VarName+` is `+Color+`$`, everyPixelOfCIsColor)
	s.Step(`^every pixel of `+VarName+` is set to `+ Color+`$`, setEveryPixelOfCIsColor)
	s.Step(`^pixel_at\(`+VarName+`, `+Number+`, `+Number+`\) = `+VarName+`$`, pixel_atcRed)
	s.Step(`^write_pixel\(`+VarName+`, `+Number+`, `+Number+`, `+VarName+`\)$`, write_pixelcRed)
	s.Step(`^lines `+Number+`-`+Number+` of `+VarName+` are$`, linesOfPpmAre)
	s.Step(`^`+VarName+` ← canvas_to_ppm\(`+VarName+`\)$`, ppmCanvas_to_ppmc)

	s.BeforeScenario(func(sc *godog.Scenario) {
		canvases = make(map[string]*canvas.Canvas)
		ppms = make(map[string]string)
	})
}
