package shapelib

import (
	"shapelib/functs/2D"
	"testing"
)

func TestAreaOfTriangleBH(t *testing.T) {
	res := functs.AreaOfTriangleBH(10.0, 5.0)
	expected := 25.0
	if res != expected {
		t.Errorf("expected %f, got %f", expected, res)
	}
}
