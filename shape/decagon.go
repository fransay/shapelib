package shape

import (
	"math"
	"shapelib/utils"
)

type Decagon struct {
	side float64
}

// Area returns area of a decagon
func (d *Decagon) Area() (area float64) {
	area = (5 / 2) * math.Pow(d.side, 2) * utils.Cot(math.Pi/10)
	return area
}

// Perimeter returns the perimeter of a decagon
func (d *Decagon) Perimeter() (perimeter float64) {
	return d.side * 10
}
