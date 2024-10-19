package tests

import (
	"math"
	"testing"
)

// funct_tests surface area of a right-angled pyramid
func TestSurfAreaOfRightAngledPyramid(t *testing.T) {
	var res = SurfAreaOfRightAngledPyramid(10.0, 10.0, 10.0)
	var exp = 300.0
	if exp != res {
		t.Errorf("Expected %v Got %v", exp, res)
	}
}

// funct_tests surface area of tetrahedron
func TestSurfAreaOfTetrahedron(t *testing.T) {
	var res = SurfAreaOfTetrahedron(10)
	var exp = math.Trunc(173.20508075688772)
	if exp != math.Trunc(res) {
		t.Errorf("Expected %v Got %v", exp, res)
	}
}

// funct_tests surface area of cube
func TestSurfAreaOfCube(T *testing.T) {
	var res = SurfAreaOfCube(10)
	var exp = 60.0
	if res != exp {
		T.Errorf("Expected %v Got %v", exp, res)
	}
}

// funct_tests surface area of a cuboid
func TestSurfAreaOfCuboid(t *testing.T) {
	var res = SurfAreaOfCuboid(3.0, 5.0, 7.0)
	var exp = 30.0
	if res != exp {
		t.Errorf("Expected %v Got %v", exp, res)
	}

}

// funct_tests surface of a heptahedron
func TestSurfAreaOfHeptahedron(t *testing.T) {
	var res = SurfAreaOfHeptahedron(10, 10.0)
	var exp = 175.00
	if math.Trunc(res) != exp {
		t.Errorf("Expected %v Got %v", exp, res)
	}
}

// funct_tests surface area of a pentagon
func TestSurfAreaOfPentahedron(t *testing.T) {
	var res = SurfAreaOfPentahedron(10.0, 5.0, 8)
	var exp = 30.0
	if math.Trunc(res) != exp {
		t.Errorf("Expected %v Got %v", exp, res)
	}

}

// funct_tests surface area of right cylinder
func TestSurfAreaOfRightCylinder(t *testing.T) {
	var res = SurfAreaOfRightAngledCylinder(10.0, 10.0)
	var exp = 1256.6370614359173
	if res != exp {
		t.Errorf("Expected %v Got %v", exp, res)
	}

}

// funct_tests surface area of a hexahedron
func TestSurfAreaOfHexahedron(t *testing.T) {
	var res = SurfAreaOfHexahedron(10.0)
	var exp = math.Trunc(600.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %v got %v", exp, res)
	}
}

// funct_tests surface area of sphere
func TestSurfAreaOfSphere(t *testing.T) {
	var res = SurfAreaOfSphere(10.0)
	var exp = math.Trunc(1256.63706)
	if math.Trunc(res) != exp {
		t.Errorf("expected %v got %v", exp, res)
	}
}

// funct_tests surface area of torus
func TestSurfAreaOfTorus(t *testing.T) {
	var res = SurfAreaOfTorus(10.0, 5)
	var exp = math.Trunc(1973.921)
	if math.Trunc(res) != exp {
		t.Errorf("expected %v got %v", exp, res)
	}
}
