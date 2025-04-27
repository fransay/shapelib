package shape

import (
	"testing"

	"github.com/fransay/shapelib/internal/utils"
)

func TestRectangle(t *testing.T) {
	var rectangle = Rectangle{Length: 15, Width: 20}
	resultOfAreaByLengthWidth := rectangle.AreaByLengthWidth()
	expectedAreaByLengthWidth := 300.0
	if !utils.IsClose(resultOfAreaByLengthWidth, expectedAreaByLengthWidth, 0.1) {
		t.Errorf("Got %f, Expected %f", resultOfAreaByLengthWidth, expectedAreaByLengthWidth)
	}

	// area by diagonal
	resultOfAreaByDiagonal := rectangle.AreaByDiagonal(25)
	expectedOfAreaByDiagonal := 3000.0
	if !utils.IsClose(resultOfAreaByDiagonal, resultOfAreaByDiagonal, 0.1) {
		t.Errorf("Got %f, Expected %f", resultOfAreaByDiagonal, expectedOfAreaByDiagonal)
	}

	// perimeter
	resultOfPerimeter := rectangle.Perimeter()
	expectedOfPerimeter := 70.0
	if !utils.IsClose(resultOfPerimeter, expectedOfPerimeter, 0.1) {
		t.Errorf("Got %f, Expected %f", resultOfPerimeter, expectedOfPerimeter)
	}

	// diagonal
	resultOfDiagonal := rectangle.Diagonal()
	expectedOfDiagonal := 28.284271
	if !utils.IsClose(resultOfDiagonal, expectedOfDiagonal, 0.1) {
		t.Errorf("Got %f, Expected %f", resultOfDiagonal, expectedOfDiagonal)
	}

}
