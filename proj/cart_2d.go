package proj

import (
	"math"
)

type Cart2D struct {
	X float64
	Y float64
}

// Init constructs new Cart2D objects
func (c *Cart2D) Init(x, y float64) *Cart2D {
	return &Cart2D{X: x, Y: y}
}

// ToPolar returns the equivalent polar coordinate, defaults origin
// to Cart2D{X: 0, Y: 0}
func (c *Cart2D) ToPolar() Polar {
	distance := math.Sqrt(c.X*c.X + c.Y*c.Y)
	angle := math.Atan(c.Y / c.X) // angle is bearing
	return Polar{Distance: distance, Angle: angle}
}

// AddCart2D returns the sum of this.Cart2D and other.Cart2D
func (c *Cart2D) AddCart2D(cart2D Cart2D) Cart2D {
	return Cart2D{X: cart2D.X + c.X, Y: cart2D.Y + c.Y}
}

// SubCart2D returns the difference of this.Cart2D and other.Cart2D
func (c *Cart2D) SubCart2D(cart2d *Cart2D) Cart2D {
	return Cart2D{X: cart2d.X - c.X, Y: cart2d.Y - c.Y}
}

// MulCart2D returns the product of this.Cart2D and other.Cart2D
func (c *Cart2D) MulCart2D(cart2d *Cart2D) Cart2D {
	return Cart2D{X: cart2d.X * c.X, Y: cart2d.Y * c.Y}
}

// MulScalar returns the product of this.Cart2D and a scalar
func (c *Cart2D) MulScalar(scalar float64) Cart2D {
	return Cart2D{X: scalar * c.X, Y: scalar * c.Y}
}

// DivCart2D returns the quotient of this.Cart2D and other.Cart2D
func (c *Cart2D) DivCart2D(cart2d *Cart2D) Cart2D {
	return Cart2D{X: cart2d.X / c.X, Y: cart2d.Y / c.Y}
}

// DivScalar returns the quotient of this.Cart2D and a scalar
func (c *Cart2D) DivScalar(scalar float64) Cart2D {
	return Cart2D{X: scalar / c.X, Y: scalar / c.Y}
}

// Distance returns the Euclidean distance between this.Cart2D and other.Cart2D
// in a cartesian coordinate system.
func (c *Cart2D) Distance(c2 *Cart2D) float64 {
	return math.Sqrt(math.Pow(c2.X-c.X, 2) + math.Pow(c2.Y-c.Y, 2))
}

// Bearing return the angular difference geog north and the line of interest
func (c *Cart2D) Bearing(c2 *Cart2D) float64 {
	return math.Atan(c2.Y - c.Y/c2.X - c.X)
}
