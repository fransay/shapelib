package functs

import "math"

func VolOfTetrahedron(side float64) (vol float64) {
	vol = (side * side * side) / (6 * math.Sqrt(2))
	return vol
}

func VolOfRightAngledTriPyramid(baseBaseTriangle, heightBaseTriangle, heightPyramid float64) (vol float64) {
	vol = (baseBaseTriangle * heightBaseTriangle * heightPyramid) / 6
	return vol
}
