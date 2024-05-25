package cart

import (
	"shapelib/types/point"
)

type Cartesian2D struct {
	X           Axis
	Y           Axis
	OriginPoint point.Point2D
}

func (c *Cartesian2D) Init(X, Y Axis) Cartesian2D {
	return Cartesian2D{
		X: X,
		Y: Y,
	}
}

func (c *Cartesian2D) XValues() (arr []float64) {
	for value := c.X.Start; value <= c.X.End; value += c.X.Step {
		arr = append(arr, value)
	}
	return arr
}

func (c *Cartesian2D) YValues() (arr []float64) {
	for value := c.Y.Start; value <= c.Y.End; value += c.Y.Step {
		arr = append(arr, value)
	}
	return arr
}

func (c *Cartesian2D) Origin() point.Point2D {
	if &c.OriginPoint != nil {
		return c.OriginPoint
	}
	return point.Point2D{X: c.X.Start, Y: c.Y.Start}
}
