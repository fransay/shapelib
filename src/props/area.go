package props

import (
	"math"

	"github.com/fransay/shapelib/src/geom"
)

// AreaOfTriangleGivenBH returns area of a triangle given the base and the height
func AreaOfTriangleGivenBH(base, height float64) (area float64) {
	area = (base * height) / 2
	return area
}

// AreaOfTriangleGivenSide returns the area of a triangle given all three sides
func AreaOfTriangleGivenSide(side1, side2, side3 float64) (area float64) {
	semiPerimeter := (side1 + side2 + side3) / 2
	area = math.Sqrt(semiPerimeter * (semiPerimeter - side1) * (semiPerimeter - side2) * (semiPerimeter - side3))
	return area
}

// AreaOfTriangleGivenAngle returns the area of a triangle given two sides and an angle
func AreaOfTriangleGivenAngle(side1, side2, angle float64) (area float64) {
	area = ((side1 * side2) * math.Sin(angle)) / 2
	return area

}

// AreaOfSquare return the area of square given the length
func AreaOfSquare(length float64) (area float64) {
	area = length * length
	return area
}

// AreaOfRectangle returns the area of a rectangle given the length and width
func AreaOfRectangle(length, width float64) (area float64) {
	area = length * width
	return area
}

// AreaOfParallelogram returns the area of a parallelogram given the base and height
func AreaOfParallelogram(base, height float64) (area float64) {
	area = base * height
	return area
}

// AreaOfKite returns the area of a kite given the main diagonals
func AreaOfKite(diag1, diag2 float64) (area float64) {
	area = (diag1 + diag2) / 2
	return area
}

// AreaOfRhombus returns the area of a rhombus given the main diagonals
func AreaOfRhombus(diag1, diag2 float64) (area float64) {
	area = (diag1 + diag2) / 2
	return area
}

// AreaOfTrapezoid returns the area trapezoid given the sides and the height
func AreaOfTrapezoid(side1, side2, height float64) (area float64) {
	area = 0.5 * (side1 + side2) * height
	return area
}

// AreaOfPentagon returns the area of a regular pentagon given the side
func AreaOfPentagon(side float64) (area float64) {
	area = 0.25 * math.Sqrt(5.0*(5.0+(2.0*math.Sqrt(5.0)))) * (side * side)
	return area

}

// AreaOfHexagon returns the area of a regular hexagon given the side
func AreaOfHexagon(side float64) (area float64) {
	area = 1.5 * math.Sqrt(3) * math.Pow(side, 2)
	return area

}

// AreaOfHeptagon returns the area of a regular heptagon given the side
func AreaOfHeptagon(side float64) (area float64) {
	const apothem float64 = 3.634
	area = apothem * side * side
	return area
}

// AreaOfOctagon returns the area of a regular octagon given the side
func AreaOfOctagon(side float64) (area float64) {
	area = 2 * (1 + math.Sqrt(2.0)) * (side * side)
	return area
}

// AreaOfNonagon returns the area of a regular nonagon given the side
func AreaOfNonagon(side float64) (area float64) {
	const cotangent = 2.747477
	area = (9 / 4) * math.Pow(side, 2.0) * cotangent
	return area

}

// AreaOfDecagon returns the area of a regular decagon given the side
func AreaOfDecagon(side float64) (area float64) {
	area = 2.5 * math.Pow(side, 2.0) * math.Sqrt(5.0+2*math.Sqrt(5.0))
	return area
}

// AreaOfDodecagon returns the area of a regular decagon given the sides
func AreaOfDodecagon(side float64) (area float64) {
	return (3 * math.Sqrt(3) * math.Pow(side, 2)) / 2
}

// AreaOfCircle returns the area of a circle given its radius
func AreaOfCircle(radius float64) (area float64) {
	const pi float64 = 22 / 7
	area = radius * radius * pi
	return area
}

// AreaOfSemiCircle returns the area of a semicircle given the radius
func AreaOfSemiCircle(radius float64) (area float64) {
	area = 0.5 * AreaOfCircle(radius)
	return area
}

// AreaOfQuadCircle returns the area of a quad-circle
func AreaOfQuadCircle(radius float64) (area float64) {
	area = AreaOfCircle(radius) / 4
	return area
}

// AreaOfOval returns the area of an oval
func AreaOfOval(semiMajorAxis float64, semiMinorAxis float64) (area float64) {
	area = math.Pi * semiMajorAxis * semiMinorAxis
	return area
}

// AreaOfEllipse returns the area of an ellipse
func AreaOfEllipse(semiMajorAxis, semiMinorAxis float64) (area float64) {
	area = AreaOfOval(semiMajorAxis, semiMinorAxis)
	return area
}

// AreaGivenCoordinates returns the area of any two-dimensional shape given the coordinates of the vertices, using the shoelace algorithm...
func AreaGivenCoordinates(cords ...geom.Point2D) float64 {
	var forPass float64
	var backPass float64
	var numberOfCords = len(cords) - 1
	for i, j := 0, 1; j <= numberOfCords; i, j = i+1, j+1 { // todo: fix area given coordinates
		forPass += cords[i].Y * cords[j].X
		backPass += cords[j].Y * cords[i].X
	}
	var area = -1 * (forPass - backPass) / 2 // results of area are negated due to orientation of types.Point2D
	return area
}

// todo: include area of irregular shapes
