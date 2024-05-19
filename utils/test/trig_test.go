package test

import (
	"math"
	"shapelib/utils"
	"testing"
)

func TestCot(t *testing.T) {
	result := utils.Cot(60)
	expect := 3.0
	if math.Trunc(result) != expect {
		t.Errorf("expect %f, actual %f", expect, result)
	}
}

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
