package types

import "math"

type LineSegment struct {
	PointA Point
	PointB Point
}

// Distance :euclidean distance of a line segment
func (l *LineSegment) Distance() (dist float64) {
	deltaX := l.PointB.X - l.PointB.X
	deltaY := l.PointB.Y - l.PointB.Y
	dist = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return dist

}

// MidPoint :midpoint of a line segment
func (l *LineSegment) MidPoint() Point {
	deltaX := (l.PointB.X + l.PointB.X) / 2
	deltaY := (l.PointB.Y + l.PointB.Y) / 2
	midPoint := Point{deltaX, deltaY}
	return midPoint

}

// Bearing :bearing of the line segment in space
func (l *LineSegment) Bearing() (bear float64) {
	deltaX := l.PointB.X - l.PointB.X
	deltaY := l.PointB.Y - l.PointB.Y
	bear = math.Atan(deltaX / deltaY)
	return bear
}
