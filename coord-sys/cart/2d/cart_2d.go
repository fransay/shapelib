package _d

import "shapelib/coord-sys/cart"

type Cart2d struct {
	X cart.Axis
	Y cart.Axis
}

func (c *Cart2d) Init(X, Y cart.Axis) Cart2d {
	return Cart2d{
		X: X,
		Y: Y,
	}
}

func (c *Cart2d) XValues() (arr []float64) {
	for value := c.X.Start; value <= c.X.End; value += c.X.Step {
		arr = append(arr, value)
	}
	return arr
}

func (c *Cart2d) YValues() (arr []float64) {
	for value := c.Y.Start; value <= c.Y.End; value += c.Y.Step {
		arr = append(arr, value)
	}
	return arr
}
