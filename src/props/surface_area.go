package props

import (
	"math"
)

// SurfaceAreaOfTetrahedron returns surface area of a tetrahedron
func SurfaceAreaOfTetrahedron(sideLength float64) float64 {
	return math.Sqrt(3) * math.Pow(sideLength, 2)
}

// SurfaceAreaOfRightAngledPyramid returns surface area of a right-angled pyramid
func SurfaceAreaOfRightAngledPyramid(baseArea float64, basePerimeter float64, pyramidHeight float64) float64 {
	return baseArea + basePerimeter*pyramidHeight/2
}

// SurfaceAreaOfCube returns surface area of a cube
func SurfaceAreaOfCube(sideLength float64) float64 {
	return 6 * sideLength
}

// SurfaceAreaOfCuboid returns surface area of cuboid
func SurfaceAreaOfCuboid(sideLength, sideWidth, sideHeight float64) float64 {
	return 2*(sideLength) + 2*(sideWidth) + 2*(sideHeight)
}

// SurfaceAreaOfPentahedron returns the surface area of pentahedron
func SurfaceAreaOfPentahedron(baseArea, basePerimeter, slantHeight float64) float64 {
	return baseArea + basePerimeter*slantHeight/2
}

// SurfaceAreaOfHexahedron returns the surface area of hexahedron
func SurfaceAreaOfHexahedron(sideLength float64) float64 {
	return 6 * (sideLength * sideLength)
}

// SurfaceAreaOfHeptahedron returns the surface area of heptahedron
func SurfaceAreaOfHeptahedron(apothem, slantHeight float64) float64 {
	return (7 * apothem * slantHeight) / 4
}

// SurfaceAreaOfRightAngledCylinder returns the surface area of a right cylinder
func SurfaceAreaOfRightAngledCylinder(radius, height float64) float64 {
	return 2*math.Pi*radius*height + 2*math.Pi*math.Pow(radius, 2)
}

// SurfaceAreaOfSphere returns the surface area of a sphere
func SurfaceAreaOfSphere(radius float64) float64 {
	return 4 * math.Pi * math.Pow(radius, 2)
}

// SurfaceAreaOfTorus returns the surface area of a Torus
func SurfaceAreaOfTorus(innerRadius, outerRadius float64) float64 {
	return 4 * math.Pow(math.Pi, 2) * outerRadius * innerRadius
}
