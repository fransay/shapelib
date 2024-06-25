package tests_test

import (
	"shapelib/shape"
	"testing"
)

func TestRectangle(t *testing.T) {
	var rectangle = shape.Rectangle{Length: 15, Width: 20}
	// area by length width
	resultOfAreaByLengthWidth := rectangle.AreaByLengthWidth()
	expectedAreaByLengthWidth := 3000.0
	if resultOfAreaByLengthWidth != expectedAreaByLengthWidth {
		t.Errorf("Got %f, Expected %f", resultOfAreaByLengthWidth, expectedAreaByLengthWidth)
	}

	// area by diagonal
	resultOfAreaByDiagonal := rectangle.AreaByDiagonal(25)
	expectedOfAreaByDiagonal := 3000.0
	if resultOfAreaByDiagonal != expectedOfAreaByDiagonal {
		t.Errorf("Got %f, Expected %f", resultOfAreaByDiagonal, expectedOfAreaByDiagonal)
	}

	// perimeter
	resultOfPerimeter := rectangle.Perimeter()
	expectedOfPerimeter := 70.0
	if resultOfAreaByDiagonal != expectedOfAreaByDiagonal {
		t.Errorf("Got %f, Expected %f", resultOfPerimeter, expectedOfPerimeter)
	}

	// diagonal
	resultOfDiagonal := rectangle.Diagonal()
	expectedOfDiagonal := 70.0
	if resultOfDiagonal != expectedOfDiagonal {
		t.Errorf("Got %f, Expected %f", expectedOfDiagonal, expectedOfDiagonal)
	}

}
