// Package functs: The perimeter of a 2D shape is the total distance around the respective shape
// Package functs: Package for computing the perimeter of REGULAR 2D shapes
package functs

// PERIMETER OF THREE SIDED FIGURES

// PerimOfRTriangle : perimeter of a regular triangle
func PerimOfRTriangle(sideA, sideB, sideC float64) (perim float64) {
	perim = sideA + sideB + sideC
	return perim
}

// PERIMETER OF FOUR SIDED FIGURES

// PerimOfRSquare : perimeter of a regular square
func PerimOfRSquare(side float64) (perim float64) {
	perim = side * 4
	return perim
}

// PerimOfRRectangle : perimeter of a regular rectangle
func PerimOfRRectangle(sideLength, sideWidth float64) (perim float64) {
	perim = (sideWidth * 2) + (sideLength * 2)
	return perim
}

// PerimOfRParallelogram :perimeter of a regular parallelogram
func PerimOfRParallelogram(northSide, westSide float64) (perim float64) {
	perim = (northSide * 2) + (westSide * 2)
	return perim

}

// PerimOfRKite :perimeter of a regular Kite
func PerimOfRKite(shortDiag, longDiag float64) (perim float64) {
	perim = (shortDiag * 2) + (longDiag * 2)
	return perim

}

// PerimOfRRhombus :perimeter of a regular Rhombus
func PerimOfRRhombus(side float64) (perim float64) {
	perim = side * 4
	return perim
}

// PerimOfRTrapezoid :a trapezoid has unequal sides
func PerimOfRTrapezoid(side1, side2, side3, side4 float64) (perim float64) {
	perim = side1 + side2 + side3 + side4
	return perim
}

// PERIMETER OF FIVE SIDED FIGURES

// PerimOfRPentagon :perimeter of regular Pentagon
func PerimOfRPentagon(side float64) (perim float64) {
	perim = side * 5
	return perim

}

// PERIMETER OF SIX SIDED FIGURES

// PerimOfRHexagon :perimeter of a regular Hexagon
func PerimOfRHexagon(side float64) (perim float64) {
	perim = side * 6
	return perim
}

// PERIMETER OF SEVEN SIDED FIGURES

// PerimOfRHeptagon :perimeter of a regular Heptagon
func PerimOfRHeptagon(side float64) (perim float64) {
	perim = side * 7
	return perim
}

// PERIMETER OF EIGHT SIDED FIGURES

// PerimOfROctagon :perimeter of a regular Octagon
func PerimOfROctagon(side float64) (perim float64) {
	perim = side * 8
	return perim
}

// PERIMETER OF NINE SIDED FIGURES

// PerimOfRNonagon  :perimeter of a regular Nonagon
func PerimOfRNonagon(side float64) (perim float64) {
	perim = side * 9
	return perim
}

// PERIMETER OF NINE SIDED FIGURES

// PerimOfRDecagon :perimeter of a regular Decagon
func PerimOfRDecagon(side float64) (perim float64) {
	perim = side * 10
	return perim
}

// PERIMETER OF GENERIC SHAPES

// PerimOfRegularGeneric :perimeter of a regular generic shape
func PerimOfRegularGeneric(sideLength float64, numberOfSides float64) (perim float64) {
	perim = sideLength * numberOfSides
	return perim
}

// PerimOfIrregularGeneric :perimeter of an irregular generic shape
func PerimOfIrregularGeneric(side ...float64) (perim float64) {
	var total float64
	for _, value := range side {
		total = total + value
	}
	return total

}
