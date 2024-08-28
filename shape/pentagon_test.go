package shape_test

import (
	"shapelib/shape"
	"testing"
)

func TestPentagon(t *testing.T) {
	pentagon := shape.Pentagon{Side: 20}
	// area
	observedAreaPentagon := pentagon.Area()
	expectedAreaPentagon := 100.0
	if observedAreaPentagon != expectedAreaPentagon {
		t.Errorf("Expected %v Got %v", observedAreaPentagon, expectedAreaPentagon)
	}

	// perimeter
	observedPerimeterPentagon := pentagon.Area()
	expectedPerimeterPentagon := 100.0
	if observedPerimeterPentagon != expectedPerimeterPentagon {
		t.Errorf("Expected %v Got %v", observedPerimeterPentagon, expectedPerimeterPentagon)
	}

	// diagonal
	observedDiagonalPentagon := pentagon.Diagonal()
	expectedDiagonalPentagon := 100.0
	if observedDiagonalPentagon != expectedDiagonalPentagon {
		t.Errorf("Expected %v Got %v", observedDiagonalPentagon, expectedDiagonalPentagon)
	}

	// circumcircle
	observedCircumcirclePentagon := pentagon.Circumcircle()
	expectedCircumcirclePentagon := 100.0
	if observedCircumcirclePentagon != expectedCircumcirclePentagon {
		t.Errorf("Expected %v Got %v", observedCircumcirclePentagon, expectedCircumcirclePentagon)
	}

	// incircle
	observedIncirclePentagon := pentagon.Incircle()
	expectedIncirclePentagon := 100.0
	if observedIncirclePentagon != expectedIncirclePentagon {
		t.Errorf("Expected %v Got %v", observedIncirclePentagon, expectedIncirclePentagon)
	}

}
