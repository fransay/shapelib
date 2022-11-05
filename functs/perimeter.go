// Package functs: The perimeter of a 2D shape is the total distance around the respective shape
// Package functs: Package for computing the perimeter of REGULAR 2D shapes
package functs

// PerimOfRTriangle : perimeter of a regular triangle
func PerimOfRTriangle(sideA, sideB, sideC float64) (perim float64) {
	if sideA != sideB { // equilateral triangles
		perim = sideA + sideB + sideC
	} else {
		perim = 3 * sideA
	}
	return perim
}

// Perimeter of four sided figures

// PerimOfRSquare : perimeter of a regular square
func PerimOfRSquare(side float64) (perim float64) {
	perim = side * 4
	return perim
}

// Perimeter of a regular rectangle

func PerimOfRRectangle(sideLength, sideWidth float64) (perim float64) {
	perim = (sideWidth * 2) + (sideLength * 2)
	return perim
}

// Perimeter of a regular parallelogram

func PerimOfRParallelogram(northSide, westSide float64) (perim float64) {
	perim = (northSide * 2) + (westSide * 2)
	return perim

}

// Perimeter of a regular Kite

func PerimOfRKite(shortDiag, longDiag float64) (perim float64) {
	perim = (shortDiag * 2) + (longDiag * 2)
	return perim

}
