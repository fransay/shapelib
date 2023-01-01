package functs

import (
	"math"
)

// VolOfTetrahedron :volume of a tetrahedron
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

// VolOfHexahedron volume of a hexahedron
func VolOfHexahedron() (vol float64) {
	return vol

}
