package cart

import (
	"shapelib/types/point"
)

// Axis defines an axis in a cartesian plane
type Axis struct {
	Start float64
	End   float64
	Step  float64
}

type Cart2d struct {
	X           Axis
	Y           Axis
	OriginPoint point.Point2D
}

func (c *Cart2d) Init(X, Y Axis) Cart2d {
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

func (c *Cart2d) Origin() point.Point2D {
	if &c.OriginPoint != nil {
		return c.OriginPoint
	}
	return point.Point2D{X: c.X.Start, Y: c.Y.Start}
}
