package point

// Point2DH Homogeneous 2D point vector
// when W is 1 PointH type is known as an augmented point vector
// alternatively when W is 0, Point2DH type is known as an ideal point or a point at infinity.
type Point2DH struct {
	X float64
	Y float64
	W float64
}

// Size of point2DH type is negligible
func (p *Point2DH) Size() (size float64) {
	return 0.0
}

// Coordinates of point2DH type returned in an array(X,Y,W)
func (p *Point2DH) Coordinates() (coordinates []float64) {
	coordinates = []float64{p.X, p.Y, p.W}
	return coordinates
}

// Translate2DH returns a new translated Point2DH
func (p *Point2DH) Translate2DH(shiftVector Point2DH) (coordinates Point2DH) {
	translateX := p.X + shiftVector.X
	translateY := p.Y + shiftVector.Y
	translateH := p.W + shiftVector.W
	translatePoint2DH := Point2DH{translateX, translateY, translateH}
	return translatePoint2DH
}

// add affine
// add projective
// add rotate
// add scaling
