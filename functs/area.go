// Package functs provide basic functions for geometric
// related problems

package functs

import "math"

// 2D figures

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
	area = math.Sqrt(averageSides * (averageSides-side1) * (averageSides-side2)
	(averageSides-side3))
	return area

}

// AreaOfTriangleL represents computation of area
// using the angle properties of the triangle
// angle described must be opposite to side3 of the triangle
func AreaOfTriangleAngle(side1, side2 , angle float64) (area float64) {
	area = (side1 * side2)


}

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


