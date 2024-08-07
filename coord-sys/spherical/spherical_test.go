package spherical

import (
	"shapelib/coord-sys/cartesian"
	"testing"
)

// tests point to point distance spherical
func TestSphericalPoint2PointDistance(t *testing.T) {
	sphericalObject := Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	observedDistance := sphericalObject.Point2PointDistance(Spherical{RadialDistance: 30.0, PolarAngle: 40.0, AzimuthAngle: 50.0})
	expectedDistance := 9.011182193785475
	if observedDistance != expectedDistance {
		t.Errorf("Expected %f, Got %f", expectedDistance, observedDistance)
	}
}

// tests spherical to cartesian coordinate in two-dimension
func TestCartesian2D(t *testing.T) {
	sphericalObject := Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	observedDistance := sphericalObject.ToCartesianIn2D()
	expectedDistance := cartesian.Cartesian2D{X: cartesian.Axis{Start: 0, End: 10, Step: 2}, Y: cartesian.Axis{Start: 100, End: 500, Step: 50}}
	if observedDistance != expectedDistance {
		t.Errorf("Expected %f, Got %f", expectedDistance, observedDistance)
	}
}

// tests spherical to cartesian coordinate in three-dimension
func TestSphericalToCartesian3D(t *testing.T) {
	sphericalObject := Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	observedDistance := sphericalObject.ToCartesianIn3D()
	expectedDistance := cartesian.Cartesian2D{X: cartesian.Axis{}, Y: cartesian.Axis{}}
	if observedDistance != expectedDistance {
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
