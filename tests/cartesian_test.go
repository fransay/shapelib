package types

import (
	cs "shapelib/coord-sys"
	"testing"
)

// test topolar
func TestToPolar(t *testing.T) {
	cartObj := cs.Cart2D{X: 20, Y: 5}
	obs := cartObj.ToPolar()
	exp := cs.Polar{Distance: 20.615528128088304, Angle: 1.3258176636680323}
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

// test point2pointdist
func TestPoint2PointDist(t *testing.T) {
	cartObj := cs.Cart2D{X: 20, Y: 5}
	obs := cartObj.Point2PointDistance2D(cs.Cart2D{X: 10, Y: 5})
	exp := 10.0
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

// test translate2D
func TestCartTranslate2D(t *testing.T) {
	cartObj := cs.Cart2D{X: 20, Y: 5}
	obs := cartObj.Translate2D([2]float64{1, 3})
	exp := cs.Cart2D{X: 21.0, Y: 8.0}
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

// test translate3D
func TestCartTranslate3D(t *testing.T) {
	cartObj := cs.Cart3D{X: 20, Y: 5, Z: 5}
	obs := cartObj.Translate3D([3]float64{1, 3, 30})
	exp := cs.Cart3D{X: 21.0, Y: 8.0, Z: 35.0}
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}

}

// test rotate2D
func TestCartRotate2D(t *testing.T) {
	cartObj := cs.Cart2D{X: 20, Y: 5}
	obs := cartObj.Rotate2D(30.0)
	exp := cs.Cart2D{X: 8.025187118215989, Y: -20.531889731295156}
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}

}

// test rotate3D
func TestCartRotate3D(t *testing.T) {
	cartObj := cs.Cart3D{X: 20, Y: 5, Z: 5}
	obs := cartObj.Rotate3D(20, 60)
	exp := cs.Cart3D{X: -2.5865880676648887, Y: -6.293551070481977, Z: 17.90637534758312}
	if obs != exp {
		t.Errorf("Expected %f, Got %f", exp, obs)
	}

}
