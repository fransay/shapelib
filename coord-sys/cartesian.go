package coord_sys

import "math"

// Cart cartesian type
type Cart struct {
	X float64
	Y float64
}

// Cart3D cartesian type
type Cart3D struct {
	X float64
	Y float64
	Z float64
}

// ToPolar convert cartesian to Polar coordinates
// reference point is origin
func (c *Cart) ToPolar() (p Polar) {
	cartDist := distance(Cart{0, 0}, *c)
	cartAngle := angle(Cart{0, 0}, *c)
	p.Distance = cartDist
	p.Angle = cartAngle
	return p
}

// Point2PointDistance distance in the cartesian coordinate system
func (c *Cart) Point2PointDistance(point Cart) (p float64) {
	return distance(*c, point)
}

// distance for point of Cart types
func distance(pointOne, pointTwo Cart) (dist float64) {
	changeInEastings := pointTwo.X - pointOne.X
	changeInNorthings := pointTwo.Y - pointOne.Y
	dist = math.Sqrt(math.Pow(changeInEastings, 2.0) + math.Pow(changeInNorthings, 2.0))
	return dist
}

// angle/bearing for point of Cart types
func angle(pointOne, pointTwo Cart) (ang float64) {
	changeInEastings := pointTwo.X - pointOne.X
	changeInNorthings := pointTwo.Y - pointOne.Y
	ang = math.Atan(changeInEastings / changeInNorthings)
	return ang
}
