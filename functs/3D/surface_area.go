// Package functs computes the surface area of 3D geometries
package functs

import "math"

// SurfAreaOfTetrahedron :surface area of a tetrahedron
func SurfAreaOfTetrahedron() (surfArea float64) {

	return surfArea

}

// SurfAreaOfRightAngledPyramid :surface area of a right-angled pyramid
func SurfAreaOfRightAngledPyramid(baseLength float64, baseWidth float64, pyramidHeight float64) (surfArea float64) {
	surfArea = (baseLength * baseWidth) + baseLength*(math.Sqrt(((baseWidth/2)*(baseWidth/2))+(pyramidHeight*pyramidHeight))) +
		baseWidth*(math.Sqrt(((baseLength/2)*(baseLength/2))+(pyramidHeight*pyramidHeight)))
	return surfArea

}

// SurfAreaOfCube :surface area of a cube
func SurfAreaOfCube(sideLength float64) (surfArea float64) {
	surfArea = 6 * sideLength
	return surfArea
}

// SurfAreaOfCuboid :surface area of cuboid
func SurfAreaOfCuboid(sideLength, sideWidth, sideHeigth float64) (surfArea float64) {
	surfArea = 2*(sideLength) + 2*(sideWidth) + 2*(sideHeigth)
	return surfArea
}

// surface
