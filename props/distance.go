package props

import (
	"math"
	"shapelib/geom"
	"shapelib/tellus"
	"shapelib/utils"
)

// todo: use an pointer intera
// EuclideanDistance returns Euclidean distance between two points in a 2-dimensional space
func EuclideanDistance(pointOne, pointTwo geom.Point2D) (euclidDist float64) {
	deltaX := pointTwo.X - pointOne.X
	deltaY := pointOne.Y - pointOne.Y
	euclidDist = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return euclidDist
}

// MinkowskiDistance returns minkowski distance of points in an N dimensional space,
func MinkowskiDistance(pointA, pointB geom.Point2D, p float64) (distMinkowski float64) {
	var deltaA = math.Pow(pointA.X-pointA.Y, p)
	var deltaB = math.Pow(pointB.X-pointB.Y, p)
	distMinkowski = math.Pow(deltaA+deltaB, 1/p)
	return distMinkowski
}

// HaversineDistance returns the physical space between to locations on earth
func HaversineDistance(pointA, pointB tellus.LatLong) (haverDist float64) {
	const radiusOfEarth = 6371 // in kilometers
	latitudeDiff := utils.Deg2Rad(pointA.Latitude - pointB.Latitude)
	longitudeDiff := utils.Deg2Rad(pointA.Longitude - pointB.Longitude)
	haversine := math.Pow(math.Sinh(latitudeDiff/2), 2) + math.Cos(pointA.Latitude)*math.Cos(pointA.Latitude)*math.Pow(math.Sin(longitudeDiff/2), 2)
	centralAngle := 2 * math.Atan2(math.Sqrt(haversine), math.Sqrt(1-haversine))
	haverDist = radiusOfEarth * centralAngle
	return haverDist
}

// ChebyshevDistance returns the max distance along any coordinate dimension
func ChebyshevDistance(pointA, pointB geom.Point2D) (chebDist float64) {
	chebDist = math.Max(pointB.X-pointA.X, pointB.Y-pointA.Y)
	return chebDist
}
