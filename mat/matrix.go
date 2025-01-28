package mat

// Mat defines a matrix type for n
type Mat2D struct {
	row int // number of row of matrix
	col int // number of columns of designated matrix
}

func NewMat(row, col int) *Mat2D {
	return &Mat2D{row: row, col: col}
}
