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

func TestPerimeterOfSquare(t *testing.T) {
	res := functs.PerimOfRSquare(5.0)
	exp := math.Trunc(20.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

func TestPerimeterOfRectangle(t *testing.T) {
	res := functs.PerimOfRRectangle(10.0, 5)
	exp := math.Trunc(30.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

func TestPerimeterOfParallelogram(t *testing.T) {
	res := functs.PerimOfRParallelogram(2.0, 5.0)
	exp := math.Trunc(14.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

func TestPerimOfKite(t *testing.T) {
	res := functs.PerimOfRKite(5.0, 4.0)
	exp := math.Trunc(18.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

func TestPerimOfRhombus(t *testing.T) {
	res := functs.PerimOfRRhombus(5.0)
	exp := math.Trunc(20.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}
