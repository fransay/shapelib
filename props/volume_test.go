package props

import (
	"math"
	"testing"
)

// Test volume of tetrahedron
func TestVolOfTetrahedron(t *testing.T) {
	var result = VolOfTetrahedron(10)
	var expected = math.Trunc(117.80)
	if math.Trunc(result) != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of right angle triangular pyramid
func TestVolOfRightAngledTriPyramid(t *testing.T) {
	var result = VolOfRightAngledTriPyramid(10, 10, 10)
	var expected = math.Trunc(166.666)
	if math.Trunc(result) != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}

}

// Test volume of cube
func TestVolOfCube(t *testing.T) {
	var result = VolOfCube(10)
	var expected = math.Trunc(1000.000)
	if math.Trunc(result) != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}

}

// Test volume of cuboid
func TestVolOfCuboid(t *testing.T) {
	var result = VolOfCuboid(10, 5.0, 7.0)
	var expected = math.Trunc(350.00)
	if math.Trunc(result) != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}

}

// Test volume of pentahedron
func TestVolOfPentahedron(t *testing.T) {
	var result = VolOfPentahedron(10.00, 5.00)
	var expected = math.Trunc(166.666)
	if math.Trunc(result) != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of cone
func TestVolOfCone(t *testing.T) {
	var result = VolOfCone(10.00, 5.00)
	var expected = 523.5987755982989
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}

}

// Test volume of cylinder
func TestVolOfCylinder(t *testing.T) {
	var result = VolOfCylinder(10.00, 5.00)
	var expected = 50.0
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}

}

// Test volume of rect prism
func TestVolOfRectPrism(t *testing.T) {
	var result = VolOfRectPrism(10.00, 5.00, 10.0)
	var expected = 500.0
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of triangular prism
func TestVolOfTriangularPrism(t *testing.T) {
	var result = VolOfTriPrism(10.00, 5.00, 10.0)
	var expected = 500.0
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of pentagonal prism
func TestVolOfPentPrism(t *testing.T) {
	var result = VolOfPentPrism(10.00, 5.00)
	var expected = 860.2387002944835
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of hexagonal prism
func TestVolOfHexagonalPrism(t *testing.T) {
	var result = VolOfHexaPrism(10.00, 5.00)
	var expected = 1299.038105676658
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of octagonal prism
func TestVolOfOctagonalPrism(t *testing.T) {
	var result = VolOfOctPrism(5, 10)
	var expected = 1207.1067811865473
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of a square pyramid
func TestVolOfSquarePyramid(t *testing.T) {
	var result = VolOfSquarePyramid(5, 10)
	var expected = 83.33333333333333
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of triangular pyramid
func TestVolOfTriPymd(t *testing.T) {
	var result = VolOfTriangularPyramid(5, 10)
	var expected = 16.666666666666668
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of a pentagonal pyramids
func TestVolOfPentPymd(t *testing.T) {
	var result = VolOfPentagonalPyramid(5.0, 10.0)
	var expected = 70.1875104841729
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of hexagonal pyramid
func TestVolOfHexaPymd(t *testing.T) {
	var result = VolOfHexagonalPyramid(5, 10)
	var expected = 216.50635094610965
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of cone frustrum
func TestVolOfConeFrustrum(t *testing.T) {
	var result = VolOfConeFrustum(5, 10)
	var expected = 261.79938779914943
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of cylinder frustum
func TestVolOfCyFrustrum(t *testing.T) {
	var result = VolOfCyFrustum(5, 10)
	var expected = 785.3981633974483
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of ellipsoid
func TestVolOfEllipsoid(t *testing.T) {
	var result = VolOfEllipsoid(5, 10, 15)
	var expected = 3141.592653589793
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of torus
func TestVolOfTorus(t *testing.T) {
	var result = VolOfTorus(5, 2)
	var expected = 394.78417604357435
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of a decahedron
func TestVolOfDecahedron(t *testing.T) {
	var result = VolOfDecahedron(5)
	var expected = 957.8898700780791
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}

// Test volume of icosahedron
func TestVolOfIcosahedron(t *testing.T) {
	var result = VolOfIcosahedron(5)
	var expected = 272.71187382811405
	if result != expected {
		t.Errorf("Expected = %v Got = %v", expected, result)
	}
}
