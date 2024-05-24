package _d

import "shapelib/coord-sys/cart"

type Cart3d struct {
	X cart.Axis
	Y cart.Axis
	Z cart.Axis
}

func (c *Cart3d) Init(X, Y, Z cart.Axis) Cart3d {
	return Cart3d{
		X: X,
		Y: Y,
		Z: Z,
	}
}

func (c *Cart3d) XValues() (arr []float64) {
	for value := c.X.Start; value <= c.X.End; value += c.X.Step {
		arr = append(arr, value)
	}
	return arr
}

func (c *Cart3d) YValues() (arr []float64) {
	for value := c.Y.Start; value <= c.Y.End; value += c.Y.Step {
		arr = append(arr, value)
	}
	return arr
}

func (c *Cart3d) ZValues() (arr []float64) {
	for value := c.Z.Start; value <= c.Z.End; value += c.Z.Step {
		arr = append(arr, value)
	}
	return arr
}
