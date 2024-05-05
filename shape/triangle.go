package shape

import (
	"math"
	"shapelib/utils"
)

type Triangle struct {
	sideOne   float64
	sideTwo   float64
	sideThree float64
	height    float64
	base      float64
}

const (
	isosceles   = "Isosceles"
	equilateral = "Equilateral"
	scalene     = "Scalene"
)

// AreaBySides returns the area of a triangle given the sides
func (t *Triangle) AreaBySides() (areaBySides float64) {
	semiPerimeter := (t.sideOne + t.sideTwo + t.sideThree) / 2
	areaBySides = math.Sqrt(
		semiPerimeter * (semiPerimeter - t.sideOne) *
			(semiPerimeter - t.sideTwo) * (semiPerimeter - t.sideThree))
	return areaBySides
}

// AreaByBaseHeight returns the area of a triangle given the base and height
func (t *Triangle) AreaByBaseHeight() (areaByBaseHeight float64) {
	areaByBaseHeight = 0.5 * t.base * t.height
	return areaByBaseHeight
}

// Perimeter returns the perimeter of a triangle
func (t *Triangle) Perimeter() (perimeter float64) {
	perimeter = t.sideOne + t.sideTwo + t.sideThree
	return perimeter
}

// Type returns the type of the triangle, either isosceles, scalene
// or equilateral
func (t *Triangle) Type() (typeT string) {
	if utils.CompareTrio(t.sideOne, t.sideTwo, t.sideThree) {
		typeT = equilateral
	}

	if utils.CompareDuo(t.sideOne, t.sideTwo) ||
		utils.CompareDuo(t.sideTwo, t.sideThree) ||
		utils.CompareDuo(t.sideOne, t.sideThree) {
		typeT = isosceles
	}

	if !utils.CompareDuo(t.sideOne, t.sideTwo) ||
		!utils.CompareDuo(t.sideTwo, t.sideThree) ||
		!utils.CompareDuo(t.sideOne, t.sideThree) {
		typeT = scalene
	}

	return typeT
}
