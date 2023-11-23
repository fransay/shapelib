package shapelib

import (
	"shapelib/types/point"
	"testing"
)

// test point2D size
func TestSize(t *testing.T) {
	// point instance one
	var pointOne = point.Point2D{X: 300, Y: 450}
	PointOneSizeResult := pointOne.Size()
	PointOneSizeExpected := 0.0
	if PointOneSizeExpected != PointOneSizeResult {
		t.Errorf("Expected %f, got %f", PointOneSizeExpected, PointOneSizeResult)
	}

	// point instance two
	var pointTwo = point.Point2D{X: 100, Y: 200}
	PointTwoSizeResult := pointTwo.Size()
	PointTwoSizeExpected := 0.0
	if PointOneSizeExpected != PointOneSizeResult {
		t.Errorf("Expected %f, got %f", PointTwoSizeExpected, PointTwoSizeResult)
	}

	// point instance three
	var pointThree = point.Point2D{X: 120, Y: 650}
	PointThreeSizeResult := pointThree.Size()
	PointThreeSizeExpected := 0.0
	if PointOneSizeExpected != PointOneSizeResult {
		t.Errorf("Expected %f, got %f", PointThreeSizeExpected, PointThreeSizeResult)
	}
}

// test point2D coordinate array representation
func TestCoordinates(t *testing.T) {
	// point instance one
	var pointOne = point.Point2D{X: 2.0, Y: 5.0}
	pointOneExpectedX := 2.0
	pointOneExpectedY := 5.0
	if pointOneExpectedY != pointOne.Y && pointOne.X != pointOneExpectedX {
		t.Errorf("Inaccurate match")
	}

	// point instance two
	var pointTwo = point.Point2D{X: 10, Y: 45}
	pointTwoExpectedX := 10.0
	pointTwoExpectedY := 45.0
	if pointTwoExpectedY != pointTwo.Y && pointTwo.X != pointTwoExpectedX {
		t.Errorf("Inaccurate match")
	}

	// point instance three
	var pointThree = point.Point2D{X: 56.9, Y: 34.2}
	pointThreeExpectedX := 56.9
	pointThreeExpectedY := 34.2
	if pointThreeExpectedY != pointThree.Y && pointThree.X != pointThreeExpectedX {
		t.Errorf("Inaccurate match")
	}
}

// test cartesian to polar converter
func TestPolar(t *testing.T) {
	// TODO: populate with test
}

// test translate2D
func TestTranslate2D(t *testing.T) {
	// TODO: populate with test
}
