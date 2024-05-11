package shape

import (
	"math"
	"shapelib/utils"
)

type Nanogon struct {
	side float64
}

// Area returns area of nanogon
func (n *Nanogon) Area() float64 {
	const angle = math.Pi / 9
	return 9 * math.Pow(n.side, 2) * utils.Cot(angle)
}
