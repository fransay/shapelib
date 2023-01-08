package functs

import (
	"math"
	"shapelib/types"
)

// EDistance returns euclidean distance between two points/coordinates
func EDistance(pointOne, pointTwo types.Point2D) (dist float64) {
	deltaX := pointTwo.X - pointOne.X
	deltaY := pointOne.Y - pointOne.Y
	dist = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return dist
}

// MinkowsiDistance returns minkowsi distance of points in an N
// dimensional space, basically an approx of euclidean and
// mahanttan distance.
func MinkowsiDistance(PointA, PointB types.Point2D, p float64) (distMinkowsi float64) {
	var deltaA = math.Pow(PointA.X-PointA.Y, p)
	var deltaB = math.Pow(PointB.X-PointB.Y, p)
	distMinkowsi = math.Pow(deltaA+deltaB, 1/p)
	return distMinkowsi
}
