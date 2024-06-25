package tests_test

import (
	"shapelib/types"
	"testing"
)

// funct_tests size
func TestPoint3DSize(t *testing.T) {
	var point3DSizeObj = types.Point3D{X: 2, Y: 3}
	var point3DSizeObserved = point3DSizeObj.Size()
	var point3DSizeExpected = 0.0
	if point3DSizeObserved != point3DSizeExpected {
		t.Errorf("Expected %f, Got %f", point3DSizeExpected, point3DSizeObserved)
	}

}

// funct_tests translate3D
func TestPoint3DTranslate3D(t *testing.T) {
	var point3DSizeObj = types.Point3D{X: 2, Y: 3, Z: 3}
	var point3DTranslate3DObserved = point3DSizeObj.Translate3D(types.Point3D{X: 2, Y: 4, Z: 5})
	var point3DTranslated3DExpected = types.Point3D{X: 4, Y: 7, Z: 8}
	if point3DTranslate3DObserved != point3DTranslated3DExpected {
		t.Errorf("Expected %f, Got %f", point3DTranslated3DExpected, point3DTranslate3DObserved)
	}

}

// funct_tests scale
func TestPoint3DScale(t *testing.T) {
	var point3DSizeObj = types.Point3D{X: 2, Y: 3, Z: 3}
	var point3DScale3DObserved = point3DSizeObj.Scale([]float64{1, 2, 3})
	var point3DScale3DExpected = types.Point3D{X: 2, Y: 6, Z: 9}
	if point3DScale3DObserved != point3DScale3DExpected {
		t.Errorf("Expected %f, Got %f", point3DScale3DExpected, point3DScale3DObserved)
	}

}
