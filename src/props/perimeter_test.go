package props

import (
	"math"
	"testing"
)

// Perimeter of triangle
func TestPerimeterOfTriangle(t *testing.T) {
	res := PerimeterOfTriangle(10.0, 10.0, 10.0)
	exp := math.Trunc(30)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of square
func TestPerimeterOfSquare(t *testing.T) {
	res := PerimeterOfSquare(5.0)
	exp := math.Trunc(20.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

// Perimeter of rectangle
func TestPerimeterOfRegularRectangle(t *testing.T) {
	res := PerimeterOfRegularRectangle(10.0, 5)
	exp := math.Trunc(30.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of rectangle
func TestPerimeterOfIrregularRectangle(t *testing.T) {
	res := PerimeterOfIrregularFourSided(1, 5, 4, 8)
	exp := math.Trunc(18.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of parallelogram
func TestPerimeterOfParallelogram(t *testing.T) {
	res := PerimeterOfParallelogram(2.0, 5.0)
	exp := math.Trunc(14.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of kite
func TestPerimOfKite(t *testing.T) {
	res := PerimeterOfRegularKite(5.0, 4.0)
	exp := math.Trunc(18.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of rhombus
func TestPerimOfRhombus(t *testing.T) {
	res := PerimeterOfRegularRhombus(5.0)
	exp := math.Trunc(20.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of trapezoid
func TestPerimOfTrapezoid(t *testing.T) {
	res := PerimeterOfRegularTrapezoid(5.0, 10.0, 5.0, 20.0)
	exp := math.Trunc(40.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of pentagon
func TestPerimOfPentagon(t *testing.T) {
	res := PerimeterOfRegularPentagon(5.0)
	exp := math.Trunc(25.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of hexagon
func TestPerimOfHexagon(t *testing.T) {
	res := PerimeterOfRegularHexagon(5.0)
	exp := math.Trunc(30.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of heptagon
func TestPerimOfHeptagon(t *testing.T) {
	res := PerimeterOfRegularHeptagon(5.0)
	exp := math.Trunc(35.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of octagon
func TestPerimOfOctagon(t *testing.T) {
	res := PerimeterOfRegularOctagon(5.0)
	exp := math.Trunc(40.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of nonagon
func TestPerimOfNonagon(t *testing.T) {
	res := PerimeterOfRegularNonagon(5.0)
	exp := math.Trunc(45.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of decagon
func TestPerimOfDecagon(t *testing.T) {
	res := PerimeterOfRegularDecagon(5.0)
	exp := math.Trunc(50.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of regular generic shapes
func TestPerimOfRegularGeneric(t *testing.T) {
	res := PerimeterOfRegularGeneric(5.0, 15.0)
	exp := math.Trunc(75.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

// Perimeter of irregular generic shape
func TestPerimOfIrregularGeneric(t *testing.T) {
	res := PerimeterOfIrregularGeneric(5.0, 6.0, 5.0, 7.0, 8.0)
	exp := math.Trunc(31.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}
