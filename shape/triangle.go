package shape

import (
	"math"
	"shapelib/utils"
)

type Triangle struct {
	SideOne   float64
	SideTwo   float64
	SideThree float64
	Height    float64
	Base      float64
}

const (
	isosceles   = "Isosceles"
	equilateral = "Equilateral"
	scalene     = "Scalene"
)

// AreaBySides returns the area of a triangle given the sides
func (t *Triangle) AreaBySides() (areaBySides float64) {
	semiPerimeter := (t.SideOne + t.SideTwo + t.SideThree) / 2
	areaBySides = math.Sqrt(
		semiPerimeter * (semiPerimeter - t.SideOne) *
			(semiPerimeter - t.SideTwo) * (semiPerimeter - t.SideThree))
	return areaBySides
}

// AreaByBaseHeight returns the area of a triangle given the Base and Height
func (t *Triangle) AreaByBaseHeight() (areaByBaseHeight float64) {
	areaByBaseHeight = 0.5 * t.Base * t.Height
	return areaByBaseHeight
}

// Perimeter returns the perimeter of a triangle
func (t *Triangle) Perimeter() (perimeter float64) {
	perimeter = t.SideOne + t.SideTwo + t.SideThree
	return perimeter
}

// Type returns the type of the triangle, either isosceles, scalene
// or equilateral
func (t *Triangle) Type() (typeT string) {
	if utils.CompareTrio(t.SideOne, t.SideTwo, t.SideThree) {
		typeT = equilateral
	}

	if utils.CompareDuo(t.SideOne, t.SideTwo) || utils.CompareDuo(t.SideTwo, t.SideThree) || utils.CompareDuo(t.SideOne, t.SideThree) {
		typeT = isosceles
	}

	if !utils.CompareDuo(t.SideOne, t.SideTwo) || !utils.CompareDuo(t.SideTwo, t.SideThree) || !utils.CompareDuo(t.SideOne, t.SideThree) {
		typeT = scalene
	}
	return typeT
}
