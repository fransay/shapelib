package shape

import "math"

type Rectangle struct {
	Length float64
	Width  float64
}

// NewRectangle initialises a new Rectangle object
func NewRectangle(length, width float64) *Rectangle {
	return &Rectangle{Length: length, Width: width}
}

// AreaByLengthWidth returns area of rectangle by Length and Width
func (r *Rectangle) AreaByLengthWidth() (area float64) {
	return r.Length * r.Width
}

// AreaByDiagonal return area of rectangle given the diagonal
func (r *Rectangle) AreaByDiagonal(diagonal float64) (area float64) {
	return 0.5 * diagonal * diagonal
}

// Perimeter returns the perimeter of rectangle
func (r *Rectangle) Perimeter() (perimeter float64) {
	return 2 * (r.Length + r.Width)
}

// Diagonal returns the Length of a rectangle's diagonal
func (r *Rectangle) Diagonal() (diagonal float64) {
	return math.Sqrt((math.Pow(r.Width, 2)) + math.Pow(r.Width, 2))
}

// NewRectangle returns a new rectangle object
func (r *Rectangle) NewRectangle(length, width float64) *Rectangle {
	return &Rectangle{Length: length, Width: width}
}

// IsQuad is a placeholder method for defining the quadrilateral interface
func (r Rectangle) IsQuad() bool {
	return true
}
