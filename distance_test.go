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
	var eDistanceTwo = functs.EDistance(types.Point2D{X: 15, Y: 20}, types.Point2D{X: 14, Y: 18})
	// distance instance one
	// distance instance one
	// distance instance one
}
