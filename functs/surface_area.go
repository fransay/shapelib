// Package functs computes the surface area of 3D geometries
package functs

import (
	"math"
)

// SurfAreaOfTetrahedron returns surface area of a tetrahedron
func SurfAreaOfTetrahedron(sideLength float64) (surfArea float64) {
	surfArea = math.Sqrt(3) * sideLength
	return surfArea

}

// SurfAreaOfRightAngledPyramid returns surface area of a right-angled pyramid
func SurfAreaOfRightAngledPyramid(baseLength float64, baseWidth float64, pyramidHeight float64) (surfArea float64) {
	surfArea = (baseLength * baseWidth) + baseLength*(math.Sqrt(((baseWidth/2)*(baseWidth/2))+(pyramidHeight*pyramidHeight))) +
		baseWidth*(math.Sqrt(((baseLength/2)*(baseLength/2))+(pyramidHeight*pyramidHeight)))
	return surfArea

}

// SurfAreaOfCube returns surface area of a cube
func SurfAreaOfCube(sideLength float64) (surfArea float64) {
	surfArea = 6 * sideLength
	return surfArea
}

// SurfAreaOfCuboid returns surface area of cuboid
func SurfAreaOfCuboid(sideLength, sideWidth, sideHeight float64) (surfArea float64) {
	surfArea = 2*(sideLength) + 2*(sideWidth) + 2*(sideHeight)
	return surfArea
}

// SurfAreaOfPentahedron returns the surface area of pentahedron
func SurfAreaOfPentahedron(sideLength, height float64) (sideArea float64) {
	sideArea = sideLength * (sideLength + math.Sqrt(math.Pow(sideLength, 2.0)+4*(math.Pow(height, 2.0))))
	return sideArea

}

// SurfAreaOfHexahedron returns the surface area of hexahedron
func SurfAreaOfHexahedron(sideLength float64) (surfArea float64) {
	surfArea = 6 * (sideLength * sideLength)
	return surfArea
}

// SurfAreaOfHeptahedron returns the surface area of heptahedron
func SurfAreaOfHeptahedron(sideLength, thickness float64) (surfArea float64) {
	surfArea = 2*(AreaOfRHeptagon(sideLength)) + 2*(AreaOfRectangle(sideLength, thickness))
	return surfArea

}

// SurfAreaOfRightCylinder returns the surface area of a right cylinder
func SurfAreaOfRightCylinder(radius, height float64) (surfArea float64) {
	surfArea = 2*math.Pi*radius*height + 2*math.Pi*math.Pow(radius, 2)
	return surfArea
}

// SurfAreaOfSphere returns the surface area of a sphere
func SurfAreaOfSphere(radius float64) (surfArea float64) {
	surfArea = 4 * math.Pi * math.Pow(radius, 2)
	return surfArea
}
