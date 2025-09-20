package shape

import (
	"math"

	"github.com/fransay/shapelib/internal/utils"
)

type Nonagon struct {
	Side float64
}

func NewNanogon(side float64) *Nonagon {
	return &Nonagon{side}
}

// InsCircle returns the radius of the inscribed circle
func (n *Nonagon) InsCircle() (radius float64) {
	radius = (n.Side / 2) * utils.Cot(math.Pi/9)
	return radius
}

// CircumCircle returns the radius of the cirumcircle
func (n *Nonagon) CircumCircle() (radius float64) {
	radius = math.Sqrt(math.Pow(n.Side, 2) + math.Pow(n.InsCircle(), 2))
	return radius
}

// Area returns area of nanogon
func (n *Nonagon) Area() float64 {
	const angle = math.Pi / 9
	return 9 * math.Pow(n.Side, 2) * utils.Cot(angle)
}

// Perimeter returns perimeter of nanogon
func (n *Nonagon) Perimeter() float64 {
	return n.Side * 9
}
