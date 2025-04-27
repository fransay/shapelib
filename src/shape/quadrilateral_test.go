package shape

import (
	"testing"
)

func TestQuadrilateral(t *testing.T) {

	var rect Quadrilateral = Rectangle{Length: 20, Width: 50}
	var square Quadrilateral = Square{Length: 20.0}

	// The quadrilateral implement only
	if rect.IsQuad() == false {
		t.Errorf("Expected: %v Got: %v", true, rect.IsQuad())
	}

}
