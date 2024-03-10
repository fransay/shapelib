package utils

import "testing"

// test IsClose
func TestIsClose(t *testing.T) {
	res := IsClose(10.00, 45.00, 0.001)
	if res != false {
		t.Errorf("Got: %v, Expected: %v", res, false)
	}
}
