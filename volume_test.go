package shapelib

import (
	"math"
	functs "shapelib/functs/3D"
	"testing"
)

func TestVolOfTetrahedron(t *testing.T) {
	var result = functs.VolOfTetrahedron(10)
	var expected = math.Trunc(117.80)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}
func TestVolOfRightAngledTriPyramid(t *testing.T) {
	var result = functs.VolOfRightAngledTriPyramid(10, 10, 10)
	var expected = math.Trunc(166.666)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}
