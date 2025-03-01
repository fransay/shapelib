package shape

import (

	"testing"
)

func TestQuadrilateral(t *testing.T) {

	var a Quadrilateral = Rectangle{Length: 20, Width: 50}

	// The quadrilateral implement only
	if a.IsQuad() == false {
		t.Errorf("Expected: %v Got: %v", true, a.IsQuad())
	}

}
