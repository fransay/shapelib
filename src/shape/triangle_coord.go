package shape

// TriangleByCoord defines a triangle given all coordinates
type TriangleByCoord [3][2]float64

// Center returns a point at the center of the triangle
func (tc *TriangleByCoord) Center() (pt [2]float64) {
	var x float64
	var y float64
	for _, pnt := range tc {
		x += pnt[0]
		y += pnt[1]
	}
	pt = [2]float64{x / 3, y / 3}
	return pt
}

// Area returns the area of triangle given the coordinates
func (tc *TriangleByCoord) Area() (area float64) {
	area = (tc[0][0]*(tc[1][1]-tc[2][1]) + tc[1][0]*(tc[2][1]-tc[1][1]) + tc[2][0]*(tc[0][1]-tc[1][1])) / 2
	return area
}

// Contains checks if a point is contained in a triangle or not.
func (tc *TriangleByCoord) Contains()    {}
func (tc *TriangleByCoord) ThirdPoint()  {}
func (tc *TriangleByCoord) LinearRings() {}
