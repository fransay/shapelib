package functs

import (
	"math"
)

// VolOfTetrahedron return volume of a tetrahedron
func VolOfTetrahedron(side float64) (vol float64) {
	vol = (side * side * side) / (6 * math.Sqrt(2))
	return vol
}

// VolOfRightAngledTriPyramid volume of a right-angled pyramid
func VolOfRightAngledTriPyramid(baseBaseTriangle, heightBaseTriangle, heightPyramid float64) (vol float64) {
	vol = (baseBaseTriangle * heightBaseTriangle * heightPyramid) / 6
	return vol
}

// VolOfCube volume of a cube
func VolOfCube(side float64) (vol float64) {
	vol = side * side * side
	return vol
}

// VolOfCuboid volume of a cuboid
func VolOfCuboid(length, width, height float64) (vol float64) {
	vol = length * width * height
	return vol
}

// VolOfPentahedron volume of a pentahedron
func VolOfPentahedron(lengthOfBase, height float64) (vol float64) {
	vol = (math.Pow(lengthOfBase, 2) * height) / 3
	return vol
}

// VolOfSphere returns the volume of a sphere
func VolOfSphere(radius float64) (vol float64) {
	vol = (4 * math.Pi * math.Pow(radius, 3)) / 3
	return vol
}

// VolOfCone return volume of a cone
func VolOfCone(radius, height float64) (vol float64) {
	vol = (math.Pi * math.Pow(radius, 2) * height) / 3
	return vol
}

// VolOfCylinder return volume of cylinder
func VolOfCylinder(radius, height float64) (vol float64) {
	vol = radius * height
	return vol
}

// VolOfRectPrism return volume of a rectangular prism
func VolOfRectPrism(length, width, height float64) (vol float64) {
	vol = length * width * height
	return vol
}

// VolOfTriPrism return volume of triangular prism
// with an equilateral base.
func VolOfTriPrism(baseSideA, baseSideB, height float64) (vol float64) {
	areaOfBase := baseSideA * baseSideB
	vol = areaOfBase * height
	return vol
}

// VolOfPentPrism return volume of pentagonal prism
func VolOfPentPrism(baseEdge, height float64) (vol float64) {
	factor := math.Sqrt(5*(5+2*math.Sqrt(5))) / 4
	vol = factor * math.Pow(baseEdge, 2) * height
	return vol
}

// VolOfHexaPrism return volume of hexagonal prism
func VolOfHexaPrism(baseEdge, height float64) (vol float64) {
	vol = (3 * math.Sqrt(3)) / 2 * math.Pow(baseEdge, 2) * height
	return vol
}

// VolOfOctPrism return volume of octagonal prism
func VolOfOctPrism(baseEdge, height float64) (vol float64) {
	factor := 2 * (1 + math.Sqrt(2))
	vol = factor * math.Pow(baseEdge, 2) * height
	return vol
}

// VolOfSquarePymd return volume of square pyramid
func VolOfSquarePymd(baseEdge, height float64) (vol float64) {
	vol = math.Pow(baseEdge, 2) * height / 3
	return vol
}

// VolOfTriPymd return volume of a triangular pyramid
func VolOfTriPymd(areaOfBase, height float64) (vol float64) {
	vol = (areaOfBase * height) / 3
	return vol
}

// VolOfPentPymd return volume of pentagonal pyramid
func VolOfPentPymd(baseEdge, height float64) (vol float64) {
	vol = (5 * math.Tan(54.0) * height * math.Pow(baseEdge, 2)) / 12
	return vol
}

// VolOfHexaPymd return volume of hexagonal pyramid
func VolOfHexaPymd(baseEdge, height float64) (vol float64) {
	vol = (math.Sqrt(3) * math.Pow(baseEdge, 2) * height) / 2
	return vol
}

// VolOfConeFrustum return volume of cone frustrum
func VolOfConeFrustum(radius, height float64) (vol float64) {
	vol = math.Pi * math.Pow(radius, 2) * height / 3
	return vol
}

// VolOfCyFrustum return volume of a cylinder frustum or right cylinder
func VolOfCyFrustum(radius, height float64) (vol float64) {
	vol = math.Pi * math.Pow(radius, 2) * height
	return vol
}

// VolOfEllipsoid return volume of ellipsoid
func VolOfEllipsoid(semiAxisA, semiAxisB, semiAxisC float64) (vol float64) {
	vol = (4 * math.Pi * semiAxisA * semiAxisB * semiAxisC) / 3
	return vol
}

// VolOfTorus return volume of a torus
func VolOfTorus(majorRadius, minorRadius float64) (vol float64) {
	vol = math.Pi * math.Pow(minorRadius, 2) * 2 * math.Pi * majorRadius
	return vol
}

// VolOfDecahedron return volume of a decahedron
func VolOfDecahedron(sides float64) (vol float64) {
	factor := (15 + (7 * math.Sqrt(5.0))) / 4
	vol = factor * math.Pow(sides, 3)
	return vol
}

// VolOfIcosahedron return volume of an icosahedron
func VolOfIcosahedron(side float64) (vol float64) {
	factor := 5 * (3.0 + math.Sqrt(5.0)) / 12
	vol = factor * math.Pow(side, 3)
	return vol
}
