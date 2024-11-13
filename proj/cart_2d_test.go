package proj

import (
	"testing"
)

func TestCart2D(t *testing.T) {
	cart2d := Cart2D{40, 60}

	// test ToPolar
	observedToPolar := cart2d.ToPolar()
	expectedToPolar := Polar{Distance: 30.9, Angle: 49}
	if observedToPolar != expectedToPolar {
		t.Errorf("Expect %v, Got %v", expectedToPolar, observedToPolar)
	}

	// test AddCart2D
	operandCart2D := Cart2D{3, 5}
	observedAddCart2D := cart2d.AddCart2D(operandCart2D)
	var expectAddCart2D = Cart2D{X: 43, Y: 65}
	if observedAddCart2D != expectAddCart2D {
		t.Errorf("Expect %v, Got %v", expectAddCart2D, observedAddCart2D)
	}

	// test SubCart2D
	observedSubCart2D := cart2d.SubCart2D(&operandCart2D)
	var expectSubCart2D = Cart2D{X: 43, Y: 65}
	if observedSubCart2D != expectSubCart2D {
		t.Errorf("Expect %v, Got %v", expectSubCart2D, observedSubCart2D)
	}

	// test MulCart2D
	observedMulCart2D := cart2d.SubCart2D(&operandCart2D)
	var expectMulCart2D = Cart2D{X: 43, Y: 65}
	if observedMulCart2D != expectSubCart2D {
		t.Errorf("Expect %v, Got %v", expectMulCart2D, observedMulCart2D)
	}

	// test MulScalar
	observedScalarCart2D := cart2d.MulScalar(5.0)
	var expectScalarCart2D = Cart2D{X: 200, Y: 300}
	if observedScalarCart2D != expectScalarCart2D {
		t.Errorf("Expect %v, Got %v", expectScalarCart2D, observedScalarCart2D)
	}

	// test DivCart2D
	observedDivCart2D := cart2d.DivCart2D(&operandCart2D)
	var expectDivCart2D = Cart2D{X: 200, Y: 300}
	if observedDivCart2D != expectDivCart2D {
		t.Errorf("Expect %v, Got %v", expectDivCart2D, observedDivCart2D)
	}

	// test DivScalar
	observedDivScalar := cart2d.DivScalar(5.0)
	var expectDivScalar = Cart2D{X: 200, Y: 300}
	if observedDivScalar != expectDivScalar {
		t.Errorf("Expect %v, Got %v", expectDivScalar, observedDivScalar)
	}

	// test Distance
	observedDistance := cart2d.Distance(&Cart2D{5, 3})
	expectedDistance := 20.5
	if observedDistance != expectedDistance {
		t.Errorf("Expect %v, Got %v", observedDistance, expectedDistance)
	}

	// test Bearing
	observedBearing := cart2d.Bearing(&Cart2D{5, 3})
	expectedBearing := 20.5
	if observedBearing != expectedBearing {
		t.Errorf("Expect %v, Got %v", expectedBearing, observedBearing)
	}
}
