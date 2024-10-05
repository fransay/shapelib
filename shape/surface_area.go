package utils

import (
	"math"
)

// SurfAreaOfTetrahedron returns surface area of a tetrahedron
func SurfAreaOfTetrahedron(sideLength float64) (surfArea float64) {
	surfArea = math.Sqrt(3) * math.Pow(sideLength, 2)
	return surfArea
}

// SurfAreaOfRightAngledPyramid returns surface area of a right-angled pyramid
// pyramidHeight is also known as slant height.
func SurfAreaOfRightAngledPyramid(baseArea float64, basePerimeter float64, pyramidHeight float64) (surfArea float64) {
	surfArea = baseArea + basePerimeter*pyramidHeight/2
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
func SurfAreaOfPentahedron(baseArea, basePerimeter, slantHeight float64) (sideArea float64) {
	sideArea = baseArea + basePerimeter*slantHeight/2
	return sideArea

}

// SurfAreaOfHexahedron returns the surface area of hexahedron
func SurfAreaOfHexahedron(sideLength float64) (surfArea float64) {
	surfArea = 6 * (sideLength * sideLength)
	return surfArea
}

// SurfAreaOfHeptahedron returns the surface area of heptahedron
func SurfAreaOfHeptahedron(apothem, slantHeight float64) (surfArea float64) {
	surfArea = (7 * apothem * slantHeight) / 4
	return surfArea
}

// SurfAreaOfRightAngledCylinder returns the surface area of a right cylinder
func SurfAreaOfRightAngledCylinder(radius, height float64) (surfArea float64) {
	surfArea = 2*math.Pi*radius*height + 2*math.Pi*math.Pow(radius, 2)
	return surfArea
}

// SurfAreaOfSphere returns the surface area of a sphere
func SurfAreaOfSphere(radius float64) (surfArea float64) {
	surfArea = 4 * math.Pi * math.Pow(radius, 2)
	return surfArea
}

// SurfAreaOfTorus returns the surface area of a Torus
func SurfAreaOfTorus(innerRadius, outerRadius float64) (surfArea float64) {
	surfArea = 4 * math.Pow(math.Pi, 2) * outerRadius * innerRadius
	return surfArea
}
