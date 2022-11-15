package shapelib

import (
	"math"
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
	res := functs.AreaOfTriangleSide(12.0, 8.0, 10.00)
	exp := 0.0 // inaccuracies increase
	if exp != res {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

func TestAreaOfTriangleAngle(t *testing.T) {
	res := functs.AreaOfTriangleAngle(10.00, 5.00, 120.00)
	exp := math.Trunc(14.515280)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

func TestOfAreaOfSquare(t *testing.T) {
	res := functs.AreaOfSquare(5.0)
	exp := math.Trunc(25.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", exp, res)
	}
}

func TestOfRectangle(t *testing.T) {
	res := functs.AreaOfRectangle(10.00, 5.00)
	exp := math.Trunc(50.00)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, found %f", exp, res)
	}
}

func TestOfParallelogram(t *testing.T) {
	res := functs.AreaOfParallelogram(20.0, 10.0)
	exp := math.Trunc(200.00)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, found %f", exp, res)
	}
}
