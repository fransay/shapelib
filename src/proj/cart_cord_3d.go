package proj

import "github.com/fransay/shapelib/src/geom"

type Cartesian3D struct {
	X           Axis
	Y           Axis
	Z           Axis
	OriginPoint geom.Point3D
}

// Init initialises a new Cartesian3D object
func (c *Cartesian3D) Init(X, Y, Z Axis) *Cartesian3D {
	return &Cartesian3D{
		X: X,
		Y: Y,
		Z: Z,
	}
}

// XValues return all values on x-axis
func (c *Cartesian3D) XValues() (arr []float64) {
	for value := c.X.Start; value <= c.X.End; value += c.X.Step {
		arr = append(arr, value)
	}
	return arr
}

// YValues return all values on y-axis
func (c *Cartesian3D) YValues() (arr []float64) {
	for value := c.Y.Start; value <= c.Y.End; value += c.Y.Step {
		arr = append(arr, value)
	}
	return arr
}

// ZValues return all values on z-axis
func (c *Cartesian3D) ZValues() (arr []float64) {
	for value := c.Z.Start; value <= c.Z.End; value += c.Z.Step {
		arr = append(arr, value)
	}
	return arr
}

// Origin returns the origin of Cartesian3D
func (c *Cartesian3D) Origin() geom.Point3D {
	return geom.Point3D{X: c.X.Start, Y: c.Y.Start, Z: c.Z.Start}
}

// Dim returns the dimension of Cartesian3D
func (c *Cartesian3D) Dim() int { return 3 }
