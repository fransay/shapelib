package shape

type Parallelogram struct {
	Base   float64
	Height float64
	Side   float64
}

// AreaByBaseHeight return area of parallelogram
func (p *Parallelogram) AreaByBaseHeight() (area float64) {
	area = p.Base * p.Height
	return area
}

// Perimeter return perimeter of a parallelogram
func (p *Parallelogram) Perimeter() (perimeter float64) {
	perimeter = 2 * (p.Base + p.Side)
	return perimeter
}
