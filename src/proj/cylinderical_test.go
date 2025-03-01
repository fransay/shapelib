package proj

import (
	"testing"
)

func TestCylinderical(t *testing.T) {
	cylindericalObject := Cylinderical{Distance: 60, Angle: 20, Height: 300}

	// test toCartesian
	observedToCartesian := cylindericalObject.ToCartesian()
	expectedToCartesian := Cart3D{X: 5000, Y: 4000} // placeholder input
	if observedToCartesian != expectedToCartesian {
		t.Errorf("Expected %v, Got %v", expectedToCartesian, observedToCartesian)
	}
}
