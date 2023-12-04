package shapelib

import (
	coordsys "shapelib/coord-sys"
	"testing"
)

// test point2pointdistance
func TestSphericalPoint2PointDistance(t *testing.T) {
	spherObj := coordsys.Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	obsDist := spherObj.Point2PointDistance(coordsys.Spherical{RadialDistance: 30.0, PolarAngle: 40.0, AzimuthAngle: 50.0})
	expDist := 20.0
	if obsDist != expDist {
		t.Errorf("Expected %f, Got %f", expDist, obsDist)
	}
}

func TestSphericalToCartesian(t *testing.T) {
	spherObj := coordsys.Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	obsDist := spherObj.ToCartesian()
	expDist := coordsys.Cart{X: 4, Y: 3} // incorrect expDist
	if obsDist != expDist {
		t.Errorf("Expected %f, Got %f", expDist, obsDist)
	}
}
