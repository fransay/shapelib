package props

// PerimeterOfTriangle return perimeter of a regular triangle
func PerimeterOfTriangle(sideA, sideB, sideC float64) (perimeter float64) {
	perimeter = sideA + sideB + sideC
	return perimeter
}

// PerimeterOfSquare returns perimeter of a regular square
func PerimeterOfSquare(side float64) (perimeter float64) {
	perimeter = side * 4
	return perimeter
}

// PerimeterOfRectangle returns the perimeter of a regular rectangle
func PerimeterOfRectangle(sideLength, sideWidth float64) (perimeter float64) {
	perimeter = (sideWidth * 2) + (sideLength * 2)
	return perimeter
}

// PerimeterOfParallelogram returns perimeter of a regular parallelogram
func PerimeterOfParallelogram(northSide, westSide float64) (perimeter float64) {
	perimeter = (northSide * 2) + (westSide * 2)
	return perimeter

}

// PerimeterOfKite returns perimeter of a regular kite
func PerimeterOfKite(shortDiag, longDiag float64) (perimeter float64) {
	perimeter = (shortDiag * 2) + (longDiag * 2)
	return perimeter

}

// PerimeterOfRhombus returns perimeter of a regular rhombus
func PerimeterOfRhombus(side float64) (perimeter float64) {
	perimeter = side * 4
	return perimeter
}

// PerimeterOfTrapezoid returns perimeter of trapezoid with unequal sides
func PerimeterOfTrapezoid(side1, side2, side3, side4 float64) (perimeter float64) {
	perimeter = side1 + side2 + side3 + side4
	return perimeter
}

// PerimeterOfPentagon returns perimeter of regular pentagon
func PerimeterOfPentagon(side float64) (perimeter float64) {
	perimeter = side * 5
	return perimeter

}

// PerimeterOfHexagon returns perimeter of a regular hexagon
func PerimeterOfHexagon(side float64) (perimeter float64) {
	perimeter = side * 6
	return perimeter
}

// PerimeterOfHeptagon returns perimeter of a regular heptagon
func PerimeterOfHeptagon(side float64) (perimeter float64) {
	perimeter = side * 7
	return perimeter
}

// PerimeterOfOctagon returns perimeter of a regular octagon
func PerimeterOfOctagon(side float64) (perimeter float64) {
	perimeter = side * 8
	return perimeter
}

// PerimeterOfNonagon returns perimeter of a regular nonagon
func PerimeterOfNonagon(side float64) (perimeter float64) {
	perimeter = side * 9
	return perimeter
}

// PerimeterOfDecagon returns perimeter of a regular decagon
func PerimeterOfDecagon(side float64) (perimeter float64) {
	perimeter = side * 10
	return perimeter
}

// PerimeterOfGenericFigure returns perimeter of a regular generic shape
func PerimeterOfGenericFigure(sideLength float64, numberOfSides float64) (perimeter float64) {
	perimeter = sideLength * numberOfSides
	return perimeter
}

// PerimeterOfIrregularGeneric returns perimeter of an irregular generic shape
func PerimeterOfIrregularGeneric(side ...float64) (perimeter float64) {
	var total float64
	for _, value := range side {
		total = total + value
	}
	return total

}

// PerimeterOfIrregularFourSided returns the perimeter of an irregular four sides figure
func PerimeterOfIrregularFourSided(sideOne, sideTwo, sideThree, sideFour float64) (perimeter float64) {
	perimeter = sideOne + sideTwo + sideThree + sideFour
	return perimeter
}
