package tests

import (
	"log"
	"shapelib/utils"
	"testing"
)

func TestCartCordAxis(t *testing.T) {
	// axis
	axis := Axis{Start: 2, End: 10, Step: 1, Sign: "+"}
	outputGenerateAxialValues := axis.GenerateAxisValues()
	expectedGenerateAxialValues := []float64{2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 10.0}
	if utils.EqualSlice(outputGenerateAxialValues, expectedGenerateAxialValues) {
		t.Errorf("expected: %v, got: %v", expectedGenerateAxialValues, outputGenerateAxialValues)
	}

	// quadrant
	quadrant := Quadrant{X: Axis{Start: 1, End: 10, Step: 1, Sign: "+"}, Y: Axis{Start: 10, End: 100, Step: 5, Sign: "+"}}
	quadrantType, err := quadrant.Type()
	if err != nil {
		log.Fatalf("Invalid Quadrant Type: %v", err)
	}
	if quadrantType != SecondQuadrant {
		t.Errorf("Invalid Quadrant Type: %v", quadrantType)
	}
}
