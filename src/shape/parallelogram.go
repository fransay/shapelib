package shape

// Parallelogram type
type Parallelogram struct {
	Base   float64
	Height float64
	Side   float64
}

// NewParallelogram initialises a new parallelogram object
func NewParallelogram(base, height, side float64) *Parallelogram {
	return &Parallelogram{
		Base:   base,
		Height: height,
		Side:   side,
	}
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

// IsQuad is a placeholder method for defining the quadrilateral interface
func (p *Parallelogram) IsQuad() bool {
	return true
}
