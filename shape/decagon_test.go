package tests

import "testing"

func TestDecagon(t *testing.T) {
	decagon := Decagon{}
	// area
	observedAreaOfDecagon := decagon.Area()
	expectedAreaOfDecagon := 40.0
	if observedAreaOfDecagon != expectedAreaOfDecagon {
		t.Errorf("Expected ,%d, got %d", expectedAreaOfDecagon, observedAreaOfDecagon)
	}

	// perimeter
	observedPerimeterOfDecagon := decagon.Perimeter()
	expectedPerimeterOfDecagon := 40.0
	if observedAreaOfDecagon != expectedPerimeterOfDecagon {
		t.Errorf("Expected %d got %d", expectedAreaOfDecagon, observedPerimeterOfDecagon)
	}
}
