package shape

import (
	"math"
)

// Rectangle defines a quadrilateral with opposite sides
// equal in length and all angle between the sides are 90 degrees
type Rectangle struct {
	Length float64
	Width  float64
}

// NewRectangle initialises a new Rectangle object
func NewRectangle(length, width float64) *Rectangle {
	return &Rectangle{Length: length, Width: width}
}

// Area returns area of rectangle
func (r Rectangle) Area() (area float64) {
	return r.Length * r.Width
}

// AreaByDiagonal return area of rectangle given the diagonal
func (r Rectangle) AreaByDiagonal(diagonal float64) (area float64) {
	if r.Length != 0 {
		area = r.Length * math.Sqrt((diagonal*diagonal)-(r.Length*r.Length))
	} else {
		area = r.Width * math.Sqrt((diagonal*diagonal)-(r.Width*r.Width))
	}
	return area
}

// Perimeter returns the perimeter of rectangle
func (r Rectangle) Perimeter() (perimeter float64) {
	return 2 * (r.Length + r.Width)
}

// Diagonal returns the Length of a rectangle's diagonal
func (r Rectangle) Diagonal() (diagonal float64) {
	return math.Sqrt((math.Pow(r.Width, 2)) + math.Pow(r.Width, 2))
}

// IsQuad is a placeholder method for defining the quadrilateral interface
func (r Rectangle) IsQuad() bool {
	return true
}

// IsSquare returns a boolean check to see if all length is equal to width
func (r Rectangle) IsSquare() bool {
	return r.Length == r.Width
}
