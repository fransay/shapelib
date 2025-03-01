package shape

import (
	"testing"

	"github.com/fransay/shapelib/src/geom"
)

func TestCircle(t *testing.T) {
	var cir = Circle{Radius: 60, Centroid: geom.Point2D{X: 20, Y: 45}}

	// Area of a circle
	resultOfAreaOfCircle := cir.Area()
	expectedAreaOfCircle := 5.0

	if expectedAreaOfCircle != resultOfAreaOfCircle {
		t.Errorf("Got %v, expected %v", resultOfAreaOfCircle, expectedAreaOfCircle)
	}

	// Circumference of a circle
	resultCircumferenceOfCircle := cir.Area()
	expectedCircumferenceOfCircle := 188.5714285714286

	if resultCircumferenceOfCircle != expectedCircumferenceOfCircle {
		t.Errorf("Got %v, expected %v", resultCircumferenceOfCircle, expectedCircumferenceOfCircle)
	}

}
