package props

import (
	"math"
)

// VolOfTetrahedron return volume of a tetrahedron
func VolOfTetrahedron(side float64) float64 {
	return (side * side * side) / (6 * math.Sqrt(2))
}

// VolOfRightAngledTriPyramid return volume of a right-angled pyramid
func VolOfRightAngledTriPyramid(
	baseBaseTriangle,
	heightBaseTriangle,
	heightPyramid float64) float64 {
	return (baseBaseTriangle * heightBaseTriangle * heightPyramid) / 6
}

// VolOfCube return volume of a cube
func VolOfCube(side float64) float64 {
	return (side * side * side)
}

// VolOfCuboid return volume of a cuboid
func VolOfCuboid(length, width, height float64) float64 {
	return length * width * height
}

// VolOfPentahedron return volume of a pentahedron
func VolOfPentahedron(lengthOfBase, height float64) float64 {
	return (math.Pow(lengthOfBase, 2) * height) / 3
}

// VolOfSphere return the volume of a sphere
func VolOfSphere(radius float64) float64 {
	return (4 * math.Pi * math.Pow(radius, 3)) / 3
}

// VolOfCone return volume of a cone
func VolOfCone(radius, height float64) float64 {
	return (math.Pi * math.Pow(radius, 2) * height) / 3
}

// VolOfCylinder return volume of cylinder
func VolOfCylinder(radius, height float64) float64 {
	return radius * height
}

// VolOfRectPrism return volume of a rectangular prism
func VolOfRectPrism(length, width, height float64) float64 {
	return length * width * height
}

// VolOfTriPrism return volume of triangular prism with an equilateral base.
func VolOfTriPrism(baseSideA, baseSideB, height float64) float64 {
	areaOfBase := baseSideA * baseSideB
	return areaOfBase * height
}

// VolOfPentPrism return volume of pentagonal prism
func VolOfPentPrism(baseEdge, height float64) float64 {
	factor := math.Sqrt(5*(5+2*math.Sqrt(5))) / 4
	return factor * math.Pow(baseEdge, 2) * height
}

// VolOfHexaPrism return volume of hexagonal prism
func VolOfHexaPrism(baseEdge, height float64) float64 {
	return (3 * math.Sqrt(3)) / 2 * math.Pow(baseEdge, 2) * height
}

// VolOfOctPrism return volume of octagonal prism
func VolOfOctPrism(baseEdge, height float64) float64 {
	factor := 2 * (1 + math.Sqrt(2))
	return factor * math.Pow(baseEdge, 2) * height
}

// VolOfSquarePyramid return volume of square pyramid
func VolOfSquarePyramid(baseEdge, height float64) float64 {
	return math.Pow(baseEdge, 2) * height / 3
}

// VolOfTriangularPyramid return volume of a triangular pyramid
func VolOfTriangularPyramid(areaOfBase, height float64) float64 {
	return (areaOfBase * height) / 3
}

// VolOfPentagonalPyramid return volume of pentagonal pyramid
func VolOfPentagonalPyramid(baseEdge, height float64) float64 {
	return (5 * math.Tan(54.0) * height * math.Pow(baseEdge, 2)) / 12
}

// VolOfHexagonalPyramid return volume of hexagonal pyramid
func VolOfHexagonalPyramid(baseEdge, height float64) float64 {
	return (math.Sqrt(3) * math.Pow(baseEdge, 2) * height) / 2
}

// VolOfConeFrustum return volume of cone frustum
func VolOfConeFrustum(radius, height float64) float64 {
	return math.Pi * math.Pow(radius, 2) * height / 3
}

// VolOfCyFrustum return volume of a cylinder frustum or right cylinder
func VolOfCyFrustum(radius, height float64) float64 {
	return math.Pi * math.Pow(radius, 2) * height
}

// VolOfEllipsoid return volume of ellipsoid
func VolOfEllipsoid(semiAxisA, semiAxisB, semiAxisC float64) float64 {
	return (4 * math.Pi * semiAxisA * semiAxisB * semiAxisC) / 3
}

// VolOfTorus return volume of a torus
func VolOfTorus(majorRadius, minorRadius float64) float64 {
	return math.Pi * math.Pow(minorRadius, 2) * 2 * math.Pi * majorRadius
}

// VolOfDecahedron return volume of a decahedron
func VolOfDecahedron(sides float64) (factor float64) {
	factor = (15 + (7 * math.Sqrt(5.0))) / 4
	return factor * math.Pow(sides, 3)
}

// VolOfIcosahedron return volume of an icosahedron
func VolOfIcosahedron(side float64) (factor float64) {
	factor = 5 * (3.0 + math.Sqrt(5.0)) / 12
	return factor * math.Pow(side, 3)
}