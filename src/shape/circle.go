package shape

import (
	"errors"
	"log"
	"math"

	"github.com/fransay/shapelib/internal/utils"
	"github.com/fransay/shapelib/src/geom"
)

var ErrorOfCollinearity = errors.New("points are collinear")

// Circle type
type Circle struct {
	Radius   float64
	Centroid geom.Point2D
}

// NewCircle constructs a new circle object
func NewCircle(radius float64, centroid geom.Point2D) (*Circle, error) {
	if radius < 0.0 {
		return nil, errors.New("can't create a new circle, radius can't be negative")
	}
	return &Circle{
		Radius:   radius,
		Centroid: centroid,
	}, nil
}

// Area returns the area of a circle
func (c *Circle) Area() (area float64) {
	area = math.Pi * math.Pow(c.Radius, 2)
	return area
}

// ArcLength returns the arc length of a circle given the angle of arc extent in degrees
func (c *Circle) ArcLength(angle, radius float64) (arc float64) {
	angleInRadians := utils.Deg2Rad(angle)
	return angleInRadians * radius
}

// SectorArea returns the sector area of a circle
func (c *Circle) SectorArea(angle, radius float64) (area float64) {
	angleInRadians := utils.Deg2Rad(angle)
	return angleInRadians * radius * 0.5
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
func (c *Circle) ContainsPoint(point geom.Point2D) bool {
	distanceBtnPointAndCenter := math.Sqrt(math.Pow(point.X-c.Centroid.X, 2) + math.Pow(point.Y-c.Centroid.Y, 2))
	return distanceBtnPointAndCenter <= c.Radius
}

// CircleFromPoints forms a new circle given three arbitrary set of point that lie
// on the circumference of a circle collinear point means all three points are highly
// unlikely to be on the circumference.
func (c *Circle) CircleFromPoints(a, b, d geom.Point2D) (Circle, error) {
	// check if points are collinear; collinear points can for a circle
	isCollinear, err := geom.IsCollinear(a, b, d)
	if err != nil {
		log.Fatal(err)
	}
	if isCollinear {
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
	intersectedPoint, err := LineIntersectGivenEqn(eqnAB, eqnBD)
	if err != nil {
		return Circle{}, err
	}
	// radius -> intersectedPoint <-> any arbitrary points on the circle
	radius := distance(intersectedPoint, a)
	return Circle{radius, intersectedPoint}, nil
}

// distance return the distance between two points
func distance(pt1, pt2 geom.Point2D) float64 {
	return math.Sqrt(math.Pow(pt2.X-pt1.X, 2) + math.Pow(pt2.X-pt1.X, 2))
}

// gradient return the gradient / slope of a perpendicular bisector.
func gradient(pt1, pt2 geom.Point2D) float64 {
	deltaY := pt2.Y - pt1.Y
	deltaX := pt2.X - pt2.X
	grad := deltaY / deltaX
	return grad
}

// LineIntersectGivenEqn returns the intersectional point of a two line given their equations
// equation of lines take this format, e.g y = 3x + 6 , which is represented in LineIntersectGivenEqn
// arguments as {}float64[1,3,6], i.e. the coefficient of the variables in the equation of the line.
func LineIntersectGivenEqn(lineOne, lineTwo [3]float64) (intersectPoint geom.Point2D, err error) {
	a1, b1, c1 := lineOne[0], lineOne[1], lineOne[2]
	a2, b2, c2 := lineTwo[0], lineTwo[1], lineTwo[2]
	det := a1*b2 - a2*b1
	if det == 0 {
		return geom.Point2D{}, errors.New("lines are parallel or coincident, no unique intersection point")
	}
	x := (b1*c2 - b2*c1) / det
	y := (a2*c1 - a1*c2) / det
	return *geom.NewPoint2D(int(x), int(y)), nil
}

// AsLineString returns a line string given a set of points that defines a circle
func (c *Circle) AsLineString(points ...geom.Point2D) (geom.LineString, error) {
	var nilError error = errors.New("no points to form line string")
	if len(points) == 0 {
		return geom.LineString{}, nilError
	}
	var lineString geom.LineString
	for _, point := range points {
		lineString = append(lineString, point)
	}
	return lineString, nil
}
