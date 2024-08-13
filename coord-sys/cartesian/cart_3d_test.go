package cartesian

import "testing"

func TestCart3D(t *testing.T) {
	cart3d := Cart3D{X: 50, Y: 20, Z: 10}

	// test distance3D
	observedDistance3D := cart3d.Distance3D(&Cart3D{X: 20, Y: 20, Z: 10})
	expectedDistance3D := 50.09
	if observedDistance3D != expectedDistance3D {
		t.Errorf("observedDistance3D = %v; want %v", observedDistance3D, expectedDistance3D)
	}
}
