package cart

import (
	"shapelib/types"
)

type Cartesian3D struct {
	X           Axis
	Y           Axis
	Z           Axis
	OriginPoint types.Point3D
}

func (c *Cartesian3D) Init(X, Y, Z Axis) Cartesian3D {
	return Cartesian3D{
		X: X,
		Y: Y,
		Z: Z,
	}
}

func (c *Cartesian3D) XValues() (arr []float64) {
	for value := c.X.Start; value <= c.X.End; value += c.X.Step {
		arr = append(arr, value)
	}
	return arr
}

func (c *Cartesian3D) YValues() (arr []float64) {
	for value := c.Y.Start; value <= c.Y.End; value += c.Y.Step {
		arr = append(arr, value)
	}
	return arr
}

func (c *Cartesian3D) ZValues() (arr []float64) {
	for value := c.Z.Start; value <= c.Z.End; value += c.Z.Step {
		arr = append(arr, value)
	}
	return arr
}

func (c *Cartesian3D) Origin() types.Point3D {
	if &c.OriginPoint != nil {
		return c.OriginPoint
	}
	return types.Point3D{X: c.X.Start, Y: c.Y.Start, Z: c.Z.Start}
}
