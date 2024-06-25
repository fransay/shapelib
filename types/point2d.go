package types

import (
	"math"
	"shapelib/functs"
	"shapelib/utils"
)

// Point2D point type
// characterise by x and y coordination
type Point2D struct {
	X float64
	Y float64
}

// Size of a point2D is dimensionless
// i.e, it has no size, hence return 0.0
func (p *Point2D) Size() (size float64) {
	return 0.0
}

// Coordinates of point2D type returned in an array
func (p *Point2D) Coordinates() (coordinates []float64) {
	coordinates = []float64{p.X, p.Y}
	return coordinates
}

// Polar returns a polar coordinate of a cartesian representation
// assuming the start coordinate is the origin (x:0, y:0)
func (p *Point2D) Polar() (distanceR float64, bearing float64) {
	distanceR = math.Pow(p.X, 2) + math.Pow(p.Y, 2)
	bearing = math.Atan(p.Y / p.X)
	return distanceR, bearing
}

// Translate2D returns a transformed 2D point undergone a translation transformation
func (p *Point2D) Translate2D(shiftVector Point2D) (translate Point2D) {
	translateX := p.X + shiftVector.X
	translateY := p.Y + shiftVector.Y
	translate = Point2D{translateX, translateY}
	return translate
}

// Rotate returns a transformed point undergone a rotation transformation
func (p *Point2D) Rotate(rotationAngle float64) (rotate Point2D) {
	angleRad := utils.Deg2Rad(rotationAngle)
	rotate.X = p.X + math.Cos(angleRad) - p.Y*math.Sin(angleRad)
	rotate.Y = p.X*math.Sin(angleRad) + p.Y*math.Cos(angleRad)
	return rotate
}

// Scale returns a transformed point undergone a scaling transformation
func (p *Point2D) Scale(scalarVector []float64) (scale Point2D) {
	scale.X = p.X * scalarVector[0]
	scale.Y = p.Y * scalarVector[1]
	return scale
}

// IsEqual returns a boolean that shows where Point instance is equal
// to another object of Point type.
func (p *Point2D) IsEqual(point Point2D) (isEqual bool) {
	isEqual = point.X == p.X && point.Y == p.Y
	return isEqual
}

// pointOnLine determines if a point falls on a line segment or not
func (p *Point2D) pointOnLine(segment LineSegment) (onLine bool) {
	var area = functs.AreaCoordinates(*p, segment.PointA, segment.PointB)
	onLine = utils.IsClose(area, 0, 0.001)
	return onLine
}
