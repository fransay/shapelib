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

func (c *Cart3D) Bearing3D(c2 *Cart3D) float64 {
	deltaX := c.X - c2.X
	deltaY := c.Y - c2.Y
	var bearing float64
	if deltaY > 0 {
		bearing = math.Atan(deltaX / deltaY)
	} else if deltaX < 0 {
		bearing = math.Atan(deltaX/deltaY) + 180
	} else if deltaY == 0 && deltaX > 0 {
		bearing = 90.0
	} else if deltaY == 0 && deltaX < 0 {
		bearing = (3 * 180) / 2
	}
	normBearing := int(bearing+360.0) % 360.0 // normalised bearing
	return normBearing
}

// AddCart3D returns the sum self and other Cart3D
func (c *Cart3D) AddCart3D(c2 *Cart3D) Cart3D {
	return Cart3D{X: c.X + c2.X, Y: c.Y + c2.Y, Z: c.Z + c2.Z}
}

// SubCart3D returns the difference self and other Cart3D
func (c *Cart3D) SubCart3D(c2 *Cart3D) Cart3D {
	return Cart3D{X: c2.X - c.X, Y: c2.Y - c.Y, Z: c2.Z - c.Z}
}

// MulCart3D returns the product between self and other Cart3D
func (c *Cart3D) MulCart3D(c2 *Cart3D) Cart3D {
	return Cart3D{X: c.X * c2.X, Y: c.Y * c2.Y, Z: c.Z * c2.Z}
}

// MulScalar3D returns the product between self and scalar
func (c *Cart3D) MulScalar3D(scalar float64) Cart3D {
	return Cart3D{c.X * scalar, c.Y * scalar, c.Z * scalar}
}

// DivCart3D return the quotient between self and other
func (c *Cart3D) DivCart3D(c2 *Cart3D) Cart3D {
	return Cart3D{X: c2.X / c.X, Y: c2.Y / c.Y, Z: c2.Z / c.Z}
}

// DivScalar returns the quotient between self and other scalar
func (c *Cart3D) DivScalar(scalar float64) Cart3D {
	return Cart3D{X: c.X * scalar, Y: c.Y * scalar, Z: c.Z * scalar}
}
