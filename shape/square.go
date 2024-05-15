package shape

import "math"

type Square struct {
	Length float64
}

const InteriorAngle = 90.0

// Area returns the area of a square
func (s *Square) Area() (area float64) {
	area = s.Length * s.Length
	return area
}

// Perimeter returns the perimeter of a square
func (s *Square) Perimeter() (perimeter float64) {
	perimeter = 2*(s.Length) + 2*(s.Length)
	return perimeter
}

// Diagional returns the diagonal Length of a square
func (s *Square) Diagional() (diag float64) {
	diag = s.Length * math.Sqrt(2)
	return diag
}

// CircumRadius returns the diagonal Length of a square
func (s *Square) CircumRadius() (circumRad float64) {
	circumRad = s.Length * (math.Sqrt(2))
	return circumRad
}

// InRadius returns the radius of the incirle
func (s *Square) InRadius() (inRadius float64) {
	inRadius = s.Length / 2
	return inRadius
}

// Apothem returns the distance from any point on the square to
// the nearest Side of the square.
func (s *Square) Apothem() (apothem float64) {
	apothem = s.Length / 2
	return apothem
}

// GoldenRation returns the ratio of the diagonal Length to the
// Side Length of a square.
func (s *Square) GoldenRation() (goldenRatio float64) {
	diag := s.Length * math.Sqrt(2)
	goldenRatio = diag / s.Length
	return goldenRatio
}

// IsQuad is a placeholder method for defining the quadrilateral interface
func (s *Square) IsQuad() bool {
	return true
}
