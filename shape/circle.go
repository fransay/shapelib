package shape

import "math"

type Circle struct {
	Diameter float64
	Radius   float64
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
