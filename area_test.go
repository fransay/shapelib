package shapelib

import (
	"shapelib/functs/2D"
	"testing"
)

func TestAreaOfTriangleBH(t *testing.T) {
	res1 := functs.AreaOfTriangleBH(10.0, 5.0)
	expected1 := 25.0
	if res1 != expected1 {
		t.Errorf("expected %f, got %f", expected1, res1)
	}
	res2 := functs.AreaOfTriangleBH(15.0, 2.0)
	expected2 := 15.0
	if res2 != expected2 {
		t.Errorf("expected %f, got %f", expected2, res2)
	}
	res3 := functs.AreaOfTriangleBH(1.0, 1.0)
	expected3 := 0.5
	if res3 != expected3 {
		t.Errorf("expected %f, got %f", expected3, res3)
	}
	res4 := functs.AreaOfTriangleBH(1.0, 5.0)
	expected4 := 2.5
	if res4 != expected4 {
		t.Errorf("expected %f, got %f", expected4, res4)
	}

}

func TestAreaOfTriangleSide(t *testing.T) {

}
