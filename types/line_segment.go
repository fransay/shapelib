package types

import "math"

// LineSegment type
type LineSegment struct {
	PointA Point2D
	PointB Point2D
}

// Distance euclidean distance of a line segment
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
	midPoint := Point2D{deltaX, deltaY}
	return midPoint

}

// Bearing of the line segment in a vector space
func (l *LineSegment) Bearing() (bearing float64) {
	deltaX := l.PointB.X - l.PointA.X
	deltaY := l.PointB.Y - l.PointA.Y
	bearing = math.Atan(deltaX / deltaY)
	return bearing
}
