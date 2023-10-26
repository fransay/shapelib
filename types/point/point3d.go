package point

// Point3D PointZ 3D point type, characterise by x,y and z coordination
type Point3D struct {
	X float64
	Y float64
	Z float64
}

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

//TODO: Add methods [consider extra additional functions]
