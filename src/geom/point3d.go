package geom

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

// CrossProduct return the cross product between self and other
func (p *Point3D) CrossProduct(q *Point3D) Point3D {
	cprod := Point3D{X: (p.Y * q.Z) - (p.Z * q.Y), Y: (p.X * q.Z) - (p.Z * q.X), Z: (p.X * q.Y) - (p.Y * q.X)}
	return cprod
}

// Normalize returns normalized point of self
func (p *Point3D) Normalize() Point3D {
	magnitude := math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
	return Point3D{X: p.X / magnitude, Y: p.Y / magnitude, Z: p.Z / magnitude}
}

// MidPoint returns midpoint between self and other point
func (p *Point3D) MidPoint(otherPoint Point3D) Point3D {
	return Point3D{X: (p.X + otherPoint.X) / 2, Y: (p.Y + otherPoint.Y) / 2, Z: (p.Z + otherPoint.Z) / 2}
}

// Scale returns a new point scaled by a factor
func (p *Point3D) Scale(scalarFactor float64) Point3D {
	return Point3D{p.X * scalarFactor, p.Y * scalarFactor, p.Z * scalarFactor}
}

// Magnitude returns the size of a Point3D
func (p *Point3D) Magnitude() float64 {
	return math.Sqrt((p.X * p.X) + (p.Y * p.Y) + (p.Z * p.Z))
}

// DistanceToOrigin returns distance of point relative to origin (0,0)
func (p *Point3D) DistanceToOrigin() float64 {
	return p.Magnitude()
}
