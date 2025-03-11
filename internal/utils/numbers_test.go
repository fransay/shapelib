package utils

import "testing"

// TestAddTwoInt tests the addition of two integers
func TestAddTwoInt(t *testing.T) {
	addTwoInts := AddTwoInt(1, 5)
	if addTwoInts != 6 {
		t.Errorf("addition operand not functional")
	}
}

// TestAddTwoFloat tests the addition of two integers
func TestAddTwoFloat(t *testing.T) {
	addTwoFloat := AddTwoFloat(1.0, 5.0)
	if addTwoFloat != 6.0 {
		t.Errorf("incorrect addition operation")
	}

}

// Test CompareThreeFloat64
func TestCompareTrio(t *testing.T) {
	res := CompareThreeFloat64(1, 2, 3)
	if res != false {
		t.Errorf("Got %v, Expected %v", res, false)
	}
}

// Test CompareTwoFloat64
func TestCompareDuo(t *testing.T) {
	res := CompareTwoFloat64(1, 1)
	if res != true {
		t.Errorf("Got %v, Expected %v", res, true)
	}
}

// Test IsClose
func TestIsClose(t *testing.T) {
	res := IsClose(10.00, 45.00, 0.001)
	if res != false {
		t.Errorf("Got: %v, Expected: %v", res, false)
	}

	res2 := IsClose(10.00, 10.00, 0.001)
	if res2 != true {
		t.Errorf("Got: %v, Expected: %v", res2, true)
	}
}

// Test Diff
func TestDiff(t *testing.T) {
	res := Diff(10.0, 15.0)
	if res != 5.0 {
		t.Errorf("Got: %v, Expected: %v", res, false)
	}
}
