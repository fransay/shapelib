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

// Point types has no methods
