package tests

import (
	"shapelib/shape"
	"shapelib/shape/interfaces"
	"testing"
)

func TestQuadrilateral(t *testing.T) {

	var a interfaces.Quadrilateral = shape.Rectangle{Length: 20, Width: 50}

	// The quadrilateral implement only
	if a.IsQuad() == false {
		t.Errorf("Expected: %v Got: %v", true, a.IsQuad())
	}

}
