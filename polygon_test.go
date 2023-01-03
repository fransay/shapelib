package shapelib

import (
	"shapelib/types"
	"testing"
)

// number of nodes test
func TestPolygonNumberOfNodes(t *testing.T) {
	var polygonOne = types.Polygon{types.Point2D{X: 7, Y: 9}, types.Point2D{X: 2, Y: 6}, types.Point2D{X: 4, Y: 8}}
	var numberOfNodesResult = polygonOne.NumberOfNodes()
	var numberOfNodesExpected = 3
	if numberOfNodesResult != numberOfNodesExpected {
		t.Errorf("Expected %d, got %d", numberOfNodesExpected, numberOfNodesResult)
	}
}

// number of line segment test
func TestPolygonLineSegment(t *testing.T) {
	var polygonOne = types.Polygon{types.Point2D{X: 7, Y: 9}, types.Point2D{X: 2, Y: 6}, types.Point2D{X: 4, Y: 8}}
	var numberOfLineSegmentResult = polygonOne.NumberOfLineSegments()
	var numberOfLineSegmentExpected = 3
	if numberOfLineSegmentExpected != numberOfLineSegmentResult {
		t.Errorf("Expected %d, got %d", numberOfLineSegmentExpected, numberOfLineSegmentResult)
	}
}

//
