package cart

import (
	"shapelib/types"
)

type Cartesian2D struct {
	X           Axis
	Y           Axis
	OriginPoint types.Point2D
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

func (c *Cartesian2D) Origin() types.Point2D {
	if &c.OriginPoint != nil {
		return c.OriginPoint
	}
	return types.Point2D{X: c.X.Start, Y: c.Y.Start}
}

// PointInCart checks if a point lie the range of self vector space
func (c *Cartesian2D) PointInCart() {}
