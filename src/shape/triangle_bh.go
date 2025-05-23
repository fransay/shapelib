package shape

// TriangleByBaseHeight
type TriangleByBaseHeight struct {
	Base   float64
	Height float64
}

func NewTriangleByBaseHeight(base, height float64) *TriangleByBaseHeight {
	return &TriangleByBaseHeight{
		Base:   base,
		Height: height,
	}
}

// Area returns the area of a triangle given the Base and Height
func (t *TriangleByBaseHeight) Area() (areaByBaseHeight float64) {
	areaByBaseHeight = 0.5 * t.Base * t.Height
	return areaByBaseHeight
}

// Perimeter returns the total distance around a given triangle given length of all three sides, i.e a, b, c
func (t *TriangleByBaseHeight) Perimeter(a, b, c float64) float64 {
	return a + b + c
}
