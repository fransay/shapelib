package utils

// MaxtrixDimension defines a dimensions of a matrix, as (row x column)
type MatrixDimension struct {
	row []int,
	col []int,
}

func NewMatrixDimension struct {}

// Matrix type with any dimension
type Matrix struct {
	dim MatrixDimension
	matType any // typical of a number type (int, float64)
}

// NewMatrix returns new matrix with any dimensions of row and columns
func NewMatrix(row int, column int, matType any) *Matrix {
	return &Matrix{
		row:     row,
		col:     column,
		matType: matType,
	}
}

func 
