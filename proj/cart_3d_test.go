package proj

import "testing"

func TestCart3D(t *testing.T) {
	cart3d := Cart3D{X: 50, Y: 20, Z: 10}

	// test Distance3D
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

	// test MulCart3D
	observedMulCart3D := cart3d.MulCart3D(&Cart3D{X: 2, Y: 3, Z: 4})
	expectedMulCart3D := Cart3D{X: 100, Y: 60, Z: 40}
	if observedMulCart3D != expectedMulCart3D {
		t.Errorf("Expect %v, Got %v", expectedMulCart3D, observedMulCart3D)
	}

	// test MulScalar3D
	observedMulScalar3D := cart3d.MulScalar3D(5.0)
	expectedMulScalar3D := Cart3D{X: 250.0, Y: 100.0, Z: 500.0}
	if observedMulScalar3D != expectedMulScalar3D {
		t.Errorf("Expect %v, Got %v", expectedMulScalar3D, observedMulScalar3D)
	}

	// test DivCart3D
	observedDivCart3D := cart3d.DivCart3D(&Cart3D{X: 2, Y: 2, Z: 2})
	expectedDivCart3D := Cart3D{X: 25, Y: 10, Z: 5}
	if observedDivCart3D != expectedDivCart3D {
		t.Errorf("Expect %v, Got %v", expectedDivCart3D, observedDivCart3D)
	}

	// test DivScalar
	observedDivScalar := cart3d.DivScalar(5.0)
	expectedDivScalar := Cart3D{X: 10.0, Y: 4.0, Z: 2.0}
	if observedDivScalar != expectedDivScalar {
		t.Errorf("Expect %v, Got %v", expectedDivScalar, observedDivScalar)
	}
}
