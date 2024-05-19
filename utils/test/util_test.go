package test

import (
	"shapelib/utils"
	"testing"
)

// test IsClose
func TestIsClose(t *testing.T) {
	res := utils.IsClose(10.00, 45.00, 0.001)
	if res != false {
		t.Errorf("Got: %v, Expected: %v", res, false)
	}
}

// test Diff
func TestDiff(t *testing.T) {
	res := utils.Diff(10.0, 15.0)
	if res != 5.0 {
		t.Errorf("Got: %v, Expected: %v", res, false)
	}
}

// test EqualSlice
func TestEqualSlice(t *testing.T) {
	var sliceOne = []float64{1, 2, 3, 4}
	var sliceTwo = []float64{1, 2, 3, 4}
	res := utils.EqualSlice(sliceOne, sliceTwo)
	if res != true {
		t.Errorf("Got %v, Expected %v", res, true)
	}
}

// test Rad2Deg
func TestRad2Deg(t *testing.T) {
	res := utils.Rad2Deg(40)
	exp := 2291.831180523293
	if res != exp {
		t.Errorf("Got %v Exp %v", res, exp)
	}

}

// test Deg2Rad
func TestDeg2Rad(t *testing.T) {
	res := utils.Deg2Rad(40)
	exp := 0.6981317007977318
	if res != exp {
		t.Errorf("Got %v Expected %v", res, exp)
	}
}

// test Distance
func TestDistance(t *testing.T) {
	res := utils.Distance([]float64{4, 6}, []float64{8, 2})
	exp := 2.8284271247461903
	if res != exp {
		t.Errorf("Got %v, Expect %v", res, exp)
	}
}
