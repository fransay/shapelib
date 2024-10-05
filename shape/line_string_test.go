package tests

import (
	"testing"
)

// funct_tests line string number of line segments
func TestLineStringNumberOfLineSegment(t *testing.T) {
	var numberOfLineSegmentObj = LineString{Point2D{X: 1, Y: 2}, Point2D{X: 3, Y: 4}, Point2D{X: 5, Y: 6}}
	var numberOfLineSegmentObserved = numberOfLineSegmentObj.NumberOfLineSegments()
	var numberOfLineSegmentExpected = 2
	if numberOfLineSegmentExpected != numberOfLineSegmentObserved {
		t.Errorf("Expected %v, Observed%v", numberOfLineSegmentExpected, numberOfLineSegmentObserved)
	}

}

// funct_tests line string length
func TestLineStringLength(t *testing.T) {
	var lineString = LineString{Point2D{X: 40, Y: 6}, Point2D{X: 100, Y: 2}, Point2D{X: 3, Y: 5}}
	lineStringLengthObserved := lineString.Length()
	lineStringLengthExpected := 18.0
	if lineStringLengthObserved != lineStringLengthExpected {
		t.Errorf("Expected %v, Observed %v", lineStringLengthExpected, lineStringLengthObserved)
	}

}
