package cartesian

import (
	"shapelib/types"
)

type Cartesian2D struct {
	X           Axis
	Y           Axis
	OriginPoint types.Point2D
}

// NewCartesian2D initialises a new Cartesian2D
func (c *Cartesian2D) NewCartesian2D(X Axis, Y Axis) *Cartesian2D {
	return &Cartesian2D{
		X: X,
		Y: Y,
	}
}

// XValues return x-only values in an array
func (c *Cartesian2D) XValues() (arr []float64) {
	for value := c.X.Start; value <= c.X.End; value += c.X.Step {
		arr = append(arr, value)
	}
	return arr
}

// YValues return y-only values in an array
func (c *Cartesian2D) YValues() (arr []float64) {
	for value := c.Y.Start; value <= c.Y.End; value += c.Y.Step {
		arr = append(arr, value)
	}
	return arr
}

// Origin return the origin of the coordinate system.
func (c *Cartesian2D) Origin() types.Point2D {
	if &c.OriginPoint != nil {
		return c.OriginPoint
	}
	return types.Point2D{X: c.X.Start, Y: c.Y.Start}
}

// PointInCartesian checks if a point lies in the cartesian vector space
func (c *Cartesian2D) PointInCartesian(point types.Point2D) bool {
	if (point.X < c.X.Start || point.X > c.X.End) || (point.Y < c.Y.Start || point.Y > c.Y.End) {
		return false
	}
	return true
}

// Quadrant return the particular quadrant as typically seen in cartesian plot
//
//		   	|
//		 2	|  1
//	 ----------------
//		 3  |  4
//			|
func (c *Cartesian2D) Quadrant() (quadrant int) {
	if c.X.Start > 0 && c.Y.Start > 0 {
		quadrant = 1
	} else if c.X.Start > 0 && c.Y.Start < 0 {
		quadrant = 2
	} else if c.X.Start < 0 && c.Y.Start < 0 {
		quadrant = 3
	} else if c.X.Start < 0 && c.Y.Start > 0 {
		quadrant = 4
	}
	return quadrant
}

// Dim returns the dimension of the coordinate system.
func (c *Cartesian2D) Dim() float64 { return 2 }
