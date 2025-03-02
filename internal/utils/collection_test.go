package utils 

// Test EqualSlice
func TestEqualSlice(t *testing.T) {
	var sliceOne = []float64{1, 2, 3, 4}
	var sliceTwo = []float64{1, 2, 3, 4}
	res := EqualSlice(sliceOne, sliceTwo)
	if res != true {
		t.Errorf("Got %v, Expected %v", res, true)
	}
}

// Test IsArrayElementFloat64
func TestIsArrayElementFloat64(t *testing.T) {
	arr := []float64{1.0, 4.0, 5.0}
	observed := IsArrayElementFloat64(arr)
	if observed != true {
		t.Errorf("Got %v, Expected %v", observed, true)
	}
}

