package props

import (
	"math"
)

// VolumeOfTetrahedron return volume of a tetrahedron
func VolumeOfTetrahedron(side float64) float64 {
	return (side * side * side) / (6 * math.Sqrt(2))
}

// VolumeOfRightAngledTrianglePyramid return volume of a right-angled pyramid
func VolumeOfRightAngledTrianglePyramid(
	baseBaseTriangle,
	heightBaseTriangle,
	heightPyramid float64) float64 {
	return (baseBaseTriangle * heightBaseTriangle * heightPyramid) / 6
}

// VolOfCube return volume of a cube
func VolOfCube(side float64) float64 {
	return side * side * side
}

// VolumeOfCuboid return volume of a cuboid
func VolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}

// VolumeOfPentahedron return volume of a pentahedron
func VolumeOfPentahedron(lengthOfBase, height float64) float64 {
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

// VolumeOfRectanglePrism return volume of a rectangular prism
func VolumeOfRectanglePrism(length, width, height float64) float64 {
	return length * width * height
}

// VolumeOfTrianglePrism return volume of triangular prism with an equilateral base.
func VolumeOfTrianglePrism(baseSideA, baseSideB, height float64) float64 {
	areaOfBase := baseSideA * baseSideB
	return areaOfBase * height
}

// VolumeOfPentagonPrism return volume of pentagonal prism
func VolumeOfPentagonPrism(baseEdge, height float64) float64 {
	factor := math.Sqrt(5*(5+2*math.Sqrt(5))) / 4
	return factor * math.Pow(baseEdge, 2) * height
}

// VolumeOfHexagonPrism return volume of hexagonal prism
func VolumeOfHexagonPrism(baseEdge, height float64) float64 {
	return (3 * math.Sqrt(3)) / 2 * math.Pow(baseEdge, 2) * height
}

// VolumeOfOctagonPrism return volume of octagonal prism
func VolumeOfOctagonPrism(baseEdge, height float64) float64 {
	factor := 2 * (1 + math.Sqrt(2))
	return factor * math.Pow(baseEdge, 2) * height
}

// VolumeOfSquarePyramid return volume of square pyramid
func VolumeOfSquarePyramid(baseEdge, height float64) float64 {
	return math.Pow(baseEdge, 2) * height / 3
}

// VolumeOfTriangularPyramid return volume of a triangular pyramid
func VolumeOfTriangularPyramid(areaOfBase, height float64) float64 {
	return (areaOfBase * height) / 3
}

// VolumeOfPentagonalPyramid return volume of pentagonal pyramid
func VolumeOfPentagonalPyramid(baseEdge, height float64) float64 {
	return (5 * math.Tan(54.0) * height * math.Pow(baseEdge, 2)) / 12
}

// VolumeOfHexagonalPyramid return volume of hexagonal pyramid
func VolumeOfHexagonalPyramid(baseEdge, height float64) float64 {
	return (math.Sqrt(3) * math.Pow(baseEdge, 2) * height) / 2
}

// VolumeOfConeFrustum return volume of cone frustum
func VolumeOfConeFrustum(radius, height float64) float64 {
	return math.Pi * math.Pow(radius, 2) * height / 3
}

// VolumeOfCylindericalFrustum return volume of a cylinder frustum or right cylinder
func VolumeOfCylindericalFrustum(radius, height float64) float64 {
	return math.Pi * math.Pow(radius, 2) * height
}

// VolumeOfEllipsoid return volume of ellipsoid
func VolumeOfEllipsoid(semiAxisA, semiAxisB, semiAxisC float64) float64 {
	return (4 * math.Pi * semiAxisA * semiAxisB * semiAxisC) / 3
}

// VolumeOfTorus return volume of a torus
func VolumeOfTorus(majorRadius, minorRadius float64) float64 {
	return math.Pi * math.Pow(minorRadius, 2) * 2 * math.Pi * majorRadius
}

// VolumeOfDecahedron return volume of a decahedron
func VolumeOfDecahedron(sides float64) (factor float64) {
	factor = (15 + (7 * math.Sqrt(5.0))) / 4
	return factor * math.Pow(sides, 3)
}

// VolumeOfIcosahedron return volume of an icosahedron
func VolumeOfIcosahedron(side float64) (factor float64) {
	factor = 5 * (3.0 + math.Sqrt(5.0)) / 12
	return factor * math.Pow(side, 3)
}
