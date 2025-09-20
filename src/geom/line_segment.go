package geom

import (
	"math"

	"github.com/fransay/shapelib/internal/utils"
)

// Dimension of a line segment === 1
const Dimension = 1

// LineSegment type
type LineSegment struct {
	PointA Point2D
	PointB Point2D
}

// Construct new line segment.
func NewLineSegment(pointA Point2D, pointB Point2D) *LineSegment {
	return &LineSegment{
		PointA: pointA,
		PointB: pointB,
	}

}

// Distance Euclidean distance of a line segment
func (l *LineSegment) Distance() (dist float64) {
	deltaX := math.Abs(l.PointB.X - l.PointA.X)
	deltaY := math.Abs(l.PointB.Y - l.PointA.Y)
	dist = math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))
	return dist

}

// MidPoint midpoint of a line segment
func (l *LineSegment) MidPoint() Point2D {
	deltaX := (l.PointB.X + l.PointA.X) / 2
	deltaY := (l.PointB.Y + l.PointA.Y) / 2
	midPoint := Point2D{X: deltaX, Y: deltaY}
	return midPoint

}

// Bearing of the line segment in a vector space
func (l *LineSegment) Bearing() (bearing float64) {
	deltaX := l.PointB.X - l.PointA.X
	deltaY := l.PointB.Y - l.PointA.Y
	bearing = math.Atan(deltaX / deltaY)
	return bearing
}

// Rotate return a transformed line segment under scale transformation
func (l *LineSegment) Rotate(rotationAngle float64) (ls LineSegment) {
	// compute is done on object prior to garbage collection
	angleRad := utils.Deg2Rad(rotationAngle)
	ls.PointA.X = l.PointA.X*math.Cos(angleRad) - l.PointA.Y*math.Sin(angleRad)
	ls.PointA.Y = l.PointA.X*math.Sin(angleRad) - l.PointA.Y*math.Cos(angleRad)
	ls.PointB.X = l.PointB.X*math.Cos(angleRad) - l.PointB.Y*math.Sin(angleRad)
	ls.PointB.Y = l.PointB.X*math.Sin(angleRad) - l.PointB.Y*math.Cos(angleRad)
	ls = *l
	return ls

}

// Scale return a transformed line segment under scale transformation
func (l *LineSegment) Scale(scalarVector []float64) (ls LineSegment) {
	// compute is done on object prior to garbage collection
	ls.PointA.X = l.PointA.X * scalarVector[0]
	ls.PointA.Y = l.PointA.Y * scalarVector[1]
	ls.PointB.X = l.PointB.X * scalarVector[0]
	ls.PointB.Y = l.PointB.Y * scalarVector[1]
	ls = *l
	return ls

}

// Gradient return the slope of line segment
func (l *LineSegment) Gradient() float64 {
	deltaY := l.PointB.Y - l.PointA.Y
	deltaX := l.PointB.X - l.PointB.Y
	return deltaY / deltaX
}
