package shape

import (
	"math"

	"github.com/fransay/shapelib/src/geom"
)

// TriangleByCoordinates defines a triangle by coordinates.
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

// Area returns the computed area of triangle by coordinates.
func (tc *TriangleByCoordinates) Area() (area float64) {
	pt1, pt2, pt3 := tc.PointOne, tc.PointTwo, tc.PointThree
	area = math.Abs(pt1.X*(pt2.Y-pt3.Y)+pt2.X*(pt3.Y-pt1.Y)+pt3.X*(pt1.Y-pt2.Y)) / 2
	return area
}

// Perimeter returns the total distance around triangle
func (tc *TriangleByCoordinates) Perimeter(a, b, c float64) (perimeter float64) {
	return a + b + c
}
