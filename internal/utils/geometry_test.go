package utils

// Test Distance
func TestDistance(t *testing.T) {
	res := Distance([]float64{4, 6}, []float64{8, 2})
	exp := 2.8284271247461903
	if res != exp {
		t.Errorf("Got %v, Expect %v", res, exp)
	}
}
