package types

// LineString type
type LineString []Point2D

// Length of a linestring
// TODO complete length function
func (l *LineString) Length() float64 {
	return 0.0
}

// NumberOfLineSegments returns the total number of elements in an instance of type LineString
func (l *LineString) NumberOfLineSegments() int {
	var indexTracker int
	for index, _ := range *l {
		indexTracker += index
	}
	return indexTracker
	// complexity : O(n)
	// TODO improve complexity to O(1): constant time
}

// Index takes a Point type argument and returns its position
func (l *LineString) Index(args Point2D) (index int) {
	var pos int
	for index, value := range *l {
		if value == args {
			pos = index
		}
	}
	return pos
	// complexity : O(n)
	// TODO improve complexity to O(1): constant time
}
