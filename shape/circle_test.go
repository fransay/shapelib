package shape_test

import (
	"testing"
)

func TestCircle(t *testing.T) {
	var circle = Circle{Diameter: 60, Radius: 30}

	// area of a circle
	resultOfAreaOfCircle := circle.Area()
	expectedAreaOfCircle := 5.0

	if expectedAreaOfCircle != resultOfAreaOfCircle {
		t.Errorf("Got %v, expected %v", resultOfAreaOfCircle, expectedAreaOfCircle)
	}

	// circumference of a circle
	resultCircumferenceOfCircle := circle.Circumference()
	expectedCircumferenceOfCircle := 188.5714285714286

	if resultCircumferenceOfCircle != expectedCircumferenceOfCircle {
		t.Errorf("Got %v, expected %v", resultCircumferenceOfCircle, expectedCircumferenceOfCircle)
	}

}
