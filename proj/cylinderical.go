package proj

import (
	"math"
)

type Cylinderical struct {
	Distance float64 // Radial distance
	Angle    float64 // Azimuth angle
	Height   float64 // Height above the xy plane
}

// Init return a new cylinderical object
func (c *Cylinderical) Init(distance, angle, height float64) *Cylinderical {
	return &Cylinderical{
		Distance: distance,
		Angle:    angle,
		Height:   height,
	}
}

// ToCartesian returns a cartesian equivalent of cylinderical representations
func (c *Cylinderical) ToCartesian() Cart3D {
	return Cart3D{
		X: c.Distance * math.Cos(c.Angle),
		Y: c.Distance * math.Sin(c.Angle),
		Z: c.Height,
	}
}
