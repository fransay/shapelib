package coord_sys

import (
	"math"
)

type Polar struct {
	Distance float64 // Distance from the origin, in dd
	Angle    float64 // referenced Angle, e.g. bearing
}

// ToCartesian converts polar to cartesian coordinates
func (p *Polar) ToCartesian() (c Cart) {
	x := p.Distance * math.Cos(p.Angle)
	y := p.Distance * math.Sin(p.Angle)
	c.X = x
	c.Y = y
	return c
}

// Distance2DMS converts Distance to degree minutes and seconds representation
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

// Point2PointDistance Distance to another point with a polar representation
func (p *Polar) Point2PointDistance(point Polar) float64 {
	return 0.0
}
