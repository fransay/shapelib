package shape

import "math"

type Hexagon struct {
	side float64
}

// Area returns the area of regular hexagon
func (h *Hexagon) Area() float64 {
	return ((3 * math.Sqrt(3)) / 2) * math.Pow(h.side, 2)
}

// Perimeter returns the perimeter of regular hexagon
func (h *Hexagon) Perimeter() float64 {
	return ((3 * math.Sqrt(3)) / 2) * math.Pow(h.side, 2)
}

//
