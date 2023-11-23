package point

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

// Size of a 3D homogenous point
// the size of a 3D homogenous point is almost negligible
func (p *Point3DH) Size() (size float64) {
	return 0.0
}

//TODO: Add methods [consider extra additional functions]
