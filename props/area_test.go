package props

import (
	"math"
	"shapelib/geom"
	"testing"
)

// Test area of triangle given base and height
func TestAreaOfTriangleBH(t *testing.T) {
	res1 := AreaOfTriangleBH(10.0, 5.0)
	expected1 := 25.0
	if res1 != expected1 {
		t.Errorf("Expected %f, Got %f", expected1, res1)
	}
	res2 := AreaOfTriangleBH(15.0, 2.0)
	expected2 := 15.0
	if res2 != expected2 {
		t.Errorf("Expected %f, Got %f", expected2, res2)
	}
	res3 := AreaOfTriangleBH(1.0, 1.0)
	expected3 := 0.5
	if res3 != expected3 {
		t.Errorf("Expected %f, Got %f", expected3, res3)
	}
	res4 := AreaOfTriangleBH(1.0, 5.0)
	expected4 := 2.5
	if res4 != expected4 {
		t.Errorf("Expected %f, Got %f", expected4, res4)
	}

}

// Test area of a triangle given the sides
func TestAreaOfTriangleSide(t *testing.T) {
	res := AreaOfTriangleSide(12.0, 8.0, 10.00)
	exp := 0.0 // inaccuracies increase
	if exp != res {
		t.Errorf("Expected %f, Got %f", exp, res)
	}

}

// Test area of a triangle given the angles
func TestAreaOfTriangleAngle(t *testing.T) {
	res := AreaOfTriangleAngle(10.00, 5.00, 120.00)
	exp := math.Trunc(14.515280)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f, Got %f", exp, res)
	}

}

// Test area of a square
func TestAreaOfSquare(t *testing.T) {
	res := AreaOfSquare(5.0)
	exp := math.Trunc(25.0)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a rectangle
func TestAreaOfRectangle(t *testing.T) {
	res := AreaOfRectangle(10.00, 5.00)
	exp := math.Trunc(50.00)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f, Got %f", exp, res)
	}
}

// Test area of a parallelogram
func TestAreaOfParallelogram(t *testing.T) {
	res := AreaOfParallelogram(20.0, 10.0)
	exp := math.Trunc(200.00)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f, Got %f", exp, res)
	}
}

// Test area of a rhombus
func TestAreaOfRhombus(t *testing.T) {
	res := AreaOfRhombus(10, 5)
	exp := math.Trunc(7.5)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f, Got %f", exp, res)
	}
}

// Test area of a trapezoid
func TestAreaOfTrapezoid(t *testing.T) {
	res := AreaOfTrapezoid(10.0, 5.0, 5.0)
	exp := math.Trunc(37.0)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f, Got %f", exp, res)
	}

}

// Test area of a regular pentagon
func TestAreaOfRPentagon(t *testing.T) {
	res := AreaOfRPentagon(5.0)
	exp := math.Trunc(43.01194)
	if math.Trunc(res) != exp {
		t.Errorf(" Expected %f, Got %f", exp, res)
	}
}

// Test area of a regular hexagon
func TestAreaOfRHexagon(t *testing.T) {
	res := AreaOfRHexagon(10.0)
	exp := math.Trunc(259.81)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f, Got %f", exp, res)
	}
}

// Test area of a regular octagon
func TestAreaOfROctagon(t *testing.T) {
	res := AreaOfROctagon(5.00)
	exp := math.Trunc(120.7)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}

}

// Test area of a regular nonagon
func TestAreaOfRNonagon(t *testing.T) {
	res := AreaOfRNonagon(5)
	exp := math.Trunc(137.373850)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a regular decagon
func TestAreaOfRDecagon(t *testing.T) {
	res := AreaOfRDecagon(5.00)
	exp := math.Trunc(192.355)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a regular heptagon
func TestAreaOfRHeptagon(t *testing.T) {
	res := AreaOfRHeptagon(5.00)
	exp := math.Trunc(90.85)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a regular dodecagon
func TestAreaOfRDodecagon(t *testing.T) {
	res := AreaOfDodecagon(5)
	exp := math.Trunc(64.9519052838)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a circle
func TestAreaOfCircle(t *testing.T) {
	res := AreaOfCircle(5.00)
	exp := math.Trunc(75.00)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a semicircle
func TestAreaOfSemiCircle(t *testing.T) {
	res := AreaOfSemiCircle(5.00)
	exp := math.Trunc(37.50)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a quad circle
func TestAreaOfQuadCircle(t *testing.T) {
	res := AreaOfQuadCircle(5.00)
	exp := math.Trunc(18.75)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of an oval
func TestAreaOfOval(t *testing.T) {
	res := AreaOfOval(10.00, 8.0)
	exp := math.Trunc(251.32)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)

	}
}

// Test area of an ellipse
func TestAreaOfEllipse(t *testing.T) {
	res := AreaOfEllipse(10.00, 8.0)
	exp := math.Trunc(251.32)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)

	}
}

// Test area given coordinates
func TestAreaCoordinates(t *testing.T) {
	var res = AreaCoordinates(geom.Point2D{X: 1, Y: 4}, geom.Point2D{X: 3, Y: 6}, geom.Point2D{X: 8, Y: 5})
	var exp = math.Trunc(-19.500)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}
