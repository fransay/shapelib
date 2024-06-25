package tests_test

import (
	s "shapelib/types"
	"testing"
)

// number of nodes funct_tests
func TestPolygonNumberOfNodes(t *testing.T) {
	var polygonOne = s.Polygon{s.Point2D{X: 7, Y: 9}, s.Point2D{X: 2, Y: 6}, s.Point2D{X: 4, Y: 8}}
	var numberOfNodesResult = polygonOne.NumberOfNodes()
	var numberOfNodesExpected = 3
	if numberOfNodesResult != numberOfNodesExpected {
		t.Errorf("Expected %d, Got %d", numberOfNodesExpected, numberOfNodesResult)
	}
}

// number of line segment funct_tests
func TestPolygonLineSegment(t *testing.T) {
	// polygon instance one
	var polygonOne = s.Polygon{s.Point2D{X: 7, Y: 9}, s.Point2D{X: 2, Y: 6}, s.Point2D{X: 4, Y: 8}}
	var numberOfLineSegmentResult = polygonOne.NumberOfLineSegments()
	var numberOfLineSegmentExpected = 3
	if numberOfLineSegmentExpected != numberOfLineSegmentResult {
		t.Errorf("Expected %d, Got %d", numberOfLineSegmentExpected, numberOfLineSegmentResult)
	}
	// polygon instance two
	var polygonTwo = s.Polygon{s.Point2D{X: 7, Y: 9}, s.Point2D{X: 2, Y: 6}, s.Point2D{X: 4, Y: 8}, s.Point2D{X: 4, Y: 50}}
	var numberOfLineSegmentTwoResult = polygonTwo.NumberOfLineSegments()
	var numberOfLineSegmentTwoExpected = 4
	if numberOfLineSegmentTwoExpected != numberOfLineSegmentTwoResult {
		t.Errorf("Expected %d, Got %d", numberOfLineSegmentTwoExpected, numberOfLineSegmentTwoResult)
	}
	// polygon instance three
	var polygonThree = s.Polygon{s.Point2D{X: 7, Y: 9}, s.Point2D{X: 2, Y: 6}, s.Point2D{X: 4, Y: 8}, s.Point2D{X: 12, Y: 43},
		s.Point2D{X: 45, Y: 645}}
	var numberOfLineSegmentThreeResult = polygonThree.NumberOfLineSegments()
	var numberOfLineSegmentThreeExpected = 5
	if numberOfLineSegmentThreeExpected != numberOfLineSegmentThreeResult {
		t.Errorf("Expected %d, Got %d", numberOfLineSegmentThreeExpected, numberOfLineSegmentThreeResult)
	}
}

// centroid funct_tests
func TestPolygonCentroid(t *testing.T) {
	// polygon instance one
	var polygonOne = s.Polygon{s.Point2D{X: 7, Y: 9}, s.Point2D{X: 2, Y: 6}, s.Point2D{X: 4, Y: 8}}
	var centroidResults = polygonOne.Centroid()
	var centroidExpected = s.Point2D{X: 6.5, Y: 11.5}
	if centroidExpected != centroidResults {
		t.Errorf("Expected %f, Got %f", centroidExpected, centroidResults)
	}
	// polygon instance two
	var polygonTwo = s.Polygon{s.Point2D{X: 1, Y: 5}, s.Point2D{X: 6, Y: 9}, s.Point2D{X: 5, Y: 10}}
	var centroidTwoResult = polygonTwo.Centroid()
	var centroidTwoExpected = s.Point2D{X: 6, Y: 12}
	if centroidTwoExpected != centroidTwoResult {
		t.Errorf("Expected %f, Got %f", centroidTwoExpected, centroidTwoResult)
	}

	// polygon instance three
	var polygonThree = s.Polygon{s.Point2D{X: 60, Y: 50}, s.Point2D{X: 22, Y: 65}, s.Point2D{X: 40, Y: 8}}
	var centroidThreeResults = polygonThree.Centroid()
	var centroidThreeExpected = s.Point2D{X: 61.0, Y: 61.5}
	if centroidThreeExpected != centroidThreeResults {
		t.Errorf("Expected %f, Got %f", centroidThreeExpected, centroidThreeResults)
	}
}

