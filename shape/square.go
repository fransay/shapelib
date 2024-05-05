package shape

import "math"

type Square struct {
	length float64
}

const InteriorAngle = 90.0

// Area returns the area of a square
func (s *Square) Area() (area float64) {
	area = s.length * s.length
	return area
}

// Perimeter returns the perimeter of a square
func (s *Square) Perimeter() (perimeter float64) {
	perimeter = 2*(s.length) + 2*(s.length)
	return perimeter
}

// Diagional returns the diagonal length of a square
func (s *Square) Diagional() (diag float64) {
	diag = s.length * math.Sqrt(2)
	return diag
}

// CircumRadius returns the diagonal length of a square
func (s *Square) CircumRadius() (circumRad float64) {
	circumRad = s.length * (math.Sqrt(2))
	return circumRad
}

// InRadius returns the radius of the incirle
func (s *Square) InRadius() (inRadius float64) {
	inRadius = s.length / 2
	return inRadius
}

// Apothem returns the distance from any point on the square to
// the nearest side of the square.
func (s *Square) Apothem() (apothem float64) {
	apothem = s.length / 2
	return apothem
}

// GoldenRation returns the ratio of the diagonal length to the
// side length of a square.
func (s *Square) GoldenRation() (goldenRatio float64) {
	diag := s.length * math.Sqrt(2)
	goldenRatio = diag / s.length
	return goldenRatio
}
