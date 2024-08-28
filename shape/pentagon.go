package shape

import "math"

type Pentagon struct {
	Side float64
}

const pentagonInteriorAngle = 105
const pentagonExteriorAngle = 108

// Area return area of pentagon
func (p *Pentagon) Area() (area float64) {
	area = 0.25 * math.Sqrt(5+(5+2*math.Sqrt(5))) * math.Pow(p.Side, 2)
	return area
}

// Perimeter return perimeter of pentagon
func (p *Pentagon) Perimeter() (perimeter float64) {
	perimeter = p.Side * 5
	return perimeter
}

// Diagonal return the Length of diagonal of pentagon
func (p *Pentagon) Diagonal() (diagonal float64) {
	diagonal = p.Side * math.Sqrt(5-2*math.Sqrt(5))
	return diagonal
}

// Circumcircle return the circumference of a circumcircle drawn around the pentagon
func (p *Pentagon) Circumcircle() float64 {
	sqrt5 := math.Sqrt(5)
	r := (p.Side / 2) * math.Sqrt((5+sqrt5)/2)
	return r
}

// Incircle return the circumference of an incircle drawn in a pentagon
func (p *Pentagon) Incircle() float64 {
	sqrt5 := math.Sqrt(5)
	r := (p.Side / 2) * math.Sqrt((5-sqrt5)/2)
	return r
}
