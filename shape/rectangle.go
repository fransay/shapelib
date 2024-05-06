package shape

import "math"

type Rectangle struct {
	length float64
	width  float64
}

// AreaByLengthWidth returns area of rectangle by length and width
func (r *Rectangle) AreaByLengthWidth() (area float64) {
	return r.length * r.width
}

// AreaByDiagonal return area of rectangle given the diagonal
func (r *Rectangle) AreaByDiagonal(diagonal float64) (area float64) {
	return 0.5 * diagonal * diagonal
}

// Perimeter returns the perimeter of rectangle
func (r *Rectangle) Perimeter() (perimeter float64) {
	return 2 * (r.length + r.width)
}

// Diagonal returns the length of a rectangle's diagonal
func (r *Rectangle) Diagonal() (diagonal float64) {
	return math.Sqrt((math.Pow(r.width, 2)) + math.Pow(r.width, 2))
}

// NewRectangle returns a new rectangle object
func NewRectangle(length, width float64) *Rectangle {
	return &Rectangle{length: length, width: width}
}
