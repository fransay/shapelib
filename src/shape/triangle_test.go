package shape

import (
	"testing"
)

var triangle = Triangle{
	SideOne:   50.0,
	SideTwo:   60.0,
	SideThree: 30.0,
	Base:      30.0,
	Height:    10.0}

func TestAreaBySides(t *testing.T) {
	resultAreaBySides := triangle.AreaBySides()
	expectedAreaBySides := 748.3314773547883
	if resultAreaBySides != expectedAreaBySides {
		t.Errorf("Got %v, expected %v", resultAreaBySides, expectedAreaBySides)
	}
}

func TestAreaByBaseAndHeight(t *testing.T) {
	resultAreaByBaseAndHeight := triangle.AreaByBaseHeight()
	expectedAreaByBaseAndHeight := 150.0
	if resultAreaByBaseAndHeight != expectedAreaByBaseAndHeight {
		t.Errorf("Got %v, expected %v", resultAreaByBaseAndHeight, expectedAreaByBaseAndHeight)
	}
}

func TestPerimeter(t *testing.T) {
	resultPerimeter := triangle.Perimeter()
	expectedPerimeter := 140.0
	if resultPerimeter != expectedPerimeter {
		t.Errorf("Got %v, expected %v", resultPerimeter, expectedPerimeter)
	}

}

func TestType(t *testing.T) {
	resultType := triangle.Type()
	expectedType := "Scalene"
	if resultType != expectedType {
		t.Errorf("Got %v, expected %v", resultType, expectedType)
	}

}
