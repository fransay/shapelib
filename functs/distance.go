package functs

import (
	"math"
	"shapelib/types"
	"shapelib/utils"
)

// EuclideanDistance returns Euclidean distance between two points in a 2-dimensional space
func EuclideanDistance(pointOne, pointTwo types.Point2D) (dist float64) {
	deltaX := pointTwo.X - pointOne.X
	deltaY := pointOne.Y - pointOne.Y
	dist = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return dist
}

// MinkowskiDistance returns minkowski distance of points in an N dimensional space,
func MinkowskiDistance(PointA, PointB types.Point2D, p float64) (distMinkowski float64) {
	var deltaA = math.Pow(PointA.X-PointA.Y, p)
	var deltaB = math.Pow(PointB.X-PointB.Y, p)
	distMinkowski = math.Pow(deltaA+deltaB, 1/p)
	return distMinkowski
}

// HaversineDistance returns the physical space between to locations on earth
func HaversineDistance(PointA, PointB types.LatLong) (dist float64) {
	const radiusOfEarth = 6371 // in kilometers
	latitudeDiff := utils.Deg2Rad(PointA.Latitude - PointB.Latitude)
	longitudeDiff := utils.Deg2Rad(PointA.Longitude - PointB.Longitude)
	haversine := math.Pow(math.Sinh(latitudeDiff/2), 2) + math.Cos(PointA.Latitude)*math.Cos(PointA.Latitude)*math.Pow(math.Sin(longitudeDiff/2), 2)
	centralAngle := 2 * math.Atan2(math.Sqrt(haversine), math.Sqrt(1-haversine))
	dist = radiusOfEarth * centralAngle
	return dist
}
