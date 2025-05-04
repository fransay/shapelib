package shape

import "math"

// Pentagon type
type Pentagon struct {
	Side float64 // length of side
}

// NewPentagon initialise a new pentagon object
func NewPentagon(side float64) *Pentagon {
	return &Pentagon{Side: side}
}

// Area return area of pentagon
func (p *Pentagon) Area() (area float64) {
	area = math.Sqrt(5*(5+(2*math.Sqrt(5)))) / 4 * (math.Pow(p.Side, 2))
	return area
}

// Perimeter return perimeter of pentagon
func (p *Pentagon) Perimeter() (perimeter float64) {
	perimeter = p.Side * 5
	return perimeter
}

// Diagonal return the length of diagonal of pentagon
func (p *Pentagon) Diagonal() (diagonal float64) {
	diagonal = p.Side * (1 + math.Sqrt(5)) / 2
	return diagonal
}

// Circumcircle return the circumference of a circumcircle drawn around the pentagon
func (p *Pentagon) Circumcircle() float64 {
	return p.Side / 1.17557
}

// Incircle return the circumference of an inner circle drawn in a pentagon
func (p *Pentagon) Incircle() float64 {
	return p.Side / 1.4531
}
