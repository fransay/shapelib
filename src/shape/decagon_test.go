package shape

import "testing"

func TestDecagon(t *testing.T) {
	decagon := Decagon{}
	// area
	observedAreaOfDecagon := decagon.Area()
	expectedAreaOfDecagon := 40.0
	if observedAreaOfDecagon != expectedAreaOfDecagon {
		t.Errorf("Expected ,%v, got %v", expectedAreaOfDecagon, observedAreaOfDecagon)
	}

	// perimeter
	observedPerimeterOfDecagon := decagon.Perimeter()
	expectedPerimeterOfDecagon := 40.0
	if observedAreaOfDecagon != expectedPerimeterOfDecagon {
		t.Errorf("Expected %v got %v", expectedAreaOfDecagon, observedPerimeterOfDecagon)
	}
}
