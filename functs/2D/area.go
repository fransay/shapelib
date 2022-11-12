// Package functs provide standalone functions for geometric related usage

package functs

import "math"

type Point struct {
	X float64
	Y float64
}

// Plane figures | 2D
// Area of triangle

// AreaOfTriangleBH :returns the results of area of a triangle
// given the base and height quantities of the triangle
func AreaOfTriangleBH(base, height float64) (area float64) {
	area = 0.5 * base * height
	return area
}

// AreaOfTriangleSide :returns the results of area of a triangle
// given the length quantity of the triangle
func AreaOfTriangleSide(side1, side2, side3 float64) (area float64) {
	averageSides := (side1 + side2 + side3) / 3
	area = math.Sqrt(averageSides * (averageSides - side1) * (averageSides - side2) *
		(averageSides - side3))
	return area

}

// AreaOfTriangleAngle :returns the results of area of a triangle
// given the angle properties of the triangle
// angle described must be opposite to side3 of the triangle as defined locally
func AreaOfTriangleAngle(side1, side2, angle float64) (area float64) {
	area = (side1 * side2) * math.Sin(angle) / 2
	return area

}

// Area of quadrilateral
// AreaOfSquare :return the area of square
// given the length property
func AreaOfSquare(length float64) (area float64) {
	area = length * length
	return area
}

// AreaOfRectangle :return the area of rectangle
// given the length and width properties
func AreaOfRectangle(length, width float64) (area float64) {
	area = length * width
	return area
}

// AreaOfParallelogram :return the area of parallelogram
// given the base and height properties
func AreaOfParallelogram(base, height float64) (area float64) {
	area = base * height
	return area
}

// AreaOfKite :return the area of kite
// given the diagonal properties
func AreaOfKite(diag1, diag2 float64) (area float64) {
	area = (diag1 + diag2) / 2
	return area
}

// AreaOfRhombus :return the area of rhombus
// given the diagonal properties
func AreaOfRhombus(diag1, diag2 float64) (area float64) {
	area = AreaOfKite(diag1, diag2)
	return area
}

// AreaOfTrapezoid :returns the area trapezoid
// given the sides and height properties
func AreaOfTrapezoid(side1, side2, height float64) (area float64) {
	area = 0.5 * (side1 + side2) * height
	return area
}

// Area of pentagon

// AreaOfRPentagon :returns the area of a regular polygon
// R in the midst of AreaOfRPentagon represents regular
// given the side length properties of the pentagon
func AreaOfRPentagon(length float64) (area float64) {
	area = 0.25 * (math.Sqrt(5 * (5 + (2 * math.Sqrt(5))) * (length * length)))
	return area

}

// AreaOfRIPentagon :returns the area of an irrgular polygon
// I in the midst of AreaOfRIPentagon represents irregular
// given the cordinates properties inf the nodes/stations of an irregular pentagon
func AreaOfIPentagon(cords ...Point) float64 {
	area := AreaC(cords)
	return area

}

// Area of hexagon

// AreaOfRHexagon :returns the area of a regular Hexagon
// R in the midst of AreaOfRHexagon represents regular
func AreaOfRHexagon(side float64) (area float64) {
	area = (3 / 2) * math.Sqrt(3) * math.Pow(side, 2)
	return area

}

// AreaOfIHexagon : return the area of irregular hexagon
// I in the midst of AreaOfRIPentagon represents irregular
// Given the coordinates
func AreaOfIHexagon(cords ...Point) (area float64) {
	return AreaC(cords[1])
}

// AreaOfRHeptagon :returns the area of a regular heptagon
func AreaOfRHeptagon(side float64) (area float64) {
	const apothem float64 = 3.634
	area = apothem * side * side
	return area
}

// AreaOfCircle :Figures with non-straight sides
// Area of a circle
func AreaOfCircle(radius float64) (area float64) {
	const pi float64 = 22 / 7
	area = radius * radius * pi
	return area
}

// AreaOfSemiCircle :Area of a semi circle
// 1/2 the area of full circle ~ area of a semi circle
func AreaOfSemiCircle(radius float64) (area float64) {
	area = 0.5 * AreaOfCircle(radius)
	return area
}

// AreaOfOval :Area of an oval
func AreaOfOval(semiMajorAxis float64, semiMinorAxis float64) (area float64) {
	area = math.Pi * semiMajorAxis * semiMajorAxis
	return area
}

// AreaOfEllipse :Area of an oval
// An Ellipse is the same as an oval but not all ovals are ellipse (Unidirectional casting)
func AreaOfEllipse(semiMajorAxis, semiMinorAxis float64) (area float64) {
	area = AreaOfOval(semiMajorAxis, semiMinorAxis)
	return area
}

// AreaCord :returns the area of any 2D shape
// given the coordinates of the shape.
func AreaCord(cords ...Point) float64 {
	// employ shoelace algorithm
	var forPass float64
	var backPass float64
	var numberOfCords = len(cords)
	for i, j := 0, 1; j <= numberOfCords; i, j = i+1, j+1 {
		forPass += cords[i].Y * cords[j].X
		backPass += cords[j].Y * cords[i].X
	}
	area := (forPass - backPass) / 2
	return area
}