// funct_tests number of line segments in polygon
func TestNumberOfLineSegments(t *testing.T) {
	var polygonOne = s.Polygon{s.Point2D{X: 60, Y: 50}, s.Point2D{X: 22, Y: 65}, s.Point2D{X: 40, Y: 8}}
	var numberOfLineSegmentsObserved = polygonOne.NumberOfLineSegments()
	var numberOfLineSegmentsExpected = 2
	if numberOfLineSegmentsObserved != numberOfLineSegmentsExpected {
		t.Errorf("Expected %v, Got %v", numberOfLineSegmentsExpected, numberOfLineSegmentsObserved)

	}
}

// funct_tests number of nodes in a polygon
func TestNumberOfNodes(t *testing.T) {
	var polygonOne = s.Polygon{s.Point2D{X: 60, Y: 50}, s.Point2D{X: 22, Y: 65}, s.Point2D{X: 40, Y: 8}}
	var numberOfNodesObserved = polygonOne.NumberOfNodes()
	var numberOfNodesExpected = 3
	if numberOfNodesObserved != numberOfNodesExpected {
		t.Errorf("Expected %v, Got %v", numberOfNodesExpected, numberOfNodesObserved)

	}
}

// funct_tests rotate transformation
func TestRotate(t *testing.T) {
	var polygonOne = s.Polygon{s.Point2D{X: 60, Y: 50}, s.Point2D{X: 22, Y: 65}, s.Point2D{X: 40, Y: 8}}
	var polygonTwo = s.Polygon{s.Point2D{X: 130.71132228476674, Y: 110.83117391955632}, s.Point2D{X: 25.446469962218117, Y: 31.677564464354067}, s.Point2D{X: 446.7709136776561, Y: 302.49969654271126}}
	var RotatePolygonObserved = polygonOne.Rotate(40)
	var RotatePolygonExpected = polygonTwo
	if RotatePolygonExpected.IsEqual(RotatePolygonObserved) {
		t.Errorf("Expected %v, Got %v", RotatePolygonExpected, RotatePolygonObserved)

	}

}

// funct_tests scale transformation
func TestScale(t *testing.T) {
	var polygonOne = s.Polygon{s.Point2D{X: 60, Y: 50}, s.Point2D{X: 22, Y: 65}, s.Point2D{X: 40, Y: 8}}
	var scaledPolygonObserved = polygonOne.Scale([]float64{2.0, 2.0})
	var scaledPolygonExpected = s.Polygon{s.Point2D{X: 120, Y: 100}, s.Point2D{X: 44, Y: 130}, s.Point2D{X: 80, Y: 16}}
	if !scaledPolygonObserved.IsEqual(scaledPolygonExpected) {
		t.Errorf("Expected %v, Got %v", scaledPolygonExpected, scaledPolygonObserved)
	}
}

// funct_tests isequal method
func TestIsEqual(t *testing.T) {
	var polygonOne = s.Polygon{s.Point2D{X: 60, Y: 50}, s.Point2D{X: 22, Y: 65}, s.Point2D{X: 40, Y: 8}}
	var polygonTwo = s.Polygon{s.Point2D{X: 10, Y: 20}, s.Point2D{X: 5, Y: 15}, s.Point2D{X: 30, Y: 15}}
	var IsEqualPolygonObserved = polygonOne.IsEqual(polygonTwo)
	var IsEqualPolygonExpected = false
	if IsEqualPolygonObserved != IsEqualPolygonExpected {
		t.Errorf("Expected %v, Got %v", IsEqualPolygonExpected, IsEqualPolygonObserved)

	}

}

// funct_tests contains method
func TestContains(t *testing.T) {
	var polygonOne = s.Polygon{s.Point2D{X: 60, Y: 50}, s.Point2D{X: 22, Y: 65}, s.Point2D{X: 40, Y: 8}}
	var polygonTwo = s.Polygon{s.Point2D{X: 10, Y: 20}, s.Point2D{X: 5, Y: 15}, s.Point2D{X: 30, Y: 15}}
	var ContainsPolygonObserved = polygonOne.Contains(polygonTwo, s.Point2D{X: 10, Y: 20})
	var ContainsPolygonExpected = true
	if ContainsPolygonObserved != ContainsPolygonExpected {
		t.Errorf("Expected %v, Got %v", ContainsPolygonExpected, ContainsPolygonObserved)

	}

}
