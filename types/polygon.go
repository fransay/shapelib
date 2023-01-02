package types

// Polygon type polygon
type Polygon []Point2D

// polygon methods

// Area returns the area of the polygon
// algorithm: shoelace
// Optimise for resource and time complexities
func (p *Polygon) Area() (area float64) {
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

// ShortestLineString returns the shortest line string
// use ranking algorithm to sort out distances
func (p *Polygon) ShortestLineString() (perim float64) {
	return perim
}

// IsClosed returns a boolean value whether polygon is closed or not
func (p *Polygon) IsClosed() (isClosedStatus float64) {
	return isClosedStatus
}

// IsOpened returns a boolean value whether polygon is opened or not
func (p *Polygon) IsOpened() (isOpenedStatus float64) {
	return isOpenedStatus
}
