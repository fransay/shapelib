package props

import (
	"math"

	"github.com/fransay/shapelib/internal/utils"
	"github.com/fransay/shapelib/src/geog"
	"github.com/fransay/shapelib/src/geom"
)

// EuclideanDistance returns Euclidean distance between two points in a 2-dimensional space
func EuclideanDistance(pointOne, pointTwo geom.Point2D) (euclideanDistance float64) {
	deltaX := pointTwo.X - pointOne.X
	deltaY := pointOne.Y - pointOne.Y
	euclideanDistance = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return euclideanDistance
}

// MinkowskiDistance returns minkowski distance of points in an N dimensional space,
func MinkowskiDistance(pointA, pointB geom.Point2D, p float64) (minkowskiDistance float64) {
	var deltaA = math.Pow(pointA.X-pointA.Y, p)
	var deltaB = math.Pow(pointB.X-pointB.Y, p)
	minkowskiDistance = math.Pow(deltaA+deltaB, 1/p)
	return minkowskiDistance
}

// HaversineDistance returns the physical space between to locations on earth
func HaversineDistance(pointA, pointB geog.LatLong) (haversineDistance float64) {
	const radiusOfEarth = 6371 // in kilometers
	latitudeDiff := utils.Deg2Rad(pointA.Latitude - pointB.Latitude)
	longitudeDiff := utils.Deg2Rad(pointA.Longitude - pointB.Longitude)
	haversine := math.Pow(math.Sinh(latitudeDiff/2), 2) + math.Cos(pointA.Latitude)*math.Cos(pointA.Latitude)*math.Pow(math.Sin(longitudeDiff/2), 2)
	centralAngle := 2 * math.Atan2(math.Sqrt(haversine), math.Sqrt(1-haversine))
	haversineDistance = radiusOfEarth * centralAngle
	return haversineDistance
}

// ChebyshevDistance returns the max distance along any coordinate dimension
func ChebyshevDistance(pointA, pointB geom.Point2D) (ChebyshevDistance float64) {
	ChebyshevDistance = math.Max(pointB.X-pointA.X, pointB.Y-pointA.Y)
	return ChebyshevDistance
}
