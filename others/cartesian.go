package others

import (
	"math"
	"shapelib/coord-sys/polar"
	"shapelib/utils"
)

// Cart2D cartesian type for 2-dim
type Cart2D struct {
	X float64
	Y float64
}

// Cart3D cartesian type for 3-dim
type Cart3D struct {
	X float64
	Y float64
	Z float64
}

// ToPolar convert cartesian to Polar coordinates return multiple objects,
// angle in radians and degrees origin (0,0):: reference point
func (c *Cart2D) ToPolar() (polarRadians, polarDegrees polar.Polar) {
	cartDist := distance(Cart2D{0, 0}, *c)
	cartAngle := angle(Cart2D{0, 0}, *c)
	polarRadians = polar.Polar{Distance: cartDist, Angle: cartAngle}
	polarDegrees = polar.Polar{Distance: cartDist, Angle: utils.Rad2Deg(cartAngle)}
	return polarRadians, polarDegrees
}

// Point2PointDistance2D distance in the cartesian coordinate system
func (c *Cart2D) Point2PointDistance2D(point Cart2D) (p float64) {
	return distance(*c, point)
}

// distance for points of Cart2D types
func distance(pointOne, pointTwo Cart2D) (dist float64) {
	changeInEastings := pointTwo.X - pointOne.X
	changeInNorthings := pointTwo.Y - pointOne.Y
	dist = math.Sqrt(
		math.Pow(changeInEastings, 2.0) + math.Pow(changeInNorthings, 2.0))
	return dist
}

// angle/bearing for point of Cart2D types
func angle(pointOne, pointTwo Cart2D) (ang float64) {
	changeInEastings := pointTwo.X - pointOne.X
	changeInNorthings := pointTwo.Y - pointOne.Y
	ang = math.Atan(changeInNorthings / changeInEastings)
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

// Rotate2D cartesian 2D point around the x-axis
func (c *Cart2D) Rotate2D(angle float64) (rotate Cart2D) {
	xPrime := c.X*math.Cos(angle) - c.Y*math.Sin(angle)
	yPrime := c.X*math.Sin(angle) - c.Y*math.Cos(angle)
	rotate.X = xPrime
	rotate.Y = yPrime
	return rotate
}

// Rotate3D cartesian 3D point around the origin
func (c *Cart3D) Rotate3D(rotationAngleAroundY, rotationAngleAroundZ float64) (rotate Cart3D) {
	xPrime := c.X*math.Cos(rotationAngleAroundY)*
		math.Cos(rotationAngleAroundZ) - c.Y*math.Cos(rotationAngleAroundY)*
		math.Sin(rotationAngleAroundZ) + c.Z*math.Sin(rotationAngleAroundY)
	yPrime := c.X*math.Sin(rotationAngleAroundZ) +
		c.Y*math.Cos(rotationAngleAroundZ) + c.Z*math.Sin(rotationAngleAroundY)
	zPrime := -1*c.X*math.Sin(rotationAngleAroundY)*
		math.Cos(rotationAngleAroundZ) + c.Y*math.Sin(rotationAngleAroundZ) +
		c.Z*math.Cos(rotationAngleAroundY)
	rotate.X = xPrime
	rotate.Y = yPrime
	rotate.Z = zPrime
	return rotate
}
