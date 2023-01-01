package functs

import "math"

// VolOfTetrahedron returns the volume of a tetrahedron
func VolOfTetrahedron(side float64) (vol float64) {
	vol = (side * side * side) / (6 * math.Sqrt(2))
	return vol
}

// VolOfRightAngledTriPyramid returns the volume of a right-angled triangular pyramid
func VolOfRightAngledTriPyramid(baseBaseTriangle, heightBaseTriangle, heightPyramid float64) (vol float64) {
	vol = (baseBaseTriangle * heightBaseTriangle * heightPyramid) / 6
	return vol
}
