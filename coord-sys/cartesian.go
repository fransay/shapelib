package coord_sys

import "math"

// Cart2D cartesian type
type Cart2D struct {
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
func (c *Cart2D) ToPolar() (p Polar) {
	cartDist := distance(Cart2D{0, 0}, *c)
	cartAngle := angle(Cart2D{0, 0}, *c)
	p.Distance = cartDist
	p.Angle = cartAngle
	return p
}

// Point2PointDistance2D distance in the cartesian coordinate system
func (c *Cart2D) Point2PointDistance2D(point Cart2D) (p float64) {
	return distance(*c, point)
}

// distance for point of Cart2D types
func distance(pointOne, pointTwo Cart2D) (dist float64) {
	changeInEastings := pointTwo.X - pointOne.X
	changeInNorthings := pointTwo.Y - pointOne.Y
	dist = math.Sqrt(math.Pow(changeInEastings, 2.0) + math.Pow(changeInNorthings, 2.0))
	return dist
}

// angle/bearing for point of Cart2D types
func angle(pointOne, pointTwo Cart2D) (ang float64) {
	changeInEastings := pointTwo.X - pointOne.X
	changeInNorthings := pointTwo.Y - pointOne.Y
	ang = math.Atan(changeInEastings / changeInNorthings)
	return ang
}

// Translate2D cartesian 2D point by a vector
func (c *Cart2D) Translate2D(vec [2]float64) (transCord Cart2D) {
	xPrime := c.X + vec[0]
	yPrime := c.Y + vec[1]
	transCord.X = xPrime
	transCord.Y = yPrime
	return transCord
}

// Translate3D cartesian 3D point by a vector
func (c *Cart3D) Translate3D(vec [3]float64) (transCord Cart3D) {
	xPrime := c.X + vec[0]
	yPrime := c.Y + vec[1]
	zPrime := c.Z + vec[2]
	transCord.X = xPrime
	transCord.Y = yPrime
	transCord.Z = zPrime
	return transCord
}

//// Rotate2D cartesian 2D point by a vector
//func (c *Cart2D) Rotate2D() {
//
//}
//
//// Reflect2D cartesian 2D point by a vector
//func (c *Cart2D) Reflect2D() {
//
//}
//
//
//// Rotate3D cartesian 3D point by a vector
//func (c *Cart2D) Rotate3D() {
//
//}
//
//// Reflect3D cartesian 3D point by a vector
//func (c *Cart2D) Reflect3D() {
//
//}
