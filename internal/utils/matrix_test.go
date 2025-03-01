package utils

import (
	"fmt"
	"testing"
)

func TestMat(t *testing.T) {
	mat1 := NewMatrix(2, 3, float64(1.0))
	// test methods on matrices.
	fmt.Println(mat1)
}
