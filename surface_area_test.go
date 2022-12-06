package shapelib

import (
	functs "shapelib/functs/3D"
	"testing"
)

func TestSurfAreaOfCube(T *testing.T) {
	var res = functs.SurfAreaOfCube(10)
	var exp = 60.0
	if res != exp {
		T.Errorf("expected %v got %v", exp, res)
	}
}
