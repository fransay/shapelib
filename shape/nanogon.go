package shape

import (
	"math"
	"shapelib/utils"
)

type Nanogon struct {
	side float64
}

// InsCircle returns the radius of the inscribed circle
func (n *Nanogon) InsCircle() (radius float64) {
	radius = (n.side / 2) * utils.Cot(math.Pi/9)
	return radius
}

// CircumCircle returns the radius of the cirumcircle
func (n *Nanogon) CircumCircle() (radius float64) {
	radius = math.Sqrt(math.Pow(n.side, 2) + math.Pow(n.InsCircle(), 2))
	return radius
}

// Area returns area of nanogon
func (n *Nanogon) Area() float64 {
	const angle = math.Pi / 9
	return 9 * math.Pow(n.side, 2) * utils.Cot(angle)
}

// Perimeter returns perimeter of nanogon
func (n *Nanogon) Perimeter() float64 {
	return n.side * 9
}
