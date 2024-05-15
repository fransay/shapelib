package shape

import (
	"math"
	"shapelib/utils"
)

type Heptagon struct {
	side float64
}

const numberOfSide = 7

// Area returns the area of a heptagon in unit squares
func (h *Heptagon) Area() (area float64) {
	area = (7 * (h.side * h.side)) / 4 * utils.Cot(math.Pi/7)
	return area
}

// Perimeter returns the area of a heptagon in unit squares
func (h *Heptagon) Perimeter() (perimeter float64) {
	perimeter = h.side * numberOfSide
	return perimeter
}

// Apothem returns the distance from the center of the hexagon
// to the midpoint on any side
func (h *Heptagon) Apothem() float64 {
	return h.side / (2 * math.Tan(math.Pi/7))
}
