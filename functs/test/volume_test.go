package test

import (
	"math"
	"shapelib/functs"
	"testing"
)

// test volume of tetrahedron
func TestVolOfTetrahedron(t *testing.T) {
	var result = functs.VolOfTetrahedron(10)
	var expected = math.Trunc(117.80)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of right angle triangular pyramid
func TestVolOfRightAngledTriPyramid(t *testing.T) {
	var result = functs.VolOfRightAngledTriPyramid(10, 10, 10)
	var expected = math.Trunc(166.666)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of cube
func TestVolOfCube(t *testing.T) {
	var result = functs.VolOfCube(10)
	var expected = math.Trunc(1000.000)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of cuboid
func TestVolOfCuboid(t *testing.T) {
	var result = functs.VolOfCuboid(10, 5.0, 7.0)
	var expected = math.Trunc(350.00)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of pentahedron
func TestVolOfPentahedron(t *testing.T) {
	var result = functs.VolOfPentahedron(10.00, 5.00)
	var expected = math.Trunc(166.666)
	if math.Trunc(result) != expected {
		t.Errorf("expected %v got %v", expected, result)
	}
}

// test volume of cone
func TestVolOfCone(t *testing.T) {
	var result = functs.VolOfCone(10.00, 5.00)
	var expected = 523.5987755982989
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of cylinder
func TestVolOfCylinder(t *testing.T) {
	var result = functs.VolOfCylinder(10.00, 5.00)
	var expected = 50.0
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}

}

// test volume of rect prism
func TestVolOfRectPrism(t *testing.T) {
	var result = functs.VolOfRectPrism(10.00, 5.00, 10.0)
	var expected = 500.0
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}
}

func TestVolOfTriPrism(t *testing.T) {
	var result = functs.VolOfTriPrism(10.00, 5.00, 10.0)
	var expected = 500.0
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}
}

// test volume of pentagonal prism
func TestVolOfPentPrism(t *testing.T) {
	var result = functs.VolOfPentPrism(10.00, 5.00)
	var expected = 860.2387002944835
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}
}

// test volume of hexagonal prism
func TestVolOfHexaPrism(t *testing.T) {
	var result = functs.VolOfHexaPrism(10.00, 5.00)
	var expected = 1299.038105676658
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}
}

// test volume of octagonal prism
func TestVolOfOctPrism(t *testing.T) {
	var result = functs.VolOfOctPrism(5, 10)
	var expected = 1207.1067811865473
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of a square pyramid
func TestVolOfSquarePyramid(t *testing.T) {
	var result = functs.VolOfSquarePymd(5, 10)
	var expected = 83.33333333333333
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of triangular pyramid
func TestVolOfTriPymd(t *testing.T) {
	var result = functs.VolOfTriPymd(5, 10)
	var expected = 16.666666666666668
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of a pentagonal pyramids
func TestVolOfPentPymd(t *testing.T) {
	var result = functs.VolOfPentPymd(5.0, 10.0)
	var expected = 70.1875104841729
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of hexagonal pyramid
func TestVolOfHexaPymd(t *testing.T) {
	var result = functs.VolOfHexaPymd(5, 10)
	var expected = 216.50635094610965
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of cone frustrum
func TestVolOfConeFrustrum(t *testing.T) {
	var result = functs.VolOfConeFrustum(5, 10)
	var expected = 261.79938779914943
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of cylinder frustum
func TestVolOfCyFrustrum(t *testing.T) {
	var result = functs.VolOfCyFrustum(5, 10)
	var expected = 785.3981633974483
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of ellipsoid
func TestVolOfEllipsoid(t *testing.T) {
	var result = functs.VolOfEllipsoid(5, 10, 15)
	var expected = 3141.592653589793
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of torus
func TestVolOfTorus(t *testing.T) {
	var result = functs.VolOfTorus(5, 2)
	var expected = 394.78417604357435
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of a decahedron
func TestVolOfDecahedron(t *testing.T) {
	var result = functs.VolOfDecahedron(5)
	var expected = 957.8898700780791
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

// test volume of icosahedron
func TestVolOfIcosahedron(t *testing.T) {
	var result = functs.VolOfIcosahedron(5)
	var expected = 272.71187382811405
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}
