package utils

import "testing"

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
