package types

// Polygon type polygon
type Polygon []Point2D

// polygon methods

// Area returns the area of the polygon
// algorithm: shoelace
// Optimise for resource and time complexities
func (p *Polygon) Area() (area float64) {
	// fetch Point2D points
	return area

}

// Centroid returns the centroid of the polygon
// algorithms: none yet
func (p *Polygon) Centroid() (cent float64) {
	return cent
}

// Perimeter returns total distance around the polygon
// traverse the total length of compositing strings
func (p *Polygon) Perimeter() (perim float64) {
	return perim
}

// ShortestLineSegment returns the shortest line string
// use ranking algorithm to sort out distances
func (p *Polygon) ShortestLineSegment() int {
	// create a slice
	// find the distance between consequence points/station
	// append them to the slice
	// sort distance
	// return smallest
	return 0.0
}

// IsClosed returns a boolean value whether polygon is closed or not
// if element of array is equal to last element of array
// then polygon is closed
func (p *Polygon) IsClosed() (isClosedStatus bool) {
	if !p.IsOpened() {
		isClosedStatus = true
	} else {
		isClosedStatus = false
	}
	return isClosedStatus
}

// IsOpened returns a boolean value whether polygon is opened or not
// if element of first array is not equal to element of array
// then polygon is opened
func (p *Polygon) IsOpened() (isOpenedStatus bool) {
	var firstElement, LastElement Point2D
	for i, v := range *p {
		if i == 0 {
			firstElement = v
		} else if i == len(*p)-1 {
			LastElement = v
		}
		if firstElement == LastElement {
			isOpenedStatus = true
		} else {
			isOpenedStatus = false
		}
	}
	return isOpenedStatus
}

// NumberOfLineSegments returns the number of compositing line segments
// number of line segments is equal to number of nodes
func (p *Polygon) NumberOfLineSegments() int {
	return p.NumberOfNodes()
}

// NumberOfNodes returns number of nodes in a polygon
func (p *Polygon) NumberOfNodes() int {
	var numbReturn int = 0
	for index, _ := range *p {
		numbReturn += index
	}
	return numbReturn
}

// return the distance of a specific line segment using the indexes

func (p *Polygon) LineSegmentDistance(ind int) (distance float64) {
	return distance

}
