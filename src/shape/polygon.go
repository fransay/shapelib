package shape

import (
	"errors"
	"math"

	"github.com/fransay/shapelib/internal/utils"
	"github.com/fransay/shapelib/src/geom"
)

// Polygon defines a polygon
type Polygon struct {
	pts []geom.Point2D
}

// NewPolygon constructs a new Polygon
func NewPolygon(pts []geom.Point2D) *Polygon {
	return &Polygon{pts}
}

// Area returns the area of a polygon using the shoelace formula
func (p Polygon) Area() float64 {
	var forwardPass, backwardPass float64
	n := len(p.pts)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		forwardPass += p.pts[i].X * p.pts[j].Y
		backwardPass += p.pts[j].X * p.pts[i].Y
	}
	area := 0.5 * (forwardPass - backwardPass)
	return math.Abs(area)
}

// Centroid returns the center (centroid) of a polygon
func (p Polygon) Centroid() (cent geom.Point2D, err error) {
	var xSum, ySum float64
	n := len(p.pts)
	if n == 0 {
		return geom.Point2D{}, errors.New("can't determine the centroid")
	}
	for _, pt := range p.pts {
		xSum += pt.X
		ySum += pt.Y
	}
	cent = geom.Point2D{
		X: xSum / float64(n),
		Y: ySum / float64(n),
	}
	return cent, nil
}

// SegmentDistances returns all segment distances
func (p Polygon) SegmentDistances() []float64 {
	n := len(p.pts)
	segmentDistances := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		startPoint := []float64{p.pts[i].X, p.pts[i].Y}
		endPoint := []float64{p.pts[j].X, p.pts[j].Y}
		segDistance := utils.Distance(startPoint, endPoint)
		segmentDistances = append(segmentDistances, segDistance)
	}
	return segmentDistances
}

// Perimeter returns the total perimeter of a polygon
func (p Polygon) Perimeter() (perimeter float64) {
	perimeter = utils.SumTotalFloat(p.SegmentDistances())
	terminalPoint := []float64{p.pts[len(p.pts)-1].X, p.pts[len(p.pts)-1].Y}
	origStartPoint := []float64{p.pts[0].X, p.pts[0].Y}
	finalSegDistance := utils.Distance(origStartPoint, terminalPoint)
	perimeter += finalSegDistance
	return perimeter
}

// ShortestLineSegment returns the shortest line segment of a polygon
func (p Polygon) ShortestLineSegment() float64 {
	segDistances := p.SegmentDistances()
	minLine := segDistances[0]
	for _, dist := range segDistances {
		if dist < minLine {
			minLine = dist
		}
	}
	return minLine
}

// IsClosed returns a boolean if a polygon is closed or not
func (p Polygon) IsClosed() bool {
	return !p.IsOpened()
}

// IsOpened returns a boolean if a polygon is closed or not
func (p Polygon) IsOpened() (isOpenedStatus bool) {
	var firstElement, LastElement geom.Point2D
	for i, v := range p.pts {
		if i == 0 {
			firstElement = v
		} else if i == len(p.pts)-1 {
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
func (p Polygon) NumberOfLineSegments() int {
	return p.NumberOfNodes() - 1
}

// NumberOfNodes returns number of nodes in a polygon
func (p Polygon) NumberOfNodes() int {
	return len(p.pts)
}

// Rotate return the set of points (forming the polygon) under rotate transformation
func (p Polygon) Rotate(rotationAngle float64) (rotate Polygon) {
	var transformPolygon Polygon
	angleRad := utils.Deg2Rad(rotationAngle)

	for _, pointInPoly := range p.pts {
		pointInPoly.X = pointInPoly.X*math.Cos(angleRad) - pointInPoly.Y*math.Sin(angleRad)
		pointInPoly.Y = pointInPoly.X*math.Sin(angleRad) + pointInPoly.Y*math.Cos(angleRad)
		transformPolygon.pts = append(transformPolygon.pts, pointInPoly)

	}
	return transformPolygon
}

// Scale return the set of points (forming the polygon) under a scale transformation
func (p Polygon) Scale(scalarVector []float64) (scale Polygon) {
	var scaledPolygon Polygon
	for _, element := range p.pts {
		element.X = element.X * scalarVector[0]
		element.Y = element.Y * scalarVector[1]
		scaledPolygon.pts = append(scaledPolygon.pts, element)
	}
	return scaledPolygon
}

// IsEqual returns equal status in boolean by comparing two polygons
func (p Polygon) IsEqual(polygon Polygon) (equal bool) {
	// for two polygons to be equal, they must be of the same or similar length
	equal = false
	containSuccessChecker := 0
	for _, polygonElement := range p.pts {
		if p.Contains(polygon, polygonElement) {
			containSuccessChecker += 1
		}
	}

	if containSuccessChecker == len(p.pts) {
		equal = true
	}
	return equal
}

// Contains checks if a point is in a polygon
func (p Polygon) Contains(poly Polygon, point geom.Point2D) (contains bool) {
	contains = false
	for _, pointElement := range poly.pts {
		if pointElement == point {
			contains = true
		}
	}
	return contains
}
