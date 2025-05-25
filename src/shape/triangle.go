package shape

import (
	"math"

	"github.com/fransay/shapelib/internal/utils"
)

const (
	Isosceles   = "Isosceles"
	Equilateral = "Equilateral"
	Scalene     = "Scalene"
)

// Triangle defines a triangle given all sides/length
type Triangle struct {
	SideOne   float64
	SideTwo   float64
	SideThree float64
}

// NewTriangle initialise a new Triangle object
func NewTriangle(sideOne, sideTwo, sideThree float64) *Triangle {
	return &Triangle{
		SideOne:   sideOne,
		SideTwo:   sideTwo,
		SideThree: sideThree,
	}
}

// Area returns the computed area of triangle given sides
func (t *Triangle) Area() (areaBySides float64) {
	semiPerimeter := (t.SideOne + t.SideTwo + t.SideThree) / 2
	areaBySides = math.Sqrt(
		semiPerimeter * (semiPerimeter - t.SideOne) *
			(semiPerimeter - t.SideTwo) * (semiPerimeter - t.SideThree))
	return areaBySides
}

// Perimeter returns the perimeter of a triangle
func (t *Triangle) Perimeter() (perimeter float64) {
	perimeter = t.SideOne + t.SideTwo + t.SideThree
	return perimeter
}

// Type returns the type of the triangle, either isosceles, scalene or equilateral
func (t *Triangle) Type() (triangleType string) {
	if utils.CompareThreeFloat64(t.SideOne, t.SideTwo, t.SideThree) {
		triangleType = Equilateral
	} else if utils.CompareTwoFloat64(t.SideOne, t.SideTwo) || utils.CompareTwoFloat64(t.SideTwo, t.SideThree) || utils.CompareTwoFloat64(t.SideOne, t.SideThree) {
		triangleType = Isosceles
	} else if !utils.CompareTwoFloat64(t.SideOne, t.SideTwo) || !utils.CompareTwoFloat64(t.SideTwo, t.SideThree) || !utils.CompareTwoFloat64(t.SideOne, t.SideThree) {
		triangleType = Scalene
	}
	return triangleType
}
