package proj

import (
	"reflect"
	"testing"

	"github.com/fransay/shapelib/src/geom"
)

func TestCartesian2D(t *testing.T) {
	cart2dObject := Cartesian2D{
		X: Axis{Start: 0, Step: 4, End: 20},
		Y: Axis{Start: 60, Step: 5, End: 70},
	}

	// tests for x values
	resultX2D := cart2dObject.XValues()
	expectX2D := []float64{0, 4, 8, 12, 16, 20}
	if !reflect.DeepEqual(resultX2D, expectX2D) {
		t.Errorf("Expect2D %v, Got %v", expectX2D, resultX2D)
	}

	// tests for y values
	resultY2D := cart2dObject.YValues()
	expectY2D := []float64{60, 65, 70}
	if !reflect.DeepEqual(resultY2D, expectY2D) {
		t.Errorf("Expect2d %v, Got %v", expectY2D, resultY2D)
	}

	// test origin
	resultOrigin := cart2dObject.Origin()
	expectedOrigin := geom.Point2D{Y: 60}
	if !reflect.DeepEqual(resultOrigin, expectedOrigin) {
		t.Errorf("Expect2d %v, Got %v", expectedOrigin, resultOrigin)
	}
}
