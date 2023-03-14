// generic functions to translate geometric in a vector space

package transformations

import "shapelib/types"

// Translate2D translate a 2D point type
// translate 2D point
func Translate2D(point, shiftVector types.Point2D) (translate types.Point2D) {
	translateX := shiftVector.X + point.X
	translateY := shiftVector.Y + point.Y
	translate = types.Point2D{X: translateX, Y: translateY}
	return translate

}

// Translate3D translate a 3D point type
func Translate3D(point, shiftVector types.Point3D) (translate types.Point3D) {
	translateX := shiftVector.X + point.X
	translateY := shiftVector.Y + point.Y
	translateZ := shiftVector.Z + point.Z
	translate = types.Point3D{X: translateX, Y: translateY, Z: translateZ}
	return translate

}

// LineSegmentTranslate translate a line segment type
// translate a lineSegment
func LineSegmentTranslate() {

}

// LineStringTranslate translate a linestring type
// translate a linestring
func LineStringTranslate() {

}

// PolygonTranslate translate a polygon type
// translate a polygon
func PolygonTranslate() {

}
