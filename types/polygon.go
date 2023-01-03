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
	return 0
}

// IsClosed returns a boolean value whether polygon is closed or not
// if element of array is equal to last element of array
// then polygon is closed
func (p *Polygon) IsClosed() (isClosedStatus float64) {
	return isClosedStatus
}

// IsOpened returns a boolean value whether polygon is opened or not
// if element of first array is not equal to element of array
// then polygon is opened
func (p *Polygon) IsOpened() (isOpenedStatus bool) {
	return isOpenedStatus
}

// NumberOfLineSegments returns the number of compositing line segments
// number of line segments is equal to number of nodes
func (p *Polygon) NumberOfLineSegments() int {
	return p.NumberOfNodes()
}

// NumberOfNodes returns number of nodes in a polygon
func (p *Polygon) NumberOfNodes() int {
	var numbReturn int = 1
	for index, _ := range *p {
		numbReturn += index
	}
	return numbReturn
}

//
