package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github/lewismetcalf/goRayChallenge/object"
)

var mat object.Material

func mMaterial() error {
	mat = *object.NewMaterial()
	return nil
}

func mcolorColor(arg1, arg2, arg3 float64) error {
	return ExpectColor(mat.Color(), arg1, arg2, arg3)
}

func mambient(arg1 float64) error {
	return ExpectFloatEquals(mat.Ambient(), arg1)
}

func mdiffuse(arg1 float64) error {
	return ExpectFloatEquals(mat.Diffuse(), arg1)
}

func mshininess(arg1 float64) error {
	return ExpectFloatEquals(mat.Shininess(), arg1)
}

func mspecular(arg1 float64) error {
	return ExpectFloatEquals(mat.Specular(), arg1)
}

func setMToAmbient(arg1 float64) error {
	mat.SetAmbient(arg1)
	return nil
}

func setSMaterialToM() error {
	s.SetMaterial(mat)
	return nil
}

func sMaterialEqualsM() error {
	return ExpectEqualMaterials(*s.Material(), mat)
}

func mEqualsNewMaterial() error {
	return ExpectEqualMaterials(*object.NewMaterial(), mat)
}

func setMToSMaterial() error {
	mat = *s.Material()
	return nil
}

func MaterialsContext(s *godog.ScenarioContext) {
	s.Step(`^m ← material\(\)$`, mMaterial)
	s.Step(`^m\.ambient = `+complexNum+`$`, mambient)
	s.Step(`^m\.color = color\((\d+), (\d+), (\d+)\)$`, mcolorColor)
	s.Step(`^m\.diffuse = `+complexNum+`$`, mdiffuse)
	s.Step(`^m\.shininess = `+complexNum+`$`, mshininess)
	s.Step(`^m\.specular = `+complexNum+`$`, mspecular)
	// sphere material
	s.Step(`^m = material\(\)$`, mEqualsNewMaterial)
	s.Step(`^m ← s\.material$`, setMToSMaterial)
	s.Step(`^m\.ambient ← (\d+)$`, setMToAmbient)
	s.Step(`^s\.material ← m$`, setSMaterialToM)
	s.Step(`^s\.material = m$`, sMaterialEqualsM)

}

func ExpectEqualMaterials(a object.Material, b object.Material) error {
	return ExpectTrue(a.Equals(b), fmt.Sprintf("Expected %v to equal %v", a, b))
}

