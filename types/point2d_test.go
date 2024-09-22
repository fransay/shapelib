package types_test

import (
	"shapelib/types"
	"testing"
)

// funct_tests point2D size
func TestSize(t *testing.T) {
	// point instance one
	var pointOne = types.Point2D{X: 300, Y: 450}
	PointOneSizeResult := pointOne.Size()
	PointOneSizeExpected := 0.0
	if PointOneSizeExpected != PointOneSizeResult {
		t.Errorf("Expected %f, got %f", PointOneSizeExpected, PointOneSizeResult)
	}

	// point instance two
	var pointTwo = types.Point2D{X: 100, Y: 200}
	PointTwoSizeResult := pointTwo.Size()
	PointTwoSizeExpected := 0.0
	if PointOneSizeExpected != PointOneSizeResult {
		t.Errorf("Expected %f, got %f", PointTwoSizeExpected, PointTwoSizeResult)
	}

	// point instance three
	var pointThree = types.Point2D{X: 120, Y: 650}
	PointThreeSizeResult := pointThree.Size()
	PointThreeSizeExpected := 0.0
	if PointOneSizeExpected != PointOneSizeResult {
		t.Errorf("Expected %f, got %f", PointThreeSizeExpected, PointThreeSizeResult)
	}
}

// funct_tests point2D coordinate array representation
func TestCoordinates(t *testing.T) {
	// point instance one
	var pointOne = types.Point2D{X: 2.0, Y: 5.0}
	pointOneExpectedX := 2.0
	pointOneExpectedY := 5.0
	if pointOneExpectedY != pointOne.Y && pointOne.X != pointOneExpectedX {
		t.Errorf("Inaccurate match")
	}

	// point instance two
	var pointTwo = types.Point2D{X: 10, Y: 45}
	pointTwoExpectedX := 10.0
	pointTwoExpectedY := 45.0
	if pointTwoExpectedY != pointTwo.Y && pointTwo.X != pointTwoExpectedX {
		t.Errorf("Inaccurate match")
	}

	// point instance three
	var pointThree = types.Point2D{X: 56.9, Y: 34.2}
	pointThreeExpectedX := 56.9
	pointThreeExpectedY := 34.2
	if pointThreeExpectedY != pointThree.Y && pointThree.X != pointThreeExpectedX {
		t.Errorf("Inaccurate match")
	}
}

// funct_tests cartesian to polar converter
func TestPolar(t *testing.T) {
	var pointPolar = types.Point2D{X: 56.9, Y: 34.2}
	pointPolarDistExpected := 23.4
	pointPolarBearObserved := 50.0
	pointPolarDistanceObserved, pointPolarBearingObserved := pointPolar.Polar()
	if pointPolarDistExpected != pointPolarDistanceObserved && pointPolarBearingObserved != pointPolarBearObserved {
		t.Errorf("Inaccurate match")
	}

}

// funct_tests translate2D
func TestTranslate2D(t *testing.T) {
	var pointOne = types.Point2D{X: 2.0, Y: 5.0}
	pointOneExpectedX := 2.0
	pointOneExpectedY := 5.0
	if pointOneExpectedY != pointOne.Y && pointOne.X != pointOneExpectedX {
		t.Errorf("Inaccurate match")
	}
}

// funct_tests rotate
func TestPointRotate(t *testing.T) {
	var pointOne = types.Point2D{X: 2.0, Y: 5.0}
	pointRotateExpected := types.Point2D{X: -0.8284271247461898, Y: 4.949747468305833}
	pointRotateObserved := pointOne.Rotate(45)
	if pointRotateExpected != pointRotateObserved {
		t.Errorf("Expected %f, got %f", pointRotateExpected, pointRotateObserved)
	}

}

// test scale
func TestPointScale(t *testing.T) {
	var pointOne = types.Point2D{X: 2.0, Y: 5.0}
	pointScaleExpected := types.Point2D{X: 4, Y: 20}
	pointScaleObserved := pointOne.Scale([]float64{2, 4})
	if pointScaleExpected != pointScaleObserved {
		t.Errorf("Expected %f, got %f", pointScaleExpected, pointScaleObserved)
	}
}

// test collinearity
func TestCollinear(t *testing.T) {
	observed := types.IsCollinear(types.Point2D{X: 3, Y: 5}, types.Point2D{X: 10, Y: 50}, types.Point2D{X: 30, Y: 70}, types.Point2D{X: 80, Y: 15})
	const expected = false
	if observed != expected {
		t.Errorf("Expected %t, Got %t", expected, observed)
	}
}
