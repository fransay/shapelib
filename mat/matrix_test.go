package mat

import (
	"fmt"
	"testing"
)

func TestMat(t *testing.T) {
	mat1 := NewMat(2, 3)
	// test methods on matrices.
	fmt.Println(mat1)
}
