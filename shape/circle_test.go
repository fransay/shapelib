package shape_test

import (
	"shapelib/shape"
	"shapelib/types"
	"testing"
)

func TestCircle(t *testing.T) {
	var cir = shape.Circle{Radius: 60, Centroid: types.Point2D{X: 20, Y: 45}}

	// area of a circle
	resultOfAreaOfCircle := cir.Area()
	expectedAreaOfCircle := 5.0

	if expectedAreaOfCircle != resultOfAreaOfCircle {
		t.Errorf("Got %v, expected %v", resultOfAreaOfCircle, expectedAreaOfCircle)
	}

	// circumference of a circle
	resultCircumferenceOfCircle := cir.Area()
	expectedCircumferenceOfCircle := 188.5714285714286

	if resultCircumferenceOfCircle != expectedCircumferenceOfCircle {
		t.Errorf("Got %v, expected %v", resultCircumferenceOfCircle, expectedCircumferenceOfCircle)
	}

}
