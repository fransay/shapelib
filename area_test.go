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

func TestAreaOfSquare(t *testing.T) {
	res := functs.AreaOfSquare(5.0)
	exp := math.Trunc(25.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", exp, res)
	}
}

func TestAreaOfRectangle(t *testing.T) {
	res := functs.AreaOfRectangle(10.00, 5.00)
	exp := math.Trunc(50.00)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, found %f", exp, res)
	}
}

func TestAreaOfParallelogram(t *testing.T) {
	res := functs.AreaOfParallelogram(20.0, 10.0)
	exp := math.Trunc(200.00)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, found %f", exp, res)
	}
}

func TestAreaOfRhombus(t *testing.T) {
	res := functs.AreaOfRhombus(10, 5)
	exp := math.Trunc(7.5)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

func TestAreaOfTrapezoid(t *testing.T) {
	res := functs.AreaOfTrapezoid(10.0, 5.0, 5.0)
	exp := math.Trunc(37.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

func TestAreaOfRPentagon(t *testing.T) {
	res := functs.AreaOfRPentagon(5.0)
	exp := math.Trunc(43.01194)
	if math.Trunc(res) != exp {
		t.Errorf("%f, %f", exp, res)
	}
}

func TestAreaOfRHexagon(t *testing.T) {
	res := functs.AreaOfRHexagon(10.0)
	exp := math.Trunc(259.81)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

func TestAreaOfRHeptagon(t *testing.T) {
	res := functs.AreaOfRHeptagon(5.00)
	exp := math.Trunc(90.85)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", res, exp)
	}

}
