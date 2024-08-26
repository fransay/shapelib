package shape

import (
	"math"
	"shapelib/utils"
)

// Triangle defines a triangle given all sides/length
type Triangle struct {
	SideOne   float64
	SideTwo   float64
	SideThree float64 // todo: remove all sides
	Height    float64
	Base      float64
}

// TriangleC defines a triangle given all coordinates
type TriangleC [3][2]float64

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
	if utils.CompareThreeFloat64(t.SideOne, t.SideTwo, t.SideThree) {
		typeT = equilateral
	}

	if utils.CompareTwoFloat64(t.SideOne, t.SideTwo) || utils.CompareTwoFloat64(t.SideTwo, t.SideThree) || utils.CompareTwoFloat64(t.SideOne, t.SideThree) {
		typeT = isosceles
	}

	if !utils.CompareTwoFloat64(t.SideOne, t.SideTwo) || !utils.CompareTwoFloat64(t.SideTwo, t.SideThree) || !utils.CompareTwoFloat64(t.SideOne, t.SideThree) {
		typeT = scalene
	}
	return typeT
}

// Center returns a point at the center of the triangle
func (tc *TriangleC) Center() (pt [2]float64) {
	var x float64
	var y float64
	for _, pnt := range tc {
		x += pnt[0]
		y += pnt[1]
	}
	pt = [2]float64{x / 3, y / 3}
	return pt
}

// todo: thirdPoint() -> ThirdPoint takes 2 points and checks which point is the 3rd in the Triangle
// todo: linearRings() -> LinearRings returns the coordinates of the linear rings
// todo: Area() -> Area returns twice the area of the oriented triangle (a,b,c), i.e. the area is positive if the triangle is oriented counterclockwise.
// todo: Contains() -> Accepts a new triangle as a parameter and checks to see if a triangle is in the TriangleC
