// Package functs computes the surface area of 3D geometries
package functs

import (
	"math"
)

// SurfAreaOfTetrahedron :surface area of a tetrahedron
func SurfAreaOfTetrahedron(sideLength float64) (surfArea float64) {
	surfArea = math.Sqrt(3) * sideLength
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

// SurfAreaOfPentahedron :surface area of pentahedron
func SurfAreaOfPentahedron(sideLength, height float64) (sideArea float64) {
	sideArea = sideLength * (sideLength + math.Sqrt(math.Pow(sideLength, 2.0)+4*(math.Pow(height, 2.0))))
	return sideArea

}

// SurfAreaOfHexahedron :surface area of hexahedron
func SurfAreaOfHexahedron(sideLength float64) (surfArea float64) {
	surfArea = 6 * (sideLength * sideLength)
	return surfArea
}

// SurfAreaOfHeptahedron surface area of heptahedron
func SurfAreaOfHeptahedron(sideLength, thickness float64) (surfArea float64) {
	surfArea = 2*(AreaOfRHeptagon(sideLength)) + 2*(AreaOfRectangle(sideLength, thickness))
	return surfArea

}

// SurfAreaOfRightCylinder :surface area of a right cylinder
func SurfAreaOfRightCylinder(radius, height float64) (surfArea float64) {
	surfArea = 2*math.Pi*radius*height + 2*math.Pi*math.Pow(radius, 2)
	return surfArea
}
