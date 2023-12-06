package shapelib

import (
	cs "shapelib/coord-sys"
	ut "shapelib/utils"
	"testing"
)

func TestToCartesian(t *testing.T) {
	polarObj := cs.Polar{Distance: 60.50, Angle: 15}
	obs := polarObj.ToCartesian()
	exp := cs.Cart2D{X: -45.961118727958684, Y: 39.34241432950557} //TODO: do calculation by hand for second check
	if exp != obs {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

func TestDistance2DMS(t *testing.T) {
	polarObj := cs.Polar{Distance: 60.50, Angle: 15}
	obs := polarObj.Distance2DMS()
	exp := []float64{60.0, 30, 0.00}
	if ut.EqualSlice(obs, exp) == false {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

func TestToRadians(t *testing.T) {
	polarObj := cs.Polar{Distance: 3, Angle: 60}
	obs := polarObj.ToRadians()
	exp := 1.0471975511965976
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

func TestPoint2PointDistance(t *testing.T) {
	polarObj := cs.Polar{Distance: 3, Angle: 60}
	var ex cs.Polar = cs.Polar{Distance: 5, Angle: 145}
	obs := polarObj.Point2PointDistance(ex)
	exp := 5.383017527390117
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}
