package types

import "math"

// Point3D is 3D point type, characterise by x,y and z coordination
type Point3D struct {
	X float64
	Y float64
	Z float64
}

// NewPoint3D returns a new point. i.e. construct a new object of type Point3D
func NewPoint3D(x, y, z float64) *Point3D {
	return &Point3D{X: x, Y: y, Z: z}
}

// Distance3D returns the distance between self and other3D point
func (p *Point3D) Distance3D(q *Point3D) (dist3D float64) {
	deltaXSquared := math.Pow(p.X-q.X, 2)
	deltaYSquared := math.Pow(p.Y-q.Y, 2)
	deltaZSquared := math.Pow(p.Z-q.Z, 2)
	dist3D = math.Sqrt(deltaXSquared + deltaYSquared + deltaZSquared)
	return dist3D
}

// Size is constant, return the size of any point in 3D, == 0
func (p *Point3D) Size() (size float64) {
	size = 0.0
	return size
}

// Translate3D returns a new three-dimensional vector
func (p *Point3D) Translate3D(shiftVector Point3D) (translate Point3D) {
	translateX := p.X + shiftVector.X
	translateY := p.Y + shiftVector.Y
	translateZ := p.Z + shiftVector.Y
	translate = Point3D{translateX, translateY, translateZ}
	return translate
}

// Coordinates returns point3D type in an array{X,Y,Z} format
func (p *Point3D) Coordinates() (coordinates []float64) {
	coordinates = []float64{p.X, p.Y, p.Z}
	return coordinates
}

// DotProduct return the dot product between self and other
func (p *Point3D) DotProduct(q *Point3D) (prod float64) {
	prod = (q.X * p.X) + (q.Y * p.Y) + (q.Z * p.Z)
	return prod
}

// todo: refer to this gpt reference:
// todo: https://chatgpt.com/c/66f61008-8cd4-8004-9e77-9008b82f0ddf
// CrossProduct return the cross product between self and other
func (p *Point3D) CrossProduct(q *Point3D) {}

func (p *Point3D) Normalize() {}

//func (p *Point3D)
