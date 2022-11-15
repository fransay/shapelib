package shapelib

import (
	"math"
	functs "shapelib/functs/2D"
	"testing"
)

func TestPerimeterOfTriangle(t *testing.T) {
	res := functs.PerimOfRTriangle(10.0, 10.0, 10.0)
	exp := math.Trunc(30)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}
