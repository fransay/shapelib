package cartesian_test

import (
	"reflect"
	"shapelib/coord-sys/cartesian"
	"shapelib/types"
	"testing"
)

func TestCart3D(t *testing.T) {
	cart3D := cartesian.Cartesian3D{
		X: cartesian.Axis{Start: 0, Step: 5, End: 15},
		Y: cartesian.Axis{Start: 5, Step: 5, End: 20},
		Z: cartesian.Axis{Start: 10, Step: 10, End: 30},
	}

	// 3D x values
	resultX3D := cart3D.XValues()
	expectedX3D := []float64{0, 5, 10, 15}
	if !reflect.DeepEqual(resultX3D, expectedX3D) {
		t.Errorf("Expected: %v, Got: %v", expectedX3D, resultX3D)
	}

	// 3D y values
	resultY3D := cart3D.YValues()
	expectedY3D := []float64{5, 10, 15, 20}
	if !reflect.DeepEqual(resultX3D, expectedX3D) {
		t.Errorf("Expected: %v, Got: %v", expectedY3D, resultY3D)
	}

	// 3D z values
	resultZ3D := cart3D.ZValues()
	expectedZ3D := []float64{10, 20, 30}
	if !reflect.DeepEqual(resultZ3D, expectedZ3D) {
		t.Errorf("Expected: %v, Got: %v", expectedY3D, resultY3D)
	}

	// test origin
	resultOrigin := cart3D.Origin()
	expectedOrigin := types.Point3D{X: 0, Y: 5, Z: 10}
	if !reflect.DeepEqual(resultOrigin, expectedOrigin) {
		t.Errorf("Expected %v, Got %v", expectedOrigin, resultOrigin)
	}
}
