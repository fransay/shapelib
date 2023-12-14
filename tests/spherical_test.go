package types

import (
	coordsys "shapelib/coord-sys"
	"testing"
)

// test point to point distance spherical
func TestSphericalPoint2PointDistance(t *testing.T) {
	spherObj := coordsys.Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	obsDist := spherObj.Point2PointDistance(coordsys.Spherical{RadialDistance: 30.0, PolarAngle: 40.0, AzimuthAngle: 50.0})
	expDist := 9.011182193785475
	if obsDist != expDist {
		t.Errorf("Expected %f, Got %f", expDist, obsDist)
	}
}

// test spherical to cartesian coordinate in two-dim
func TestSphericalToCartesian2D(t *testing.T) {
	spherObj := coordsys.Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	obsDist := spherObj.ToCartesianIn2D()
	expDist := coordsys.Cart2D{X: -9.534169523255013, Y: 2.5923465282621976} // incorrect expDist
	if obsDist != expDist {
		t.Errorf("Expected %f, Got %f", expDist, obsDist)
	}
}

// test spherical to cartesian coordinate in three-dim
func TestSphericalToCartesian3D(t *testing.T) {
	spherObj := coordsys.Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	obsDist := spherObj.ToCartesianIn3D()
	expDist := coordsys.Cart3D{X: -9.534169523255013, Y: 2.5923465282621976, Z: 9.649660284921133}
	if obsDist != expDist {
		t.Errorf("Expected %f, Got %f", expDist, obsDist)
	}
}

// test toArray
func TestToArray(t *testing.T) {
	spherObj := coordsys.Spherical{RadialDistance: 10.0, PolarAngle: 50.0, AzimuthAngle: 30.0}
	obsDist := spherObj.ToArray()
	expDist := [3]float64{10.0, 50.0, 30.0}
	if obsDist != expDist {
		t.Errorf("Expected %f, Got %f", expDist, obsDist)
	}
}
