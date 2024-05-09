package types

import (
	"shapelib/shape"
	"testing"
)

func TestCircle(t *testing.T) {
	var circle = shape.Circle{60, 30}

	// area of a circle
	resultOfAreaOfCircle := circle.Area()
	expectedAreaOfCircle := 5

	if resultOfAreaOfCircle != expectedAreaOfCircle {

	}

}
