package functs

import (
	"math"
	"testing"
)

// perimeter of triangle funct_tests
func TestPerimeterOfTriangle(t *testing.T) {
	res := PerimeterOfTriangle(10.0, 10.0, 10.0)
	exp := math.Trunc(30)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of square funct_tests
func TestPerimeterOfSquare(t *testing.T) {
	res := PerimeterOfSquare(5.0)
	exp := math.Trunc(20.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

// perimeter of rectangle funct_tests
func TestPerimeterOfRegularRectangle(t *testing.T) {
	res := PerimeterOfRegularRectangle(10.0, 5)
	exp := math.Trunc(30.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of rectangle funct_tests
func TestPerimeterOfIrregularRectangle(t *testing.T) {
	res := PerimOfIrr4Sided(1, 5, 4, 8)
	exp := math.Trunc(18.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of parallelogram funct_tests
func TestPerimeterOfParallelogram(t *testing.T) {
	res := PerimeterOfParallelogram(2.0, 5.0)
	exp := math.Trunc(14.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of kite funct_tests
func TestPerimOfKite(t *testing.T) {
	res := PerimOfRKite(5.0, 4.0)
	exp := math.Trunc(18.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of rhombus funct_tests
func TestPerimOfRhombus(t *testing.T) {
	res := PerimOfRRhombus(5.0)
	exp := math.Trunc(20.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of trapezoid funct_tests
func TestPerimOfTrapezoid(t *testing.T) {
	res := PerimOfRTrapezoid(5.0, 10.0, 5.0, 20.0)
	exp := math.Trunc(40.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of pentagon funct_tests
func TestPerimOfPentagon(t *testing.T) {
	res := PerimOfRPentagon(5.0)
	exp := math.Trunc(25.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of hexagon funct_tests
func TestPerimOfHexagon(t *testing.T) {
	res := PerimOfRHexagon(5.0)
	exp := math.Trunc(30.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of heptagon funct_tests
func TestPerimOfHeptagon(t *testing.T) {
	res := PerimOfRHeptagon(5.0)
	exp := math.Trunc(35.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of octagon funct_tests
func TestPerimOfOctagon(t *testing.T) {
	res := PerimOfROctagon(5.0)
	exp := math.Trunc(40.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of nonagon funct_tests
func TestPerimOfNonagon(t *testing.T) {
	res := PerimOfRNonagon(5.0)
	exp := math.Trunc(45.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of decagon funct_tests
func TestPerimOfDecagon(t *testing.T) {
	res := PerimOfRDecagon(5.0)
	exp := math.Trunc(50.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of regular generic shapes funct_tests
func TestPerimOfRegularGeneric(t *testing.T) {
	res := PerimOfRGeneral(5.0, 15.0)
	exp := math.Trunc(75.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// perimeter of irregular generic shape funct_tests
func TestPerimOfIrregularGeneric(t *testing.T) {
	res := PerimOfIrGeneral(5.0, 6.0, 5.0, 7.0, 8.0)
	exp := math.Trunc(31.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}
