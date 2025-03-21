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

const SumOfAnglesInTriangle = 180

// Triangle defines a triangle given all sides/length
type Triangle struct {
	SideOne   float64
	SideTwo   float64
	SideThree float64 // todo: remove all sides
	Height    float64
	Base      float64
}

// NewTriangle initialise a new triangle object
func NewTriangle(sideOne, sideTwo, sideThree, height, base float64) *Triangle {
	return &Triangle{
		SideOne:   sideOne,
		SideTwo:   sideTwo,
		SideThree: sideThree,
		Height:    height,
		Base:      base,
	}
}

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

// Type returns the type of the triangle, either isosceles, scalene or equilateral
func (t *Triangle) Type() (triangleType string) {
	if utils.CompareThreeFloat64(t.SideOne, t.SideTwo, t.SideThree) {
		triangleType = Equilateral
	}

	if utils.CompareTwoFloat64(t.SideOne, t.SideTwo) || utils.CompareTwoFloat64(t.SideTwo, t.SideThree) || utils.CompareTwoFloat64(t.SideOne, t.SideThree) {
		triangleType = Isosceles
	}

	if !utils.CompareTwoFloat64(t.SideOne, t.SideTwo) || !utils.CompareTwoFloat64(t.SideTwo, t.SideThree) || !utils.CompareTwoFloat64(t.SideOne, t.SideThree) {
		triangleType = Scalene
	}
	return triangleType
}

// PointInTriangle checks if a point lies in a triangle
func (t *Triangle) PointInTriangle(x, y float64) (in bool) {
	// x,y defines the coordinates of the points
	in = false
	// todo: put logic here
	return in
}

// ContainsTriangle checks if a triangle completely lies in instant triangle
func (t *Triangle) ContainsTriangle() (contain bool) {
	contain = false
	//todo: put logic here
	return contain
}
