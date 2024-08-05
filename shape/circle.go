package shape

import (
	"errors"
	"math"
	"shapelib/types"
)

var ErrorOfCollinearity = errors.New("points are collinear")

type Circle struct {
	Radius   float64 // from the center to the circumference of the circle
	Centroid types.Point2D
}

// Area returns the area of a circle
func (c *Circle) Area() (area float64) {
	area = math.Pi * math.Pow(c.Radius, 2)
	return area
}

// Circumference returns the circumference of a circle
func (c *Circle) Circumference() (circumference float64) {
	circumference = 2 * math.Pi * c.Radius
	return circumference
}

// ChordLength returns the chord length of circle based on sagitta
func (c *Circle) ChordLength(sagitta float64) (chord float64) {
	chord = 2 * math.Sqrt(math.Pow(c.Radius, 2)-math.Pow(sagitta, 2))
	return chord
}

// ContainsPoint checks if a point lies in the radius extent of a circle
func (c *Circle) ContainsPoint(point types.Point2D) bool {
	distanceBtnPointAndCenter := math.Sqrt(math.Pow(point.X-c.Centroid.X, 2) + math.Pow(point.Y-c.Centroid.Y, 2))
	if distanceBtnPointAndCenter <= c.Radius {
		return true
	}
	return false
}

// CircleFromPoints forms a new circle given three arbitrary set of point that lie
// on the circumference of a circle collinear point means all three points are highly
// unlikely to be on the circumference.
func (c *Circle) CircleFromPoints(a, b, d types.Point2D) (Circle, error) {
	// check if points are collinear; collinear points can for a circle
	if types.IsCollinear(a, b, d) == true {
		return Circle{}, ErrorOfCollinearity
	}
	slopeAB := gradient(a, b)
	slopeBD := gradient(b, d)

	abXSum := a.X + b.X
	abYSum := a.Y + d.Y
	bdXSum := b.X + d.X
	bdYSum := b.Y + d.Y

	lineConstantAB := abXSum + (abYSum * slopeAB)
	lineConstantBD := bdXSum + (bdYSum * slopeBD)

	eqnAB := [3]float64{1.0, -2.0, lineConstantAB}
	eqnBD := [3]float64{1.0, -2.0, lineConstantBD}
	intersectedPoint := LineIntersectGivenEqn(eqnAB, eqnBD)
	// radius -> intersectedPoint <-> any arbitrary points on the circle
	radius := distance(intersectedPoint, a)
	return Circle{radius, intersectedPoint}, nil
}

// distance return the distance between two points
func distance(pt1, pt2 types.Point2D) float64 {
	return math.Sqrt(math.Pow(pt2.X-pt1.X, 2) + math.Pow(pt2.X-pt1.X, 2))
}

// gradient return the gradient / slope of a perpendicular bisector.
func gradient(pt1, pt2 types.Point2D) float64 {
	deltaY := pt2.Y - pt1.Y
	deltaX := pt2.X - pt2.X
	grad := deltaY / deltaX
	return grad
}

// LineIntersectGivenEqn returns the intersectional point of a two line given their equations
// equation of lines take this format, e.g y = 3x + 6 , which is represented in LineIntersectGivenEqn
// arguments as [1,3,6], i.e. the coefficient of the variables in the equation of the line.
func LineIntersectGivenEqn(lineOne, lineTwo [3]float64) (intersectPoint types.Point2D) {} // todo: complete function

// AsLineString returns a line string given a set of points that defines a circle
func (c *Circle) AsLineString(points ...types.Point2D) (types.LineString, error) {
	var nilError error = errors.New("no points to form line string")
	if points == nil || len(points) == 0 {
		return types.LineString{}, nilError
	}
	var lineString types.LineString
	for _, point := range points {
		lineString = append(lineString, point)
	}
	return lineString, nil
}
