package shape

import (
	"testing"

	"github.com/fransay/shapelib/src/geom"
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
		t.Errorf("Got = %v, Expected = %v", resultAreaBySides, expectedAreaBySides)
	}
}

func TestAreaByBaseAndHeight(t *testing.T) {
	resultAreaByBaseAndHeight := triangle.AreaByBaseHeight()
	expectedAreaByBaseAndHeight := 150.0
	if resultAreaByBaseAndHeight != expectedAreaByBaseAndHeight {
		t.Errorf("Got = %v, Expected = %v", resultAreaByBaseAndHeight, expectedAreaByBaseAndHeight)
	}
}

func TestPerimeter(t *testing.T) {
	resultPerimeter := triangle.Perimeter()
	expectedPerimeter := 140.0
	if resultPerimeter != expectedPerimeter {
		t.Errorf("Got = %v, Expected = %v", resultPerimeter, expectedPerimeter)
	}

}

func TestType(t *testing.T) {
	resultType := triangle.Type()
	expectedType := "Scalene"
	if resultType != expectedType {
		t.Errorf("Got =  %v, Expected = %v", resultType, expectedType)
	}

}

var traingleByCoordinates = TriangleByCoordinates{
	PointOne:   geom.Point2D{X: 50, Y: 160},
	PointTwo:   geom.Point2D{X: 70, Y: 120},
	PointThree: geom.Point2D{X: 45, Y: 500},
}

// triangle by coordinates
func TestCenter(t *testing.T) {
	var observedCenterCoordinates = traingleByCoordinates.Center()
	var expectedCenterCoordinates = geom.Point2D{X: 55, Y: 260}
	if !observedCenterCoordinates.IsEqual(expectedCenterCoordinates) {
		t.Errorf("observed coordinates is not equal to expected.")
	}

}

func TestArea(t *testing.T) {
	var observedAreaCoordinates = traingleByCoordinates.Area()
	var expectedAreaCoordinates = 3300.00
	if observedAreaCoordinates != expectedAreaCoordinates {
		t.Errorf("observed area %v is not equal to expected area %v", observedAreaCoordinates, expectedAreaCoordinates)
	}
}
