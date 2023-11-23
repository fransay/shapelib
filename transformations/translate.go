// generic functions to translate geometric in a vector space

package transformations

import (
	"shapelib/types/point"
)

// Translate2D translate a 2D point type
// translate 2D point
func Translate2D(point, shiftVector point.Point2D) (translate point.Point2D) {
	translateX := shiftVector.X + point.X
	translateY := shiftVector.Y + point.Y
	translate = point.Point2D{X: translateX, Y: translateY}
	return translate

}

// Translate3D translate a 3D point type
func Translate3D(point, shiftVector point.Point3D) (translate point.Point3D) {
	translateX := shiftVector.X + point.X
	translateY := shiftVector.Y + point.Y
	translateZ := shiftVector.Z + point.Z
	translate = point.Point3D{X: translateX, Y: translateY, Z: translateZ}
	return translate

}

// LineSegmentTranslate translate a line of type line segment
func LineSegmentTranslate() {

}

// LineStringTranslate translate a linestring of type line segment
func LineStringTranslate() {

}

// PolygonTranslate translate a polygon type
// translate a polygon
func PolygonTranslate() {

}
