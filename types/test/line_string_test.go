package test

import (
	"shapelib/types"
	"shapelib/types/point"
	"testing"
)

// test line string number of line segments
func TestLineStringNumberOfLineSegment(t *testing.T) {
	var numberOfLineSegmentObj = types.LineString{point.Point2D{X: 1, Y: 2}, point.Point2D{X: 3, Y: 4}, point.Point2D{X: 5, Y: 6}}
	var numberOfLineSegmentObserved = numberOfLineSegmentObj.NumberOfLineSegments()
	var numberOfLineSegmentExpected = 2
	if numberOfLineSegmentExpected != numberOfLineSegmentObserved {
		t.Errorf("Expected %v, Observed%v", numberOfLineSegmentExpected, numberOfLineSegmentObserved)
	}

}

// test line string length
func TestLineStringLength(t *testing.T) {
	var lineString = types.LineString{point.Point2D{X: 40, Y: 6}, point.Point2D{X: 100, Y: 2}, point.Point2D{X: 3, Y: 5}}
	lineStringLengthObserved := lineString.Length()
	lineStringLengthExpected := 18.0
	if lineStringLengthObserved != lineStringLengthExpected {
		t.Errorf("Expected %v, Observed %v", lineStringLengthExpected, lineStringLengthObserved)
	}

}
