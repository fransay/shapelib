package shape

import (
	"math"

	"github.com/fransay/shapelib/internal/utils"
	"github.com/fransay/shapelib/src/geom"
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

// AreaBySides returns the area of a triangle given the sides, using t
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
	} else if utils.CompareTwoFloat64(t.SideOne, t.SideTwo) || utils.CompareTwoFloat64(t.SideTwo, t.SideThree) || utils.CompareTwoFloat64(t.SideOne, t.SideThree) {
		triangleType = Isosceles
	} else if !utils.CompareTwoFloat64(t.SideOne, t.SideTwo) || !utils.CompareTwoFloat64(t.SideTwo, t.SideThree) || !utils.CompareTwoFloat64(t.SideOne, t.SideThree) {
		triangleType = Scalene
	}
	return triangleType
}

// TriangleByCoordinates defined by cartesian coordinates.
type TriangleByCoordinates struct {
	PointOne   geom.Point2D
	PointTwo   geom.Point2D
	PointThree geom.Point2D
}

// NewTriangleByCoordinates initialises a new TriangleByCoordinates object
func NewTriangleByCoordinates(pointOne, pointTwo, pointThree geom.Point2D) *TriangleByCoordinates {
	return &TriangleByCoordinates{
		PointOne:   pointOne,
		PointTwo:   pointTwo,
		PointThree: pointThree,
	}
}

// Center returns the centroid of a triangle
func (tc *TriangleByCoordinates) Center() (pt geom.Point2D) {
	x := tc.PointOne.X + tc.PointTwo.X + tc.PointThree.X
	y := tc.PointOne.Y + tc.PointTwo.Y + tc.PointThree.Y
	pt.X = x / 3
	pt.Y = y / 3
	return pt
}

// Area returns the area of triangle given the coordinates
func (tc *TriangleByCoordinates) Area() (area float64) {
	pt1 := tc.PointOne
	pt2 := tc.PointTwo
	pt3 := tc.PointThree
	area = math.Abs(pt1.X*(pt2.Y-pt3.Y)+pt2.X*(pt3.Y-pt1.Y)+pt3.X*(pt1.Y-pt2.Y)) / 2
	return area
}
