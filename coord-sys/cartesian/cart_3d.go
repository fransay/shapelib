package cartesian

type Cart3D struct {
	X float64
	Y float64
	Z float64
}

// Init constructs new Cart3D objects
func (c *Cart3D) Init(x float64, y float64, z float64) *Cart3D {
	return &Cart3D{
		X: x,
		Y: y,
		Z: z,
	}
}
