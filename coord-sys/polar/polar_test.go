package polar

import (
	"shapelib/types"
	ut "shapelib/utils"
	"testing"
)

// funct_tests to cartesian
func TestToCartesian(t *testing.T) {
	polarObj := Polar{Distance: 60.50, Angle: 15}
	obs := polarObj.ToCartesian()
	exp := types.Point2D{X: -45.961118727958684, Y: 39.34241432950557}
	if exp != obs {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

// funct_tests distance to degree minutes second representation
func TestDistance2DMS(t *testing.T) {
	polarObj := Polar{Distance: 60.50, Angle: 15}
	obs := polarObj.Distance2DMS()
	exp := []float64{60.0, 30, 0.00}
	if ut.EqualSlice(obs, exp) == false {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

// funct_tests toRadians
func TestToRadians(t *testing.T) {
	polarObj := Polar{Distance: 3, Angle: 60}
	obs := polarObj.ToRadians()
	exp := 1.0471975511965976
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

// funct_tests point to point distance
func TestPoint2PointDistance(t *testing.T) {
	polarObj := Polar{Distance: 3, Angle: 60}
	var ex Polar = Polar{Distance: 5, Angle: 145}
	obs := polarObj.Point2PointDistance(ex)
	exp := 5.383017527390117
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}
