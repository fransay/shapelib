package types

// Point2D point type
// characterise by x and y coordination
type Point2D struct {
	X float64
	Y float64
}

// Point3D PointZ 3D point type
// characterise by x,y and z coordination
type Point3D struct {
	X float64
	Y float64
	Z float64
}

// Point2DH Homogeneous 2D point vector
// when W is 1 PointH type is known as an augmented point vector
// alternatively when W is 0, Point2DH type is known as an ideal point
// or a point at infinity
type Point2DH struct {
	X float64
	Y float64
	W float64
}

// Point3DH Homogeneous 2D point vector
// when W is 1 PointH type is known as an augmented point vector
// alternatively when W is 0, Point3DH type is known as an ideal point
// or a point at infinity
type Point3DH struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Point methods

// Size of a point2D is negligible
func (p *Point2D) Size() (size float64) {
	size = 0.0
	return size
}

// Coordinates of point2D type returned in an array
func (p *Point2D) Coordinates() (coordinates []float64) {
	coordinates = []float64{p.X, p.Y}
	return coordinates
}

// Size of point2DH type is negligible
func (p *Point2DH) Size() (size float64) {
	size = 0.0
	return size
}

// Coordinates of point2DH type returned in an array(X,Y,W}
func (p *Point2DH) Coordinates() (coordinates []float64) {
	coordinates = []float64{p.X, p.Y, p.W}
	return coordinates
}

func (p *Point3D) Size() (size float64) {
	size = 0.0
	return size
}

// Coordinates of point3D type returned in an array{X,Y,Z} format
func (p *Point3D) Coordinates() (coordinates []float64) {
	coordinates = []float64{p.X, p.Y, p.Z}
	return coordinates
}

// Size of a 3D homogenous point
// the size of a 3D homogenous point is almost negligible
// 0.0 is returned as negligible
func (p *Point3DH) Size() (size float64) {
	size = 0.0
	return size
}

// Translate2D translates a 2D point
// Translate2D function takes a shift vector
// and return a new point
func (p *Point2D) Translate2D(shiftVector Point2D) (translate Point2D) {
	translateX := p.X + shiftVector.X
	translateY := p.Y + shiftVector.Y
	translate = Point2D{translateX, translateY}
	return translate
}

// Translate3D translates a 3D point
// Translate3D function takes a 3D shift vector
// and return a new point
func (p *Point3D) Translate3D(shiftVector Point3D) (translate Point3D) {
	translateX := p.X + shiftVector.X
	translateY := p.Y + shiftVector.Y
	translateZ := p.Z + shiftVector.Y
	translate = Point3D{translateX, translateY, translateZ}
	return translate
}
