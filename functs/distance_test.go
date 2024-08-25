package functs

import (
	"math"
	"shapelib/types"
	"testing"
)

// funct_tests Euclidean distance {EuclideanDistance} function
func TestEDistance(t *testing.T) {
	// object one
	var eDistanceOneResult = EuclideanDistance(types.Point2D{X: 5, Y: 20}, types.Point2D{X: 4, Y: 8})
	var eDistanceOneExpected = 1.0
	if math.Trunc(eDistanceOneResult) != eDistanceOneExpected {
		t.Errorf("Expected %f, Got %f", eDistanceOneExpected, eDistanceOneResult)
	}
}

// funct_tests minkowski distance
func TestMinkowskiDistance(t *testing.T) {
	// object one
	var minkowskiDistExpected = MinkowskiDistance(types.Point2D{X: 4, Y: 8}, types.Point2D{X: 2, Y: 6}, 2)
	var minkowskiDistResult = 5.656854249492381
	if math.Trunc(minkowskiDistExpected) != math.Trunc(minkowskiDistResult) {
		t.Errorf("Ëxpected %f, Got %f", minkowskiDistExpected, minkowskiDistResult)
	}
}

// funct_tests haversine distance
func TestHaversineDistance(t *testing.T) {
	stationA := types.LatLong{Latitude: 20.0, Longitude: 60.0}
	stationB := types.LatLong{Latitude: 40.0, Longitude: 100.0}
	haversineDistanceObs := HaversineDistance(stationA, stationB)
	haversineDistanceExp := 2880.8708298130377
	if haversineDistanceExp != haversineDistanceObs {
		t.Errorf("Ëxpected %f, Got %f", haversineDistanceExp, haversineDistanceObs)
	}
}

func TestChebyshevDistance(t *testing.T) {
	firstPoint := types.Point2D{X: 5, Y: 20}
	secondPoint := types.Point2D{X: 4, Y: 8}
	var observedChebyshevDistance = ChebyshevDistance(firstPoint, secondPoint)
	const expectedChebyshevDistance = 1
	if observedChebyshevDistance != expectedChebyshevDistance {
		t.Errorf("Ëxpected %f, Got %f", expectedChebyshevDistance, observedChebyshevDistance)
	}
}
