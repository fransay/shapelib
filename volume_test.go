package shapelib

import (
	"math"
	"shapelib/functs"
	"testing"
)

// test volume of tetrahedron
func TestVolOfTetrahedron(t *testing.T) {
	var result = functs.VolOfTetrahedron(10)
	var expected = math.Trunc(117.80)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of right angle triangular pyramid
func TestVolOfRightAngledTriPyramid(t *testing.T) {
	var result = functs.VolOfRightAngledTriPyramid(10, 10, 10)
	var expected = math.Trunc(166.666)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of cube
func TestVolOfCube(t *testing.T) {
	var result = functs.VolOfCube(10)
	var expected = math.Trunc(1000.000)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of cuboid
func TestVolOfCuboid(t *testing.T) {
	var result = functs.VolOfCuboid(10, 5.0, 7.0)
	var expected = math.Trunc(350.00)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of pentahedron
func TestVolOfPentahedron(t *testing.T) {
	var result = functs.VolOfPentahedron(10.00, 5.00)
	var expected = math.Trunc(166.666)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}
