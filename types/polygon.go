package types

import (
	"math"
	"shapelib/types/point"
)

// Polygon type polygon
type Polygon []point.Point2D

// Area returns area of a polygon using shoelace
func (p Polygon) Area() (area float64) {
	var forwardPass = 0.0
	var backwardPass = 0.0
	for i, j := 0, 1; j < len(p); i, j = i+1, j+1 {
		startPoint := p[i]
		terminalPoint := p[j]
		forwardPass = forwardPass + startPoint.X*terminalPoint.Y
		backwardPass = backwardPass + terminalPoint.X*startPoint.Y
		area = (forwardPass - backwardPass) / 2
	}
	return math.Abs(area) // abs for magnitude sense only
}

// Centroid returns the centre point of polygon
func (p *Polygon) Centroid() (cent point.Point2D) {
	var xSum, ySum float64
	for _, values := range *p {
		xSum = xSum + values.X
		ySum = ySum + values.Y
	}
	cent = point.Point2D{
		X: xSum / 2, Y: ySum / 2,
	}
	return cent
}

// Perimeter returns the total perimeter of a polygon
func (p *Polygon) Perimeter() (perim float64) {
	return perim
}

// ShortestLineSegment returns the shortest line segment of a polygon
func (p *Polygon) ShortestLineSegment() int {
	// create a slice
	// find the distance between consequence points/station
	// append them to the slice
	// sort distance
	// return smallest
	return 0.0
}

// IsClosed returns a boolean if a polygon is closed or not
func (p *Polygon) IsClosed() (isClosedStatus bool) {
	if !p.IsOpened() {
		isClosedStatus = true
	} else {
		isClosedStatus = false
	}
	return isClosedStatus
}

// IsOpened returns a boolean if a polygon is closed or not
func (p *Polygon) IsOpened() (isOpenedStatus bool) {
	var firstElement, LastElement point.Point2D
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

// NumberOfLineSegments returns the number of line segments forming a polygon
func (p *Polygon) NumberOfLineSegments() int {
	return p.NumberOfNodes()
}

// NumberOfNodes returns number of nodes in a polygon
func (p *Polygon) NumberOfNodes() int {

	/*
		var numbReturn int = 0
		for index, _ := range *p {
			numbReturn += index
		}
	*/
	return len(*p)
}

// LineSegmentDistance returns the distance of a specific line segment using the indexes
func (p *Polygon) LineSegmentDistance(ind int) (distance float64) {
	return distance

}

// Affine returns the transformed set of points (forming the polygon) under affine transformation
func (p *Polygon) Affine() (affine Polygon) {
	return affine
}

// Projective return the transformed set of points (forming the polygon) under projective transformation
func (p *Polygon) Projective() (projective Polygon) {
	return projective
}

// Rotate return the set of points (forming the polygon) under rotate transformation
func (p *Polygon) Rotate() (rotate Polygon) {
	return rotate
}

// Scale return the set of points (forming the polygon) under a scale transformation
func (p *Polygon) Scale() (scale Polygon) {
	return scale
}
