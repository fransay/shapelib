// Package functs provide standalone functions for geometric related usage

package functs

import "math"

// Plane figures | 2D
// Area of a triangle

// AreaOfTriangle represent computation of area
// using the base and height properties of the triangle

func AreaOfTriangle(base, height float64) (area float64) {
	area = base * height
	return area
}

// AreaOfTriangleSide represent computation of area
// using the length properties of the triangle
func AreaOfTriangleSide(side1, side2, side3 float64) (area float64) {
	averageSides := (side1 + side2 + side3) / 3
	area = math.Sqrt(averageSides * (averageSides - side1) * (averageSides - side2) *
		(averageSides - side3))
	return area

}

// AreaOfTriangleL represents computation of area
// using the angle properties of the triangle
// angle described must be opposite to side3 of the triangle
func AreaOfTriangleAngle(side1, side2, angle float64) (area float64) {
	area = (side1 * side2) * math.Sin(angle) / 2
	return area

}

// Figures with four sides , aka quadilaterals
// Area of a square
func AreaOfSquare(length float64) (area float64) {
	area = length * length
	return area
}

// Area of a rectangle
func AreaOfRectangle(length float64) (area float64) {
	area = length * length
	return area
}

// Area of a parallelogram
func AreaOfParallelogram(base, height float64) (area float64) {
	area = base * height
	return area
}

// Area of a kite
// function arguments are the two diagonals in a kite
func AreaOfKite(diag1, diag2 float64) (area float64) {
	area = (diag1 + diag2) / 2
	return area
}

// Area of a rhombus
// func arguments are the two diagonals of the rhombus
// area algo same as kite
func AreaOfRhombus(diag1, diag2 float64) (area float64) {
	area = AreaOfKite(diag1, diag2)
	return area
}

// Area of a trapezoid
func AreaOfTrapezoid(side1, side2, height float64) (area float64) {
	area = 0.5 * (side1 + side2) * height
	return area
}

// Figures with 5 sides
// AreaOfRPentagon returns the area of a regular polygon
// R in the midst of AreaOfRPentagon represents regular
func AreaOfRPentagon(length float64) (area float64) {
	area = 0.25 * (math.Sqrt(5 * (5 + (2 * math.Sqrt(5))) * (length * length)))
	return area

}

// Figures with 6 sides
// AreaOfRHexagon returns the area of a regular Hexagon
// R in the midst of AreaOfRHexagon represents regular
func AreaOfRHexagon() {

}

// Figures with non-straight sides
// Area of a circle
func AreaOfCircle(radius float64) (area float64) {
	const pi float64 = 22 / 7
	area = radius * radius * pi
	return area

}

// Area of a semi circle
func AreaOfSemiCircle(radius float64) (area float64) {
	area = 0.5 * AreaOfCircle(radius)
	return area
}

// Area of an oval
func AreaOfOval() {

}
