// Package functs computes the surface area of 3D geometries
package functs

// SurfAreaOfCube :surface area of a cube
func SurfAreaOfCube(sideLength float64) (surfArea float64) {
	surfArea = 6 * sideLength
	return surfArea
}
