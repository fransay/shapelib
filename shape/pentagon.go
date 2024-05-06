package shape

import "math"

type Pentagon struct {
	Side float64
}

const interiorAngle = 105
const exteriorAngle = 108

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
