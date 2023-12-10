package types

import (
	s "shapelib/types"
	"shapelib/types/point"
	"testing"
)

// number of nodes test
func TestPolygonNumberOfNodes(t *testing.T) {
	var polygonOne = s.Polygon{point.Point2D{X: 7, Y: 9}, point.Point2D{X: 2, Y: 6}, point.Point2D{X: 4, Y: 8}}
	var numberOfNodesResult = polygonOne.NumberOfNodes()
	var numberOfNodesExpected = 3
	if numberOfNodesResult != numberOfNodesExpected {
		t.Errorf("Expected %d, Got %d", numberOfNodesExpected, numberOfNodesResult)
	}
}

// number of line segment test
func TestPolygonLineSegment(t *testing.T) {
	// polygon instance one
	var polygonOne = s.Polygon{point.Point2D{X: 7, Y: 9}, point.Point2D{X: 2, Y: 6}, point.Point2D{X: 4, Y: 8}}
	var numberOfLineSegmentResult = polygonOne.NumberOfLineSegments()
	var numberOfLineSegmentExpected = 3
	if numberOfLineSegmentExpected != numberOfLineSegmentResult {
		t.Errorf("Expected %d, Got %d", numberOfLineSegmentExpected, numberOfLineSegmentResult)
	}
	// polygon instance two
	var polygonTwo = s.Polygon{point.Point2D{X: 7, Y: 9}, point.Point2D{X: 2, Y: 6}, point.Point2D{X: 4, Y: 8}, point.Point2D{X: 4, Y: 50}}
	var numberOfLineSegmentTwoResult = polygonTwo.NumberOfLineSegments()
	var numberOfLineSegmentTwoExpected = 4
	if numberOfLineSegmentTwoExpected != numberOfLineSegmentTwoResult {
		t.Errorf("Expected %d, Got %d", numberOfLineSegmentTwoExpected, numberOfLineSegmentTwoResult)
	}
	// polygon instance three
	var polygonThree = s.Polygon{point.Point2D{X: 7, Y: 9}, point.Point2D{X: 2, Y: 6}, point.Point2D{X: 4, Y: 8}, point.Point2D{X: 12, Y: 43},
		point.Point2D{X: 45, Y: 645}}
	var numberOfLineSegmentThreeResult = polygonThree.NumberOfLineSegments()
	var numberOfLineSegmentThreeExpected = 5
	if numberOfLineSegmentThreeExpected != numberOfLineSegmentThreeResult {
		t.Errorf("Expected %d, Got %d", numberOfLineSegmentThreeExpected, numberOfLineSegmentThreeResult)
	}
}

// centroid test
func TestPolygonCentroid(t *testing.T) {
	// polygon instance one
	var polygonOne = s.Polygon{point.Point2D{X: 7, Y: 9}, point.Point2D{X: 2, Y: 6}, point.Point2D{X: 4, Y: 8}}
	var centroidResults = polygonOne.Centroid()
	var centroidExpected = point.Point2D{X: 6.5, Y: 11.5}
	if centroidExpected != centroidResults {
		t.Errorf("Expected %f, Got %f", centroidExpected, centroidResults)
	}
	// polygon instance two
	var polygonTwo = s.Polygon{point.Point2D{X: 1, Y: 5}, point.Point2D{X: 6, Y: 9}, point.Point2D{X: 5, Y: 10}}
	var centroidTwoResult = polygonTwo.Centroid()
	var centroidTwoExpected = point.Point2D{X: 6, Y: 12}
	if centroidTwoExpected != centroidTwoResult {
		t.Errorf("Expected %f, Got %f", centroidTwoExpected, centroidTwoResult)
	}

	// polygon instance three
	var polygonThree = s.Polygon{point.Point2D{X: 60, Y: 50}, point.Point2D{X: 22, Y: 65}, point.Point2D{X: 40, Y: 8}}
	var centroidThreeResults = polygonThree.Centroid()
	var centroidThreeExpected = point.Point2D{X: 61.0, Y: 61.5}
	if centroidThreeExpected != centroidThreeResults {
		t.Errorf("Expected %f, Got %f", centroidThreeExpected, centroidThreeResults)
	}
}
