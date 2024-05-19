package test

import (
	"shapelib/utils"
	"testing"
)

func TestSumOfExteriorAngle(t *testing.T) {
	result := utils.SumOfInteriorAngle(5)
	expect := 540
	if result != expect {
		t.Errorf("SumOfInteriorAngle(5)=%d, want %d", result, expect)
	}
}

func TestExteriorAngle(t *testing.T) {
	result := utils.ExteriorAngle(5)
	expect := 72.0
	if result != expect {
		t.Errorf("ExteriorAngle(5)=%f, want %f", result, expect)
	}
}
