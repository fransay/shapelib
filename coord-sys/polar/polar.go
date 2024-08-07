package polar

import (
	"math"
	cart2 "shapelib/coord-sys/cartesian"
)

// Polar type
type Polar struct {
	Distance float64 // Distance from the origin
	Angle    float64 // referenced Angle, e.g. bearing
}

// ToCartesian converts polar to cartesian coordinates
func (p *Polar) ToCartesian() (c cart2.Cart2D) {
	x := p.Distance * math.Cos(p.Angle)
	y := p.Distance * math.Sin(p.Angle)
	c = cart2.Cart2D{X: x, Y: y}
	return c
}

// Distance2DMS converts distance to degree minutes and seconds representation
func (p *Polar) Distance2DMS() (dms []float64) {
	var degree = math.Floor(p.Distance) // whole number part
	var decimalMinutes = (p.Distance - degree) * 60
	var minutes = math.Floor(decimalMinutes)
	var sec = (decimalMinutes - minutes) * 60
	dms = append(dms, degree, minutes, sec)
	return dms
}

// ToRadians converts bearing in degrees to radians
func (p *Polar) ToRadians() (rad float64) {
	rad = p.Angle * (math.Pi / 180)
	return rad
}

// Point2PointDistance returns to distance between self.point and other point
func (p *Polar) Point2PointDistance(point Polar) (dist float64) {
	addDist := p.Distance + point.Distance
	prodDist := p.Distance * point.Distance
	addAngle := p.Angle + point.Angle
	dist = math.Sqrt(addDist - 2*prodDist*math.Cos(addAngle))
	return dist
}
