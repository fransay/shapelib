package shapelib

import (
	"math"
	"shapelib/functs"
	"testing"
)

// perimeter of triangle test
func TestPerimeterOfTriangle(t *testing.T) {
	res := functs.PerimOfRTriangle(10.0, 10.0, 10.0)
	exp := math.Trunc(30)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of square test
func TestPerimeterOfSquare(t *testing.T) {
	res := functs.PerimOfRSquare(5.0)
	exp := math.Trunc(20.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

// perimeter of rectangle test
func TestPerimeterOfRectangle(t *testing.T) {
	res := functs.PerimOfRRectangle(10.0, 5)
	exp := math.Trunc(30.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of parallelogram test
func TestPerimeterOfParallelogram(t *testing.T) {
	res := functs.PerimOfRParallelogram(2.0, 5.0)
	exp := math.Trunc(14.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of kite test
func TestPerimOfKite(t *testing.T) {
	res := functs.PerimOfRKite(5.0, 4.0)
	exp := math.Trunc(18.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of rhombus test
func TestPerimOfRhombus(t *testing.T) {
	res := functs.PerimOfRRhombus(5.0)
	exp := math.Trunc(20.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of trapezoid test
func TestPerimOfTrapezoid(t *testing.T) {
	res := functs.PerimOfRTrapezoid(5.0, 10.0, 5.0, 20.0)
	exp := math.Trunc(40.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of pentagon test
func TestPerimOfPentagon(t *testing.T) {
	res := functs.PerimOfRPentagon(5.0)
	exp := math.Trunc(25.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of hexagon test
func TestPerimOfHexagon(t *testing.T) {
	res := functs.PerimOfRHexagon(5.0)
	exp := math.Trunc(30.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of heptagon test
func TestPerimOfHeptagon(t *testing.T) {
	res := functs.PerimOfRHeptagon(5.0)
	exp := math.Trunc(35.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of octagon test
func TestPerimOfOctagon(t *testing.T) {
	res := functs.PerimOfROctagon(5.0)
	exp := math.Trunc(40.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of nonagon test
func TestPerimOfNonagon(t *testing.T) {
	res := functs.PerimOfRNonagon(5.0)
	exp := math.Trunc(45.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of decagon test
func TestPerimOfDecagon(t *testing.T) {
	res := functs.PerimOfRDecagon(5.0)
	exp := math.Trunc(50.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of regular generic shapes test
func TestPerimOfRegularGeneric(t *testing.T) {
	res := functs.PerimOfRegularGeneric(5.0, 15.0)
	exp := math.Trunc(75.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of irregular generic shape test
func TestPerimOfIrregularGeneric(t *testing.T) {
	res := functs.PerimOfIrregularGeneric(5.0, 6.0, 5.0, 7.0, 8.0)
	exp := math.Trunc(31.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}
