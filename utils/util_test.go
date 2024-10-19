package utils

import (
	"testing"
)

// Test IsClose
func TestIsClose(t *testing.T) {
	res := IsClose(10.00, 45.00, 0.001)
	if res != false {
		t.Errorf("Got: %v, Expected: %v", res, false)
	}

	res2 := IsClose(10.00, 10.00, 0.001)
	if res2 != true {
		t.Errorf("Got: %v, Expected: %v", res2, true)
	}
}

// Test Diff
func TestDiff(t *testing.T) {
	res := Diff(10.0, 15.0)
	if res != 5.0 {
		t.Errorf("Got: %v, Expected: %v", res, false)
	}
}

// Test EqualSlice
func TestEqualSlice(t *testing.T) {
	var sliceOne = []float64{1, 2, 3, 4}
	var sliceTwo = []float64{1, 2, 3, 4}
	res := EqualSlice(sliceOne, sliceTwo)
	if res != true {
		t.Errorf("Got %v, Expected %v", res, true)
	}
}

// Test Rad2Deg
func TestRad2Deg(t *testing.T) {
	res := Rad2Deg(40)
	exp := 2291.831180523293
	if res != exp {
		t.Errorf("Got %v Exp %v", res, exp)
	}

}

// Test Deg2Rad
func TestDeg2Rad(t *testing.T) {
	res := Deg2Rad(40)
	exp := 0.6981317007977318
	if res != exp {
		t.Errorf("Got %v Expected %v", res, exp)
	}
}

// Test Distance
func TestDistance(t *testing.T) {
	res := Distance([]float64{4, 6}, []float64{8, 2})
	exp := 2.8284271247461903
	if res != exp {
		t.Errorf("Got %v, Expect %v", res, exp)
	}
}

// Test CompareThreeFloat64
func TestCompareTrio(t *testing.T) {
	res := CompareThreeFloat64(1, 2, 3)
	if res != false {
		t.Errorf("Got %v, Expected %v", res, false)
	}
}

// Test CompareTwoFloat64
func TestCompareDuo(t *testing.T) {
	res := CompareTwoFloat64(1, 1)
	if res != true {
		t.Errorf("Got %v, Expected %v", res, true)
	}
}

// Test IsArrayElementFloat64
func TestIsArrayElementFloat64(t *testing.T) {
	arr := []float64{1, 4, 5}
	observed := IsArrayElementFloat64(arr)
	if observed != true {
		t.Errorf("Got %v, Expected %v", observed, true)
	}
}
