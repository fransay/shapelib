package types

import (
	"errors"
	"math"
	"shapelib/functs"
	"shapelib/shape"
	"shapelib/utils"
)

var collinearityError = errors.New("collinearity error, points don't lie on the same line")

// Point2D point type characterise by x and y coordination
type Point2D struct {
	X float64
	Y float64
}

// NewPoint2D returns a new point. i.e constructor
func NewPoint2D(x, y int) *Point2D {
	return &Point2D{float64(x), float64(y)}
}

// DistanceTo returns the distance between self and other point
func (p *Point2D) DistanceTo(pt Point2D) (distance float64) {
	distance = math.Sqrt(math.Pow(p.X-pt.X, 2) + math.Pow(p.Y-pt.Y, 2))
	return distance
}

// BearingTo returns the bearing between self and other point.
func (p *Point2D) BearingTo(pt Point2D) (bearing float64) {
	bearing = math.Atan2(p.Y-pt.Y, p.X-pt.X)
	return bearing
}

// Size of a point2D is dimensionless i.e, it has no size, hence return 0.0
func (p *Point2D) Size() (size float64) {
	return 0.0
}

// Len return the number of points in the receiver, always equal to 1
func (p *Point2D) Len() float64 {
	return 1.0
}

// Coordinates of point2D type returned in an array
func (p *Point2D) Coordinates() (coordinates []float64) {
	coordinates = []float64{p.X, p.Y}
	return coordinates
}

// Buffer returns a circle with a given radius where this.Point2D is the center
func (p *Point2D) Buffer(radius float64) *shape.Circle {
	return &shape.Circle{Radius: radius, Centroid: *p}
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

// IsEqual returns a boolean that shows where Point instance is equal to another object of Point type.
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

// Determine the gradient between two points.
func slope(pointOne, pointTwo Point2D) float64 {
	grad := (pointTwo.Y - pointOne.Y) / (pointTwo.X - pointOne.X)
	return grad
}

// IsCollinear checks if a set of the points lie on the same line
func IsCollinear(points ...Point2D) (isCollinear bool, err error) {
	var slopeList []float64
	numberOfPoints := len(points)
	currentIndex := 0

	if numberOfPoints <= 2 {
		return false, collinearityError
	}
	for currentIndex < numberOfPoints-2 {
		nextIndex := currentIndex + 1
		firstSlope := slope(points[currentIndex], points[nextIndex])
		secondSlope := slope(points[nextIndex], points[nextIndex+1])
		slopeList = append(slopeList, firstSlope, secondSlope)
		currentIndex++
	}

	if utils.IsArrayElementFloat64(slopeList) {
		isCollinear = true
	} else {
		isCollinear = false
	}
	return isCollinear, collinearityError
}
