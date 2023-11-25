// generic functions to translate geometric in a vector space

package transformations

import p "shapelib/types/point"

// Translate2D a 2D point type with a shift vector
func Translate2D(point, shiftVector p.Point2D) (translate p.Point2D) {
	translateX := shiftVector.X + point.X
	translateY := shiftVector.Y + point.Y
	translate = p.Point2D{X: translateX, Y: translateY}
	return translate

}

// Translate3D translate a 3D point type with a shift vector
func Translate3D(point, shiftVector p.Point3D) (translate p.Point3D) {
	translateX := shiftVector.X + point.X
	translateY := shiftVector.Y + point.Y
	translateZ := shiftVector.Z + point.Z
	translate = p.Point3D{X: translateX, Y: translateY, Z: translateZ}
	return translate

}

// LineSegmentTranslate translate a line of type line segment
func LineSegmentTranslate() {

}

// LineStringTranslate translate a linestring of type line segment
func LineStringTranslate() {

}

// PolygonTranslate translate a polygon type
func PolygonTranslate() {

}
