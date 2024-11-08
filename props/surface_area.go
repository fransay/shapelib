package props

import (
	"math"
)

// SurfAreaOfTetrahedron returns surface area of a tetrahedron
func SurfAreaOfTetrahedron(sideLength float64) float64 {
	return math.Sqrt(3) * math.Pow(sideLength, 2)
}

// SurfAreaOfRightAngledPyramid returns surface area of a right-angled pyramid
func SurfAreaOfRightAngledPyramid(
	baseArea float64,
	basePerimeter float64,
	pyramidHeight float64) float64 {
	return baseArea + basePerimeter*pyramidHeight/2
}

// SurfAreaOfCube returns surface area of a cube
func SurfAreaOfCube(sideLength float64) float64 {
	return 6 * sideLength
}

// SurfAreaOfCuboid returns surface area of cuboid
func SurfAreaOfCuboid(
	sideLength,
	sideWidth,
	sideHeight float64) float64 {
	return 2*(sideLength) + 2*(sideWidth) + 2*(sideHeight)
}

// SurfAreaOfPentahedron returns the surface area of pentahedron
func SurfAreaOfPentahedron(
	baseArea,
	basePerimeter,
	slantHeight float64) float64 {
	return baseArea + basePerimeter*slantHeight/2
}

// SurfAreaOfHexahedron returns the surface area of hexahedron
func SurfAreaOfHexahedron(sideLength float64) float64 {
	return 6 * (sideLength * sideLength)
}

// SurfAreaOfHeptahedron returns the surface area of heptahedron
func SurfAreaOfHeptahedron(
	apothem,
	slantHeight float64) float64 {
	return (7 * apothem * slantHeight) / 4
}

// SurfAreaOfRightAngledCylinder returns the surface area of a right cylinder
func SurfAreaOfRightAngledCylinder(radius, height float64) float64 {
	return 2*math.Pi*radius*height + 2*math.Pi*math.Pow(radius, 2)
}

// SurfAreaOfSphere returns the surface area of a sphere
func SurfAreaOfSphere(radius float64) float64 {
	return 4 * math.Pi * math.Pow(radius, 2)
}

// SurfAreaOfTorus returns the surface area of a Torus
func SurfAreaOfTorus(innerRadius, outerRadius float64) float64 {
	return 4 * math.Pow(math.Pi, 2) * outerRadius * innerRadius
}
