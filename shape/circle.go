package shape

import (
	"math"
	"shapelib/types/point"
)

type Circle struct {
	Diameter float64
	Radius   float64
	Chord    float64
	Centroid point.Point2D
}

// Area returns the area of a circle
func (c *Circle) Area() (area float64) {
	area = math.Pi * math.Pow(c.Radius, 2)
	return area
}

// Circumference returns the circumference of a circle
func (c *Circle) Circumference() (circumference float64) {
	circumference = 2 * math.Pi * c.Radius
	return circumference
}

// ComputeChord returns the chord length of circle based on sagitta
func (c *Circle) ComputeChord(sagitta float64) (chord float64) {
	chord = 2 * math.Sqrt(math.Pow(c.Radius, 2)-math.Pow(sagitta, 2))
	return chord
}
