package shape

import (
	"math"

	"github.com/fransay/shapelib/internal/utils"
)

// Hexagon type
type Hexagon struct {
	Side float64
}

// Area returns the area of regular hexagon
func (h *Hexagon) Area() float64 {
	return ((3 * math.Sqrt(3)) / 2) * math.Pow(h.Side, 2)
}

// Perimeter returns the perimeter of regular hexagon
func (h *Hexagon) Perimeter() float64 {
	return h.Side * 6
}

// Apothem returns the apothem of a regular polygon
func (h *Hexagon) Apothem() float64 {
	return (h.Side / 2) * utils.Cot(math.Pi/h.Side)
}
