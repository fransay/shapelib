package types

import functs "shapelib/functs/2D"

// LineString type
type LineString [...]Point2D

// Length of a linestring
// TODO improve complexity to O(1): constant time
func (l *LineString) Length() float64 {
	var totLength float64
	for i, j := 0, 1; i < len(l); i, j = i+1, j+1 {
		var pointOne Point2D = l[i]
		var pointTwo Point2D = l[j]
		totLength += functs.EDistance(pointOne, pointTwo)
	}
	return totLength
}

// NumberOfLineSegments returns the total number of elements in an instance of type LineString
// complexity : O(n)
// TODO improve complexity to O(1): constant time
func (l *LineString) NumberOfLineSegments() int {
	var indexTracker int
	for index, _ := range *l {
		indexTracker += index
	}
	return indexTracker

}

// Index takes a Point type argument and returns its position
// complexity : O(n)
// TODO improve complexity to O(1): constant time
func (l *LineString) Index(args Point2D) (index int) {
	var pos int
	for index, value := range *l {
		if value == args {
			pos = index
		}
	}
	return pos

}
