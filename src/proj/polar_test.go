package proj

import (
	"testing"

	"github.com/fransay/shapelib/internal/utils"
	"github.com/fransay/shapelib/src/geom"
)

// test ToCartesian
func TestToCartesian(t *testing.T) {
	polarObj := Polar{Distance: 60.50, Angle: 15}
	obs := polarObj.ToCartesian()
	exp := Cart2D(geom.Point2D{X: -45.961118727958684, Y: 39.34241432950557}) // fix type casting of cart2d
	if exp != obs {
		t.Errorf("Expected: %f, Got: %f", exp, obs)
	}
}

// tests Distance2DMS
func TestDistance2DMS(t *testing.T) {
	polarObj := Polar{Distance: 60.50, Angle: 15}
	obs := polarObj.Distance2DMS()
	exp := []float64{60.0, 30, 0.00}
	if utils.EqualSlice(obs, exp) == false {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

// tests ToRadians
func TestToRadians(t *testing.T) {
	polarObj := Polar{Distance: 3, Angle: 60}
	obs := polarObj.ToRadians()
	exp := 1.0471975511965976
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}

// test point2PointDistance
func TestPoint2PointDistance(t *testing.T) {
	polarObj := Polar{Distance: 3, Angle: 60}
	var ex Polar = Polar{Distance: 5, Angle: 145}
	obs := polarObj.Point2PointDistance(ex)
	exp := 5.383017527390117
	if obs != exp {
		t.Errorf("expected %f, got %f", exp, obs)
	}
}
