package cart_test

import (
	"reflect"
	"shapelib/coord-sys/cart"
	"testing"
)

func TestCart2d(t *testing.T) {
	cart2dObject := cart.Cartesian2D{
		X: cart.Axis{Start: 0, Step: 4, End: 20},
		Y: cart.Axis{Start: 60, Step: 5, End: 70},
	}

	// test for x values
	resultX2d := cart2dObject.XValues()
	expectX2d := []float64{0, 4, 8, 12, 16, 20}
	if !reflect.DeepEqual(expectX2d, resultX2d) {
		t.Errorf("Expect2d %v, Got %v", expectX2d, resultX2d)
	}

	// test for y values
	resultY2d := cart2dObject.YValues()
	expectY2d := []float64{60, 65, 70}
	if !reflect.DeepEqual(expectY2d, resultY2d) {
		t.Errorf("Expect2d %v, Got %v", expectY2d, resultY2d)
	}
}
