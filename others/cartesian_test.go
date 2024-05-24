package others

import (
	"math"
	"shapelib/coord-sys/polar"
	"testing"
)

// test ToPolar
func TestToPolar(t *testing.T) {
	// obj 1
	cartObj := Cart2D{X: 20, Y: 5}
	obsRR, obsRD := cartObj.ToPolar()
	expRR, expRD := polar.Polar{
		Distance: 20.615528128088304, Angle: 0.24497866312686414},
		polar.Polar{
			Distance: 20.615528128088304, Angle: 14.036243467926479}
	if obsRR != expRR && obsRD != expRD {
		t.Errorf("Expected %f, Got %f", expRR, obsRR)
		t.Errorf("Expected %f, Got %f", expRD, obsRD)
	}

	// obj 2
	cartObj2 := Cart2D{X: 0, Y: 1}
	obsRR2, obsRD2 := cartObj2.ToPolar()
	expRR2, expRD2 := polar.Polar{
		Distance: 1.0, Angle: 1.5707963267948966},
		polar.Polar{Distance: 1.0, Angle: 90}
	if obsRR2 != expRR2 && obsRD2 != expRD2 {
		t.Errorf("Expected %f, Got %f", expRR2, obsRR2)
		t.Errorf("Expected %f, Got %f", expRD2, obsRD2)
	}
}

// test point2pointDist
func TestPoint2PointDist(t *testing.T) {
	// obj1
	cartObj := Cart2D{X: 20, Y: 5}
	obs := cartObj.Point2PointDistance2D(Cart2D{X: 10, Y: 5})
	exp := 10.0
	if obs != exp {
		t.Errorf("Expected %f, Got %f", exp, obs)
	}

	// obj2
	cartObj2 := Cart2D{X: 20, Y: 5}
	obs2 := cartObj2.Point2PointDistance2D(Cart2D{X: 50, Y: 10})
	exp2 := math.Sqrt(925)
	if obs != exp {
		t.Errorf("Expected %f, Got %f", exp2, obs2)
	}
}

// test translate2D
func TestCartTranslate2D(t *testing.T) {
	cartObj := Cart2D{X: 20, Y: 5}
	obs := cartObj.Translate2D([2]float64{1, 3})
	exp := Cart2D{X: 21.0, Y: 8.0}
	if obs != exp {
		t.Errorf("Expected %f, Got %f", exp, obs)
	}
}

// test translate3D
func TestCartTranslate3D(t *testing.T) {
	cartObj := Cart3D{X: 20, Y: 5, Z: 5}
	obs := cartObj.Translate3D([3]float64{1, 3, 30})
	exp := Cart3D{X: 21.0, Y: 8.0, Z: 35.0}
	if obs != exp {
		t.Errorf("Expected %f, Got %f", exp, obs)
	}
}

// test rotate2D
func TestCartRotate2D(t *testing.T) {
	cartObj := Cart2D{X: 20, Y: 5}
	obs := cartObj.Rotate2D(30.0)
	exp := Cart2D{X: 8.025187118215989, Y: -20.531889731295156}
	if obs != exp {
		t.Errorf("Expected %f, Got %f", exp, obs)
	}
}

// test rotate3D
func TestCartRotate3D(t *testing.T) {
	cartObj := Cart3D{X: 20, Y: 5, Z: 5}
	obs := cartObj.Rotate3D(20, 60)
	exp := Cart3D{
		X: -2.5865880676648887, Y: -6.293551070481977, Z: 17.90637534758312}
	if obs != exp {
		t.Errorf("Expected %f, Got %f", exp, obs)
	}
}
