package cart_test

import (
	"reflect"
	"shapelib/coord-sys/cart"
	"testing"
)

func TestCart3d(t *testing.T) {
	cart3D := cart.Cart3d{
		X: cart.Axis{Start: 0, Step: 5, End: 15},
		Y: cart.Axis{Start: 5, Step: 5, End: 20},
		Z: cart.Axis{Start: 10, Step: 10, End: 30},
	}

	// 3d x_values
	resultX3d := cart3D.XValues()
	expectedX3d := []float64{0, 5, 10, 15}
	if !reflect.DeepEqual(resultX3d, expectedX3d) {
		t.Errorf("expected: %v, got: %v", expectedX3d, resultX3d)
	}

	// 3d y_values
	resultY3d := cart3D.YValues()
	expectedY3d := []float64{5, 10, 15, 20}
	if !reflect.DeepEqual(resultX3d, expectedX3d) {
		t.Errorf("expected: %v, got: %v", expectedY3d, resultY3d)
	}

	// 3d z_values
	resultZ3d := cart3D.ZValues()
	expectedZ3d := []float64{10, 20, 30}
	if !reflect.DeepEqual(resultZ3d, expectedZ3d) {
		t.Errorf("expected: %v, got: %v", expectedY3d, resultY3d)
	}
}
