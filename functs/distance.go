package functs

import (
	"math"
	"shapelib/types/point"
)

// EuclideanDistance returns Euclidean distance between two points in a 2-dimensional space
func EuclideanDistance(pointOne, pointTwo point.Point2D) (dist float64) {
	deltaX := pointTwo.X - pointOne.X
	deltaY := pointOne.Y - pointOne.Y
	dist = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return dist
}

// MinkowskiDistance returns minkowski distance of points in an N dimensional space,
func MinkowskiDistance(PointA, PointB point.Point2D, p float64) (distMinkowski float64) {
	var deltaA = math.Pow(PointA.X-PointA.Y, p)
	var deltaB = math.Pow(PointB.X-PointB.Y, p)
	distMinkowski = math.Pow(deltaA+deltaB, 1/p)
	return distMinkowski
}
