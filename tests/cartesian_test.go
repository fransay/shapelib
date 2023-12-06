package types

import (
	cs "shapelib/coord-sys"
	"testing"
)

func TestToPolar(t *testing.T) {
	cartObj := cs.Cart2D{X: 20, Y: 5}
	obs := cartObj.ToPolar()
	exp := cs.Polar{Distance: 20.615528128088304, Angle: 1.3258176636680323}
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

func TestPoint2PointDist(t *testing.T) {
	cartObj := cs.Cart2D{X: 20, Y: 5}
	obs := cartObj.Point2PointDistance(cs.Cart2D{X: 10, Y: 5})
	exp := 10.0
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}
