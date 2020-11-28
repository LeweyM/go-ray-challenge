package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github/lewismetcalf/goRayChallenge/matrix"
	"github/lewismetcalf/goRayChallenge/object"
	"github/lewismetcalf/goRayChallenge/tuple"
	"github/lewismetcalf/goRayChallenge/world"
	"regexp"
	"strconv"
	"strings"
)

var w *world.World
var s1 *object.Sphere
var s2 *object.Sphere

func parseSphereFromTable(mm *messages.PickleStepArgument_PickleTable) *object.Sphere {
	sphere := object.NewSphere()
	material := object.NewMaterial()
	for _, row := range mm.GetRows() {
		cells := row.GetCells()
		left := cells[0].GetValue()
		right := cells[1].GetValue()
		if left == "material.color" {
			s := removeWhitespace(right)
			v := strings.Split(s[1:len(s)-1], ",")
			r, _ := strconv.ParseFloat(v[0], 64)
			g, _ := strconv.ParseFloat(v[1], 64)
			b, _ := strconv.ParseFloat(v[2], 64)
			material.SetColor(tuple.NewColor(r, g, b))
		}
		if left == "material.diffuse" {
			f, _ := strconv.ParseFloat(right, 64)
			material.SetDiffuse(f)
		}
		if left == "material.specular" {
			f, _ := strconv.ParseFloat(right, 64)
			material.SetSpecular(f)
		}
		if left == "transform" {
			s := removeWhitespace(right)
			v := strings.Split(s[8:len(s)-1], ",")
			x, _ := strconv.ParseFloat(v[0], 64)
			y, _ := strconv.ParseFloat(v[1], 64)
			z, _ := strconv.ParseFloat(v[2], 64)
			scale := matrix.NewScale(x, y, z)
			sphere.SetTransform(scale)
		}
	}
	sphere.SetMaterial(*material)
	return sphere
}

func wContainsNoObjects() error {
	objs := w.Objects()
	return ExpectTrue(len(objs) == 0, "Expected w to contain no objects")
}

func wHasNoLightSource() error {
	ok, _ := w.Light()
	return ExpectFalse(ok, "Expected w to have no light")
}

func wWorld() error {
	w = world.NewWorld()
	return nil
}

func s1SphereWith(mm *messages.PickleStepArgument_PickleTable) error {
	s1 = parseSphereFromTable(mm)
	return nil
}

func s2SphereWith(mm *messages.PickleStepArgument_PickleTable) error {
	s2 = parseSphereFromTable(mm)
	return nil
}

func wContainsS1() error {
	return ExpectEqualsSpheres(w.Objects()[0], *s1)
}

func wContainsS2() error {
	return ExpectEqualsSpheres(w.Objects()[1], *s2)
}

func wDefault_world() error {
	w = world.NewDefaultWorld()
	return nil
}

func wlightLight() error {
	_, pointLight := w.Light()
	return ExpectTrue(pointLight.Equals(&l), fmt.Sprintf("Expected %v to equal %v.", pointLight, l))
}

func xsIntersect_worldwR() error {
	xs = w.Intersect(r)
	return nil
}

func xsT(arg1 int, f float64) error {
	intersection := xs.Get(arg1)
	return ExpectFloatEquals(intersection.Time(), f)
}

func WorldContext(s *godog.ScenarioContext) {
	s.Step(`^w contains no objects$`, wContainsNoObjects)
	s.Step(`^w has no light source$`, wHasNoLightSource)
	s.Step(`^w ← world\(\)$`, wWorld)
	s.Step(`^s1 ← sphere\(\) with:$`, s1SphereWith)
	s.Step(`^s2 ← sphere\(\) with:$`, s2SphereWith)
	s.Step(`^w contains s1$`, wContainsS1)
	s.Step(`^w contains s2$`, wContainsS2)
	s.Step(`^w ← default_world\(\)$`, wDefault_world)
	s.Step(`^w\.light = light$`, wlightLight)

	s.Step(`^xs ← intersect_world\(w, r\)$`, xsIntersect_worldwR)
	s.Step(`^xs\[(\d+)\]\.t = `+complexNum+`$`, xsT)
}

func removeWhitespace(src string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(src, "")
}

func ExpectEqualsSpheres(s2 object.Sphere, s1 object.Sphere) error {
	return ExpectTrue(s2.Equals(s1), fmt.Sprintf("Expected \n%v, to equal \n%v", s2, s1))
}


