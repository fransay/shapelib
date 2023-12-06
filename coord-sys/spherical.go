package coord_sys

import "math"

// Spherical type
type Spherical struct {
	RadialDistance float64
	PolarAngle     float64 // in radians
	AzimuthAngle   float64 // in radians
}

// ToCartesianIn2D converts spherical coordinate to cartesian coordinates in 3D
func (s *Spherical) ToCartesianIn2D() (cart Cart) {
	x := s.RadialDistance * math.Sin(s.AzimuthAngle) * math.Cos(s.PolarAngle)
	y := s.RadialDistance * math.Sin(s.AzimuthAngle) * math.Sin(s.PolarAngle)
	cart.X = x
	cart.Y = y
	return cart
}

// ToCartesianIn3D converts spherical coordinate to cartesian coordinates in 3D
func (s *Spherical) ToCartesianIn3D() (cart Cart3D) {
	x := s.RadialDistance * math.Sin(s.AzimuthAngle) * math.Cos(s.PolarAngle)
	y := s.RadialDistance * math.Sin(s.AzimuthAngle) * math.Sin(s.PolarAngle)
	z := s.RadialDistance * math.Cos(s.PolarAngle)
	cart.X = x
	cart.Y = y
	cart.Z = z
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
