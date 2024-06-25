package types

import (
	"math"
	"shapelib/utils"
)

// Polygon type polygon
type Polygon []Point2D

// Area returns area of a polygon using shoelace
func (p Polygon) Area() (area float64) {
	var forwardPass float64
	var backwardPass float64
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
func (p *Polygon) Centroid() (cent Point2D) {
	var xSum, ySum float64
	for _, values := range *p {
		xSum = xSum + values.X
		ySum = ySum + values.Y
	}
	cent = Point2D{
		X: xSum / 2, Y: ySum / 2,
	}
	return cent
}

// Perimeter TODO: Fix perimeter
// Perimeter returns the total perimeter of a polygon
func (p Polygon) Perimeter() (perim float64) {
	for i, j := 0, 1; j < len(p); i, j = i+1, j+1 {
		// destruct element
		startPoint := []float64{p[i].X, p[i].Y}
		nextPoint := []float64{p[j].X, p[j].Y}
		segDistance := utils.Distance(startPoint, nextPoint)
		perim += segDistance
	}
	// add perimeter of closing segment
	terminalPoint := []float64{p[len(p)-1].X, p[len(p)-1].Y}
	origStartPoint := []float64{p[0].X, p[0].Y}
	finalSegDistance := utils.Distance(origStartPoint, terminalPoint)
	perim += finalSegDistance

	return perim
}

// ShortestLineSegment returns the shortest line segment of a polygon
func (p *Polygon) ShortestLineSegment() int {
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

// NumberOfLineSegments returns the number of line segments forming a polygon
func (p *Polygon) NumberOfLineSegments() int {
	return p.NumberOfNodes() - 1
}

// NumberOfNodes returns number of nodes in a polygon
func (p *Polygon) NumberOfNodes() int {
	return len(*p)
}

// Rotate return the set of points (forming the polygon) under rotate transformation
func (p *Polygon) Rotate(rotationAngle float64) (rotate Polygon) {
	var transformPolygon Polygon
	angleRad := utils.Deg2Rad(rotationAngle)

	for _, pointInPoly := range *p {
		pointInPoly.X = pointInPoly.X*math.Cos(angleRad) - pointInPoly.Y*math.Sin(angleRad)
		pointInPoly.Y = pointInPoly.X*math.Sin(angleRad) + pointInPoly.Y*math.Cos(angleRad)
		transformPolygon = append(transformPolygon, pointInPoly)

	}

	return transformPolygon
}

// Scale return the set of points (forming the polygon) under a scale transformation
func (p *Polygon) Scale(scalarVector []float64) (scale Polygon) {
	var scaledPolygon Polygon
	for _, element := range *p {
		element.X = element.X * scalarVector[0]
		element.Y = element.Y * scalarVector[1]
		scaledPolygon = append(scaledPolygon, element)
	}
	return scaledPolygon
}

// IsEqual returns equal status in boolean by comparing two polygons
func (p *Polygon) IsEqual(polygon Polygon) (equal bool) {
	// for two polygons to be equal, they must be of the same or similar length
	equal = false
	containSuccessChecker := 0
	for _, polygonElement := range *p {
		if p.Contains(polygon, polygonElement) {
			containSuccessChecker += 1
		}
	}

	if containSuccessChecker == len(*p) {
		equal = true

	}
	return equal
}

// Contains returns a boolean representing the presence of a point in an array of Polygon types
func (p *Polygon) Contains(poly Polygon, point Point2D) (contains bool) {
	contains = false
	for _, pointElement := range poly {
		if pointElement == point {
			contains = true
		}
	}
	return contains
}
