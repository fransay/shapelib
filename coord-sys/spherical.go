package coord_sys

// Spherical type
type Spherical struct {
	RadialDistance float64
	PolarAngle     float64 // in degrees
	AzimuthAngle   float64 // in degrees
}

// ToCartesian converts spherical coordinate to cartesian coordinates
func (s *Spherical) ToCartesian() (cart Cart) {
	return cart
}

// Point2PointDistance returns distance two spherical points
// using spherical laws of cosine
func (s *Spherical) Point2PointDistance(point Spherical) (dist float64) {
	return dist
}
