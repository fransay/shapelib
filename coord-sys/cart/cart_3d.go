package cart

import (
	"shapelib/types/point"
)

type Cart3d struct {
	X           Axis
	Y           Axis
	Z           Axis
	OriginPoint point.Point3D
}

func (c *Cart3d) Init(X, Y, Z Axis) Cart3d {
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

func (c *Cart3d) Origin() point.Point3D {
	if &c.OriginPoint != nil {
		return c.OriginPoint
	}
	return point.Point3D{X: c.X.Start, Y: c.Y.Start, Z: c.Z.Start}
}
