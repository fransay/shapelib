package point

// Point generic point type
// constraints are point2D, point3D, Point2DH, Point3DH
type Point interface {
	Point2D | Point3D | Point2DH | Point3DH
}
