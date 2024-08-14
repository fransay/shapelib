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

	// test AddCart3D
	observedAddCart3D := cart3d.AddCart3D(&Cart3D{X: 20, Y: 20, Z: 10})
	expectedAddCart3D := Cart3D{X: 70, Y: 40, Z: 20}
	if observedAddCart3D != expectedAddCart3D {
		t.Errorf("observedAddCart3D = %v; want %v", observedAddCart3D, expectedAddCart3D)
	}

	// test SubCart3D
	observedSubCart3D := cart3d.SubCart3D(&Cart3D{X: 20, Y: 20, Z: 10})
	expectedSubCart3D := Cart3D{X: 30, Y: 0, Z: 0}
	if observedSubCart3D != expectedSubCart3D {
		t.Errorf("Expect %v, Got %v", expectedSubCart3D, observedSubCart3D)
	}
}
