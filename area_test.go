package shapelib

import (
	"math"
	"shapelib/functs"
	"shapelib/types"
	"testing"
)

func TestAreaOfTriangleBH(t *testing.T) {
	res1 := functs.AreaOfTriangleBH(10.0, 5.0)
	expected1 := 25.0
	if res1 != expected1 {
		t.Errorf("expected %f, got %f", expected1, res1)
	}
	res2 := functs.AreaOfTriangleBH(15.0, 2.0)
	expected2 := 15.0
	if res2 != expected2 {
		t.Errorf("expected %f, got %f", expected2, res2)
	}
	res3 := functs.AreaOfTriangleBH(1.0, 1.0)
	expected3 := 0.5
	if res3 != expected3 {
		t.Errorf("expected %f, got %f", expected3, res3)
	}
	res4 := functs.AreaOfTriangleBH(1.0, 5.0)
	expected4 := 2.5
	if res4 != expected4 {
		t.Errorf("expected %f, got %f", expected4, res4)
	}

}

func TestAreaOfTriangleSide(t *testing.T) {
	res := functs.AreaOfTriangleSide(12.0, 8.0, 10.00)
	exp := 0.0 // inaccuracies increase
	if exp != res {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

func TestAreaOfTriangleAngle(t *testing.T) {
	res := functs.AreaOfTriangleAngle(10.00, 5.00, 120.00)
	exp := math.Trunc(14.515280)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

func TestAreaOfSquare(t *testing.T) {
	res := functs.AreaOfSquare(5.0)
	exp := math.Trunc(25.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", exp, res)
	}
}

func TestAreaOfRectangle(t *testing.T) {
	res := functs.AreaOfRectangle(10.00, 5.00)
	exp := math.Trunc(50.00)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, found %f", exp, res)
	}
}

func TestAreaOfParallelogram(t *testing.T) {
	res := functs.AreaOfParallelogram(20.0, 10.0)
	exp := math.Trunc(200.00)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, found %f", exp, res)
	}
}

func TestAreaOfRhombus(t *testing.T) {
	res := functs.AreaOfRhombus(10, 5)
	exp := math.Trunc(7.5)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

func TestAreaOfTrapezoid(t *testing.T) {
	res := functs.AreaOfTrapezoid(10.0, 5.0, 5.0)
	exp := math.Trunc(37.0)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}

}

func TestAreaOfRPentagon(t *testing.T) {
	res := functs.AreaOfRPentagon(5.0)
	exp := math.Trunc(43.01194)
	if math.Trunc(res) != exp {
		t.Errorf("%f, %f", exp, res)
	}
}

func TestAreaOfRHexagon(t *testing.T) {
	res := functs.AreaOfRHexagon(10.0)
	exp := math.Trunc(259.81)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

func TestAreaOfROctagon(t *testing.T) {
	res := functs.AreaOfROctagon(5.00)
	exp := math.Trunc(120.7)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", exp, res)
	}

}
func TestAreaOfRNonagon(t *testing.T) {
	res := functs.AreaOfRNonagon(5.00)
	exp := math.Trunc(154.55)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", exp, res)
	}

}
func TestAreaOfRDecagon(t *testing.T) {
	res := functs.AreaOfRDecagon(5.00)
	exp := math.Trunc(192.355)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", exp, res)
	}

}
func TestAreaOfRHeptagon(t *testing.T) {
	res := functs.AreaOfRHeptagon(5.00)
	exp := math.Trunc(90.85)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", res, exp)
	}

}

func TestAreaOfCircle(t *testing.T) {
	res := functs.AreaOfCircle(5.00)
	exp := math.Trunc(75.00)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", res, exp)
	}
}

func TestAreaOfSemiCircle(t *testing.T) {
	res := functs.AreaOfSemiCircle(5.00)
	exp := math.Trunc(37.50)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", res, exp)
	}

}

func TestAreaOfOval(t *testing.T) {
	res := functs.AreaOfOval(10.00, 8.0)
	exp := math.Trunc(251.32)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", res, exp)

	}

}

func TestAreaOfEllipse(t *testing.T) {
	res := functs.AreaOfEllipse(10.00, 8.0)
	exp := math.Trunc(251.32)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", res, exp)

	}

}

func TestAreaCoordinates(t *testing.T) {
	var res = functs.AreaCoordinates(types.Point2D{X: 1, Y: 4}, types.Point2D{X: 3, Y: 6}, types.Point2D{X: 8, Y: 5})
	var exp = math.Trunc(-19.500)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f got %f", exp, res)
	}

}
