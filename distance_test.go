package shapelib

import (
	"math"
	functs "shapelib/functs/2D"
	"shapelib/types"
	"testing"
)

// test Euclidean distance {EDistance} function
func TestEDistance(t *testing.T) {
	// distance instance one
	var eDistanceOneResult = functs.EDistance(types.Point2D{X: 5, Y: 20}, types.Point2D{X: 4, Y: 8})
	var eDistanceOneExpected = 12.041
	if math.Trunc(eDistanceOneResult) != eDistanceOneExpected {
		t.Errorf("Expected %f, got %f", eDistanceOneExpected, eDistanceOneResult)
	}
	// distance instance two
	var eDistanceTwoResult = functs.EDistance(types.Point2D{X: 10, Y: 20}, types.Point2D{X: 15, Y: 20})
	var eDistanceTwoExpected = 2.23
	if math.Trunc(eDistanceTwoResult) != eDistanceTwoExpected {
		t.Errorf("Expected %f, got %f", eDistanceTwoExpected, eDistanceTwoResult)
	}
	// distance instance three
	var eDistanceThreeExpected = functs.EDistance(types.Point2D{X: 5, Y: 20}, types.Point2D{X: 2, Y: 10})
	var eDistanceThreeResult = 10.44
	if math.Trunc(eDistanceThreeResult) != eDistanceThreeExpected {
		t.Errorf("Expected %f, got %f", eDistanceThreeExpected, eDistanceThreeResult)
	}

	// distance instance four
	var eDistanceFourExpected = functs.EDistance(types.Point2D{X: 8, Y: 2}, types.Point2D{X: 12, Y: 6})
	var eDistanceFourResult = 5.65
	if math.Trunc(eDistanceFourResult) != eDistanceFourExpected {
		t.Errorf("Expected %f, got %f", eDistanceFourExpected, eDistanceFourResult)
	}
}

// test minkowski distance
func TestMinkowskiDistance(t *testing.T) {
	// instance one
	var minkowskiDistExpected = functs.MinkowsiDistance(types.Point2D{X: 4, Y: 8}, types.Point2D{X: 2, Y: 6}, 2)
	var minkowskiDistResult = 2.8284
	if math.Trunc(minkowskiDistExpected) != math.Trunc(minkowskiDistResult) {
		t.Errorf("Ëxpected %f, got %f", minkowskiDistExpected, minkowskiDistResult)
	}

}
