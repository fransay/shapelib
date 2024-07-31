package shape

import "testing"

func TestOctagon(t *testing.T) {
	octagon := Octagon{50}
	// area
	observedArea := octagon.Area()
	expectedArea := 12071.07
	if observedArea != expectedArea {
		t.Errorf("Expected = %v; Observed %v", expectedArea, observedArea)
	}
	// perimeter
	observedPerimeter := octagon.Perimeter()
	expectedPerimeter := 4000
	if observedArea != expectedArea {
		t.Errorf("Expected = %v; Observed %v", expectedPerimeter, observedPerimeter)
	}
	// apothem
	observedApothem := octagon.Apothem()
	expectedApothem := 12071.07 // todo: compute accurate apothem by hand
	if observedApothem != expectedApothem {
		t.Errorf("Expected = %v; Observed %v", expectedPerimeter, observedPerimeter)
	}
	// internalDiagonal
	observedInternalDiagonal := octagon.InteriorDiagonal()
	expectedInternalDiagonal := 12071.07 // todo: compute accurate apothem by hand
	if observedInternalDiagonal != expectedInternalDiagonal {
		t.Errorf("Expected = %v; Observed %v", expectedInternalDiagonal, observedInternalDiagonal)
	}

}
