package types

type LineString []Point2D

// Length of a line string

func (ll *LineString) Distance() float64 {
	return 0.0
}

// returns the total number of elements in an instance of type LineString
func (ll *LineString) length() int {
	var indexTracker int
	for index, _ := range *ll {
		indexTracker += index
	}
	return indexTracker
	// complexity : O(n)
}

// takes a Point type argument and returns its position
func (ll *LineString) index(args Point2D) (index int) {
	var pos int
	for index, value := range *ll {
		if value == args {
			pos = index
		}
	}
	return pos
	// complexity : O(n)
}
