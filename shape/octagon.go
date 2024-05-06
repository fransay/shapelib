package shape

import "math"

type Octagon struct {
	Side float64
}

// Area returns the area of octagon
func (o *Octagon) Area() float64 {
	return 2 * (1 + math.Sqrt(o.Side)) * math.Pow(o.Side, 2)
}

// Perimeter returns the perimeter of octagon
func (o *Octagon) Perimeter() float64 {
	return o.Side * 8
}
