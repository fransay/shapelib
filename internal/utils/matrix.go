package utils

// MatrixDimension defines a dimensions of a matrix, as (row x column)
type MatrixDimension struct {
	row []int
	col []int
}

// NewMatrixDimension constructs new object of MatrixDimension
func NewMatrixDimension(row, col []int) *MatrixDimension {
	return &MatrixDimension{row: row, col: col}
}

// Matrix type with any dimension
type Matrix struct {
	dim     MatrixDimension
	matType any // typical of a number type (int, float64)
}

// NewMatrix returns new matrix with any dimensions of row and columns
func NewMatrix(dim MatrixDimension, matType any) *Matrix {
	return &Matrix{
		dim:     dim,
		matType: matType,
	}
}

// transpose
// conjugate transpose
// inverse
// pseudo-inverse
// determinant
// trace
// rank
// norms
// rotation
// scaling
// translation
// reflection
// shearing
// orthogonality
// adjudicate
// cofactor matrix
// compute eigenvalues
// compute eigenvectors
// matrix decomposition
// matrix multiplication
// matrix arithmetic operations
