package test

import (
	"reflect"
	"shapelib/coord-sys/cart"
	_d "shapelib/coord-sys/cart/2d"
	"testing"
)

func TestCard2d(t *testing.T) {
	cart2dObject := _d.Cart2d{
		X: cart.Axis{Start: 0, Step: 4, End: 20},
		Y: cart.Axis{Start: 60, Step: 5, End: 70},
	}

	// test for x values
	resultX2d := cart2dObject.XValues()
	expectX2d := []float64{0, 4, 8, 12, 16, 20}
	if !reflect.DeepEqual(expectX2d, resultX2d) {
		t.Errorf("expect2d %v, got %v", expectX2d, resultX2d)
	}

	// test for y values
	resultY2d := cart2dObject.YValues()
	expectY2d := []float64{60, 65, 70}
	if !reflect.DeepEqual(expectY2d, resultY2d) {
		t.Errorf("expect2d %v, got %v", expectY2d, resultY2d)
	}
}
