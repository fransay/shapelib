package shape_test

import (
	"shapelib/shape"
	"testing"
)

func TestTriangle(t *testing.T) {
	triangle := shape.Triangle{SideOne: 5, SideTwo: 10, SideThree: 15, Base: 20, Height: 30}
	// area by sides
	resultAreaBySides := triangle.AreaBySides()
	expectedAreaBySides := 34.0
	if resultAreaBySides != expectedAreaBySides {
		t.Errorf("Got %v, expected %v", resultAreaBySides, expectedAreaBySides)
	}
	// area by base && height
	resultAreaByBaseAndHeight := triangle.AreaByBaseHeight()
	expectedAreaByBaseAndHeight := 34.0
	if resultAreaByBaseAndHeight != expectedAreaByBaseAndHeight {
		t.Errorf("Got %v, expected %v", resultAreaByBaseAndHeight, expectedAreaByBaseAndHeight)
	}
	// perimeter
	resultPerimeter := triangle.Perimeter()
	expectedPerimeter := 15.0
	if resultPerimeter != expectedPerimeter {
		t.Errorf("Got %v, expected %v", resultPerimeter, expectedPerimeter)
	}
	// type
	resultType := triangle.Type()
	expectedType := "Scalene"
	if resultType != expectedType {
		t.Errorf("Got %v, expected %v", resultType, expectedPerimeter)
	}
}
