package utils

import (
	"testing"
)

func TestIsClose(t *testing.T) {
	res := IsClose(10.00, 45.00, 0.001)
	if res != false {
		t.Errorf("Got: %v, Expected: %v", res, false)
	}
}

func TestDiff(t *testing.T) {
	res := Diff(10.0, 15.0)
	if res != 5.0 {
		t.Errorf("Got: %v, Expected: %v", res, false)
	}
}

func TestEqualSlice(t *testing.T) {
	var sliceOne = []float64{1, 2, 3, 4}
	var sliceTwo = []float64{1, 2, 3, 4}
	res := EqualSlice(sliceOne, sliceTwo)
	if res != true {
		t.Errorf("Got %v, Expected %v", res, true)
	}
}

func TestRad2Deg(t *testing.T) {
	res := Rad2Deg(40)
	exp := 2291.831180523293
	if res != exp {
		t.Errorf("Got %v Exp %v", res, exp)
	}

}

func TestDeg2Rad(t *testing.T) {
	res := Deg2Rad(40)
	exp := 0.6981317007977318
	if res != exp {
		t.Errorf("Got %v Expected %v", res, exp)
	}
}

func TestDistance(t *testing.T) {
	res := Distance([]float64{4, 6}, []float64{8, 2})
	exp := 2.8284271247461903
	if res != exp {
		t.Errorf("Got %v, Expect %v", res, exp)
	}
}

func TestCompareTrio(t *testing.T) {
	res := CompareTrio(1, 2, 3)
	if res != false {
		t.Errorf("Got %v, Expected %v", res, false)
	}
}

func TestCompareDuo(t *testing.T) {
	res := CompareDuo(1, 1)
	if res != true {
		t.Errorf("Got %v, Expected %v", res, true)
	}
}
