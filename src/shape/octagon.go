package shape

import "math"

// Octagon type
type Octagon struct {
	Side float64
}

// NewOctagon inits a new octagon object
func NewOctagon(side float64) *Octagon {
	return &Octagon{Side: side}
}

// Area returns the area of octagon
func (o *Octagon) Area() float64 {
	return 2 * (1 + math.Sqrt(o.Side)) * math.Pow(o.Side, 2)
}

// Perimeter returns the perimeter of octagon
func (o *Octagon) Perimeter() float64 {
	return o.Side * 8
}

// Apothem returns the distance from the center of octagon to
// the midpoint on any Side.
func (o *Octagon) Apothem() float64 {
	return (o.Side / 2) * (math.Sqrt(2 + math.Sqrt(2)))
}

// InteriorDiagonal returns the length of the inner diameter
func (o *Octagon) InteriorDiagonal() float64 {
	return o.Side * math.Sqrt(4+(2*math.Sqrt2))
}
