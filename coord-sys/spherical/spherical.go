package spherical

import (
	"math"
	"shapelib/others"
)

// Spherical type
type Spherical struct {
	RadialDistance float64
	PolarAngle     float64 // in radians
	AzimuthAngle   float64 // in radians
}

// ToCartesianIn2D converts spherical coordinate to cartesian coordinates in 3D
func (s *Spherical) ToCartesianIn2D() (cart others.Cart2D) {
	x := s.RadialDistance * math.Sin(s.AzimuthAngle) * math.Cos(s.PolarAngle)
	y := s.RadialDistance * math.Sin(s.AzimuthAngle) * math.Sin(s.PolarAngle)
	cart = others.Cart2D{X: x, Y: y}
	return cart
}

// ToCartesianIn3D converts spherical coordinate to cartesian coordinates in 3D
func (s *Spherical) ToCartesianIn3D() (cart others.Cart3D) {
	x := s.RadialDistance * math.Sin(s.AzimuthAngle) * math.Cos(s.PolarAngle)
	y := s.RadialDistance * math.Sin(s.AzimuthAngle) * math.Sin(s.PolarAngle)
	z := s.RadialDistance * math.Cos(s.PolarAngle)
	cart = others.Cart3D{X: x, Y: y, Z: z}
	return cart
}

// Point2PointDistance returns distance two spherical points using spherical laws of cosine
func (s *Spherical) Point2PointDistance(point Spherical) (dist float64) {
	addRadialDist := s.RadialDistance + point.RadialDistance
	prodRadialDist := s.RadialDistance * point.RadialDistance
	angArithmetic := math.Sin(s.AzimuthAngle)*math.Sin(point.AzimuthAngle)*math.Cos(s.PolarAngle-point.PolarAngle) +
		math.Cos(s.AzimuthAngle)*math.Cos(point.AzimuthAngle)
	dist = math.Sqrt(addRadialDist - 2*prodRadialDist*angArithmetic)
	return dist
}

// ToArray return the array form of the spherical structure
// array form: {azimuthAngle, polarAngle, radialDistance}
func (s *Spherical) ToArray() [3]float64 {
	return [3]float64{s.RadialDistance, s.PolarAngle, s.AzimuthAngle}
}
