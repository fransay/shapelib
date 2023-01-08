package shapelib

import (
	"math"
	"shapelib/functs"
	"testing"
)

func TestSurfAreaOfRightAngledPyramid(t *testing.T) {
	var res = functs.SurfAreaOfRightAngledPyramid(10.0, 10.0, 10.0)
	var exp = 1000.00
	if exp != res {
		t.Errorf("expected %v got %v", exp, res)
	}
}
func TestSurfAreaOfTetrahedron(t *testing.T) {
	var res = functs.SurfAreaOfTetrahedron(10)
	var exp = math.Trunc(17.320)
	if exp != math.Trunc(res) {
		t.Errorf("expected %v got %v", exp, res)
	}

}
func TestSurfAreaOfCube(T *testing.T) {
	var res = functs.SurfAreaOfCube(10)
	var exp = 60.0
	if res != exp {
		T.Errorf("expected %v got %v", exp, res)
	}
}

func TestSurfAreaOfCuboid(t *testing.T) {
	var res = functs.SurfAreaOfCuboid(3.0, 5.0, 7.0)
	var exp = 30.0
	if res != exp {
		t.Errorf("expected %v got %v", exp, res)
	}

}

func TestSurfAreaOfHeptahedron(t *testing.T) {
	var res = functs.SurfAreaOfHeptahedron(10.0, 10.0)
	var exp = math.Trunc(1257.1)
	if math.Trunc(res) != exp {
		t.Errorf("expected %v got %v", exp, res)
	}
}

func TestSurfAreaOfPentahedron(t *testing.T) {
	var res = functs.SurfAreaOfPentahedron(10.0, 10.0)
	var exp = math.Trunc(1257.1)
	if math.Trunc(res) != exp {
		t.Errorf("expected %v got %v", exp, res)
	}

}

func TestSurfAreaOfRightCylinder(t *testing.T) {
	var res = functs.SurfAreaOfRightCylinder(10.0, 10.0)
	var exp = math.Trunc(1257.1)
	if math.Trunc(res) != exp {
		t.Errorf("expected %v got %v", exp, res)
	}

}

func TestSurfAreaOfHexahedron(t *testing.T) {
	var res = functs.SurfAreaOfHexahedron(10.0)
	var exp = math.Trunc(600.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %v got %v", exp, res)
	}

}
