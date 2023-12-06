package types

import (
	"math"
	s "shapelib/types"
	"shapelib/types/point"
	"testing"
)

// distance test
func TestLineSegmentDistance(t *testing.T) {
	// line segment one
	var lineSegmentOne s.LineSegment = s.LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 5, Y: 10}, PointB: struct {
		X float64
		Y float64
	}{X: 10, Y: 15}}
	// line segment two
	var lineSegmentTwo s.LineSegment = s.LineSegment{PointA: struct {
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

// midpoint test
func TestLineSegmentMidPoint(t *testing.T) {
	var lineSegmentOne s.LineSegment = s.LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 10, Y: 14}, PointB: struct {
		X float64
		Y float64
	}{X: 20, Y: 22}}

	// midpoint test one
	var midPointLineResultLineSegmentOne point.Point2D = lineSegmentOne.MidPoint()
	var midPointLineExpectedLineSegmentOne point.Point2D = point.Point2D{X: 15, Y: 18}
	if midPointLineExpectedLineSegmentOne != midPointLineResultLineSegmentOne {
		t.Errorf("Expected %f, got %f", midPointLineExpectedLineSegmentOne, midPointLineResultLineSegmentOne)
	}
	// midpoint test two
	var lineSegmentTwo s.LineSegment = s.LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 10, Y: 14}, PointB: struct {
		X float64
		Y float64
	}{X: 20, Y: 22}}

	// midpoint test one
	var midPointLineResultLineSegmentTwo point.Point2D = lineSegmentTwo.MidPoint()
	var midPointLineExpectedLineSegmentTwo point.Point2D = point.Point2D{X: 15, Y: 18}
	if midPointLineExpectedLineSegmentTwo != midPointLineResultLineSegmentTwo {
		t.Errorf("Expected %f, got %f", midPointLineExpectedLineSegmentTwo, midPointLineResultLineSegmentTwo)
	}

}

// bearing test
func TestLineSegmentBearing(t *testing.T) {
	// instance 1 :line segment
	var lineSegmentOne s.LineSegment = s.LineSegment{
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
	var lineSegmentTwo s.LineSegment = s.LineSegment{
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
