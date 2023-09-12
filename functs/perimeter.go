// Package functs: The perimeter of a 2D shape is the total distance around the respective shape
// Package functs: Package for computing the perimeter of REGULAR 2D shapes
package functs

// PERIMETER OF THREE SIDED FIGURES

// PerimeterOfTriangle : perimeter of a regular triangle
func PerimeterOfTriangle(sideA, sideB, sideC float64) (perim float64) {
	perim = sideA + sideB + sideC
	return perim
}

// PERIMETER OF FOUR SIDED FIGURES

// PerimeterOfSquare : perimeter of a regular square
func PerimeterOfSquare(side float64) (perimeter float64) {
	perimeter = side * 4
	return perimeter
}

// PerimeterOfRegularRectangle returns the perimeter of a regular rectangle
func PerimeterOfRegularRectangle(sideLength, sideWidth float64) (perimeter float64) {
	perimeter = (sideWidth * 2) + (sideLength * 2)
	return perimeter
}

// PerimeterOfParallelogram :perimeter of a regular parallelogram
func PerimeterOfParallelogram(northSide, westSide float64) (perim float64) {
	perim = (northSide * 2) + (westSide * 2)
	return perim

}

// PerimeterOfKite returns the perimeter of a kite
func PerimeterOfKite(shortDiag, longDiag float64) (perimeter float64) {
	perimeter = (shortDiag * 2) + (longDiag * 2)
	return perimeter

}

// PerimeterOfRhombus returns the perimeter of a rhombus
func PerimeterOfRhombus(side float64) (perimeter float64) {
	perimeter = side * 4
	return perimeter
}

// PerimeterOfTrapezoid :a trapezoid has unequal sides
func PerimeterOfTrapezoid(side1, side2, side3, side4 float64) (perimeter float64) {
	perimeter = side1 + side2 + side3 + side4
	return perimeter
}

// PERIMETER OF FIVE SIDED FIGURES

// PerimeterOfPentagon returns the perimeter of a pentagon
func PerimeterOfPentagon(side float64) (perimeter float64) {
	perimeter = side * 5
	return perimeter

}

// PERIMETER OF SIX SIDED FIGURES

// PerimeterOfHexagon :perimeter of a regular Hexagon
func PerimeterOfHexagon(side float64) (perimeter float64) {
	perimeter = side * 6
	return perimeter
}

// PERIMETER OF SEVEN SIDED FIGURES

// PerimeterOfHeptagon :perimeter of a regular Heptagon
func PerimeterOfHeptagon(side float64) (perimeter float64) {
	perimeter = side * 7
	return perimeter
}

// PERIMETER OF EIGHT SIDED FIGURES

// PerimeterOfOctagon returns the perimeter of an octagon
func PerimeterOfOctagon(side float64) (perimeter float64) {
	perimeter = side * 8
	return perimeter
}

// PERIMETER OF NINE SIDED FIGURES

// PerimeterOfNonagon returns the perimeter of a nonagon
func PerimeterOfNonagon(side float64) (perimeter float64) {
	perimeter = side * 9
	return perimeter
}

// PERIMETER OF NINE SIDED FIGURES

// PerimeterOfDecagon returns the perimeter of a decagon
func PerimeterOfDecagon(side float64) (perimeter float64) {
	perimeter = side * 10
	return perimeter
}

// PERIMETER OF GENERIC SHAPES

// PerimeterOfShape returns the perimeter of any regular shape given the number of sides
func PerimeterOfShape(sideLength float64, numberOfSides float64) (perimeter float64) {
	perimeter = sideLength * numberOfSides
	return perimeter
}

// PerimeterOfIrregularShape returns the perimeter of a shape given the sides
func PerimeterOfIrregularShape(side ...float64) (perimeter float64) {
	var total float64
	for _, value := range side {
		total = total + value

	}
	return total

}

// PerimeterOfIrregularFourSidedFigures returns the perimeter of an irregular rectangle
func PerimeterOfIrregularFourSidedFigures(sideOne, sideTwo, sideThree, sideFour float64) (perimeter float64) {
	perimeter = sideOne + sideTwo + sideThree + sideFour
	return perimeter
}
