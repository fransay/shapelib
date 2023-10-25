package shapelib

import (
	"math"
	"shapelib/functs"
	"shapelib/types/point"
	"testing"
)

// test Euclidean distance {EuclideanDistance} function
func TestEDistance(t *testing.T) {
	// distance instance one
	var eDistanceOneResult = functs.EuclideanDistance(point.Point2D{X: 5, Y: 20}, point.Point2D{X: 4, Y: 8})
	var eDistanceOneExpected = 12.041
	if math.Trunc(eDistanceOneResult) != eDistanceOneExpected {
		t.Errorf("Expected %f, got %f", eDistanceOneExpected, eDistanceOneResult)
	}
	// distance instance two
	var eDistanceTwoResult = functs.EuclideanDistance(point.Point2D{X: 10, Y: 20}, point.Point2D{X: 15, Y: 20})
	var eDistanceTwoExpected = 2.23
	if math.Trunc(eDistanceTwoResult) != eDistanceTwoExpected {
		t.Errorf("Expected %f, got %f", eDistanceTwoExpected, eDistanceTwoResult)
	}
	// distance instance three
	var eDistanceThreeExpected = functs.EuclideanDistance(point.Point2D{X: 5, Y: 20}, point.Point2D{X: 2, Y: 10})
	var eDistanceThreeResult = 10.44
	if math.Trunc(eDistanceThreeResult) != eDistanceThreeExpected {
		t.Errorf("Expected %f, got %f", eDistanceThreeExpected, eDistanceThreeResult)
	}

	// distance instance four
	var eDistanceFourExpected = functs.EuclideanDistance(point.Point2D{X: 8, Y: 2}, point.Point2D{X: 12, Y: 6})
	var eDistanceFourResult = 5.65
	if math.Trunc(eDistanceFourResult) != eDistanceFourExpected {
		t.Errorf("Expected %f, got %f", eDistanceFourExpected, eDistanceFourResult)
	}
}

// test minkowski distance
func TestMinkowskiDistance(t *testing.T) {
	// instance one
	var minkowskiDistExpected = functs.MinkowskiDistance(point.Point2D{X: 4, Y: 8}, point.Point2D{X: 2, Y: 6}, 2)
	var minkowskiDistResult = 2.8284
	if math.Trunc(minkowskiDistExpected) != math.Trunc(minkowskiDistResult) {
		t.Errorf("Ëxpected %f, got %f", minkowskiDistExpected, minkowskiDistResult)
	}
	// TODO include a few couple test instances

}
