package props

import (
	"math"
	"testing"

	"github.com/fransay/shapelib/src/geog"
	"github.com/fransay/shapelib/src/geom"
)

// Test euclidean distance
func TestEuclideanDistance(t *testing.T) {
	// object one
	var eDistanceOneResult = EuclideanDistance(geom.Point2D{X: 5, Y: 20}, geom.Point2D{X: 4, Y: 8})
	var eDistanceOneExpected = 1.0
	if math.Trunc(eDistanceOneResult) != eDistanceOneExpected {
		t.Errorf("Expected %f, Got %f", eDistanceOneExpected, eDistanceOneResult)
	}
}

// Test minkowski distance
func TestMinkowskiDistance(t *testing.T) {
	// object one
	var minkowskiDistExpected = MinkowskiDistance(geom.Point2D{X: 4, Y: 8}, geom.Point2D{X: 2, Y: 6}, 2)
	var minkowskiDistResult = 5.656854249492381
	if math.Trunc(minkowskiDistExpected) != math.Trunc(minkowskiDistResult) {
		t.Errorf("Ëxpected %f, Got %f", minkowskiDistExpected, minkowskiDistResult)
	}
}

// Test haversine distance
func TestHaversineDistance(t *testing.T) {
	stationA := geog.LatLong{Latitude: 20.0, Longitude: 60.0}
	stationB := geog.LatLong{Latitude: 40.0, Longitude: 100.0}
	haversineDistanceObs := HaversineDistance(stationA, stationB)
	haversineDistanceExp := 2880.8708298130377
	if haversineDistanceExp != haversineDistanceObs {
		t.Errorf("Ëxpected %f, Got %f", haversineDistanceExp, haversineDistanceObs)
	}
}

// Test chebyshev distance
func TestChebyshevDistance(t *testing.T) {
	firstPoint := geom.Point2D{X: 5, Y: 20}
	secondPoint := geom.Point2D{X: 4, Y: 8}
	var observedChebyshevDistance = ChebyshevDistance(firstPoint, secondPoint)
	const expectedChebyshevDistance = 1
	if observedChebyshevDistance != expectedChebyshevDistance {
		t.Errorf("Ëxpected %f, Got %f", expectedChebyshevDistance, observedChebyshevDistance)
	}
}

// other distances to explore:
// 1. hamming distance
// 2. frechet distance
