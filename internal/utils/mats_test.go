package utils

import (
	"math"
	"testing"
)

func TestCot(t *testing.T) {
	result := Cot(60)
	expect := 3.0
	if math.Trunc(result) != expect {
		t.Errorf("expect %f, actual %f", expect, result)
	}
}

func TestInteriorAngle(t *testing.T) {
	geom := 5
	expectedInteriorAngle := 45.0
	observedInteriorAngle := InteriorAngle(geom)
	if expectedInteriorAngle != observedInteriorAngle {
		t.Errorf("expect %f, actual %f", expectedInteriorAngle, observedInteriorAngle)
	}
}

func TestSumOfExteriorAngle(t *testing.T) {
	result := SumOfInteriorAngle(5)
	expect := 540
	if result != expect {
		t.Errorf("SumOfInteriorAngle(5)=%d, want %d", result, expect)
	}
}

func TestExteriorAngle(t *testing.T) {
	result := ExteriorAngle(5)
	expect := 72.0
	if result != expect {
		t.Errorf("ExteriorAngle(5)=%f, want %f", result, expect)
	}
}
