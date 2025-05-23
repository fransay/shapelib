package shape

import (
	"errors"
	"math"

	"github.com/fransay/shapelib/internal/utils"
)

// Heptagon type
type Heptagon struct {
	Side float64
}

// NewHeptagon initialises a new pentagon object
func NewHeptagon(side float64) (*Heptagon, error) {
	if side < 0 {
		return nil, errors.New("can't create heptagon object, Side length can't be les than zero")

	}
	return &Heptagon{Side: side}, nil

}

// Area returns the area of a heptagon in unit squares
func (h *Heptagon) Area() (area float64) {
	area = (7 * (h.Side * h.Side)) / 4 * utils.Cot(math.Pi/7)
	return area
}

// Perimeter returns the area of a heptagon in unit squares
func (h *Heptagon) Perimeter() (perimeter float64) {
	perimeter = h.Side * 7
	return perimeter
}

// Apothem returns the distance from the center of the hexagon to the midpoint on any Side
func (h *Heptagon) Apothem() float64 {
	return h.Side / (2 * math.Tan(math.Pi/7))
}
