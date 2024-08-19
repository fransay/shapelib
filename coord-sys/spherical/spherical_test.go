package spherical

import (
	"shapelib/coord-sys/cartesian"
	"testing"
)

// tests point to point distance spherical
func TestSphericalPoint2PointDistance(t *testing.T) {
	sphericalObject := Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	observedDistance := sphericalObject.Distance(Spherical{RadialDistance: 30.0, PolarAngle: 40.0, AzimuthAngle: 50.0})
	expectedDistance := 9.011182193785475
	if observedDistance != expectedDistance {
		t.Errorf("Expected %f, Got %f", expectedDistance, observedDistance)
	}
}

// tests spherical to cartesian coordinate in two-dimension
func TestCart2D(t *testing.T) {
	sphericalObject := Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	observedDistance := sphericalObject.ToCart2D()
	expectedDistance := cartesian.Cart2D{X: 675.65, Y: 43.54}
	if observedDistance != expectedDistance {
		t.Errorf("Expected %f, Got %f", expectedDistance, observedDistance)
	}
}

// tests spherical to cartesian coordinate in three-dimension
func TestSphericalToCartesian3D(t *testing.T) {
	sphericalObject := Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	observedDistance := sphericalObject.ToCart3D()
	expectedDistance := cartesian.Cart3D{X: 343.5, Y: 433.43, Z: 344.3}
	if expectedDistance != observedDistance {
		t.Errorf("Expected %f, Got %f", expectedDistance, observedDistance)
	}
}

// tests toArray
func TestToArray(t *testing.T) {
	sphericalObject := Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	observedDistance := sphericalObject.ToArray()
	expectedDistance := [3]float64{10.0, 50.0, 30.0}
	if observedDistance != expectedDistance {
		t.Errorf("Expected %f, Got %f", expectedDistance, observedDistance)
	}
}
