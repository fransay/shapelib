package cartesian

import "math"

type Cart3D struct {
	X float64
	Y float64
	Z float64
}

// Init constructs a new Cart3D object
func (c *Cart3D) Init(x float64, y float64, z float64) *Cart3D {
	return &Cart3D{
		X: x,
		Y: y,
		Z: z,
	}
}

// Distance3D returns Euclidean distance in a three-dimensional space
func (c *Cart3D) Distance3D(c2 *Cart3D) float64 {
	return math.Sqrt((c2.X - c.X) + (c2.Y - c.Y) + (c2.Z - c.Z))
}
