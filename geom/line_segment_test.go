package geom

import (
	"math"
	"testing"
)

// distance funct_tests
func TestLineSegmentDistance(t *testing.T) {
	// line segment one
	var lineSegmentOne LineSegment = LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 5, Y: 10}, PointB: struct {
		X float64
		Y float64
	}{X: 10, Y: 15}}
	// line segment two
	var lineSegmentTwo LineSegment = LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 234.543, Y: 24043.354}, PointB: struct {
		X float64
		Y float64
	}{X: 544.654, Y: 354.544}}

	// line segment one result
	var distanceResultLineSegmentOne = lineSegmentOne.Distance()
	var distanceExpectedLineSegmentOne = 7.017
	if math.Trunc(distanceResultLineSegmentOne) != math.Trunc(distanceExpectedLineSegmentOne) {
		t.Errorf("Expected %f, got %f", distanceExpectedLineSegmentOne, distanceResultLineSegmentOne)
	}

	// line segment two results
	var distanceResultLineSegmentTwo = lineSegmentTwo.Distance()
	var distanceExpectedLineSegmentTwo = 23690.84
	if math.Trunc(distanceResultLineSegmentTwo) != math.Trunc(distanceExpectedLineSegmentTwo) {
		t.Errorf("Expected %f, got %f", distanceExpectedLineSegmentTwo, distanceResultLineSegmentTwo)
	}

}

// midpoint funct_tests
func TestLineSegmentMidPoint(t *testing.T) {
	var lineSegmentOne LineSegment = LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 10, Y: 14}, PointB: struct {
		X float64
		Y float64
	}{X: 20, Y: 22}}

	// midpoint funct_tests one
	var midPointLineResultLineSegmentOne Point2D = lineSegmentOne.MidPoint()
	var midPointLineExpectedLineSegmentOne Point2D = Point2D{X: 15, Y: 18}
	if midPointLineExpectedLineSegmentOne != midPointLineResultLineSegmentOne {
		t.Errorf("Expected %f, got %f", midPointLineExpectedLineSegmentOne, midPointLineResultLineSegmentOne)
	}
	// midpoint funct_tests two
	var lineSegmentTwo LineSegment = LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 10, Y: 14}, PointB: struct {
		X float64
		Y float64
	}{X: 20, Y: 22}}

	// midpoint funct_tests one
	var midPointLineResultLineSegmentTwo Point2D = lineSegmentTwo.MidPoint()
	var midPointLineExpectedLineSegmentTwo Point2D = Point2D{X: 15, Y: 18}
	if midPointLineExpectedLineSegmentTwo != midPointLineResultLineSegmentTwo {
		t.Errorf("Expected %f, got %f", midPointLineExpectedLineSegmentTwo, midPointLineResultLineSegmentTwo)
	}

}

// bearing funct_tests
func TestLineSegmentBearing(t *testing.T) {
	// instance 1 :line segment
	var lineSegmentOne LineSegment = LineSegment{
		PointA: struct {
			X float64
			Y float64
		}{X: 20, Y: 70}, PointB: struct {
			X float64
			Y float64
		}{X: 80, Y: 100}}

	// bearing one
	var bearingLineResultLineSegmentOne float64 = lineSegmentOne.Bearing()
	var bearingLineExpectedLineSegmentOne float64 = 1.107
	if math.Trunc(bearingLineExpectedLineSegmentOne) != math.Trunc(bearingLineResultLineSegmentOne) {
		t.Errorf("Expected %f got %f", bearingLineExpectedLineSegmentOne, bearingLineResultLineSegmentOne)
	}

	// instance 2 :line segment
	var lineSegmentTwo LineSegment = LineSegment{
		PointA: struct {
			X float64
			Y float64
		}{X: -20, Y: 50}, PointB: struct {
			X float64
			Y float64
		}{X: 80, Y: -100}}

	// bearing two
	var bearingLineResultLineSegmentTwo float64 = lineSegmentTwo.Bearing()
	var bearingLineExpectedLineSegmentTwo float64 = 0.5880026
	if math.Trunc(bearingLineExpectedLineSegmentTwo) != math.Trunc(bearingLineResultLineSegmentTwo) {
		t.Errorf("Expected %f got %f", bearingLineExpectedLineSegmentTwo, bearingLineResultLineSegmentTwo)
	}

}
