package shape

import "github.com/fransay/shapelib/src/geom"

type TriangleRight struct {
	PointOne   geom.Point2D
	PointTwo   geom.Point2D
	PointThree geom.Point2D
}

func NewTriangleRight(pointOne, pointTwo, pointThree geom.Point2D) *TriangleRight {
	return &TriangleRight{
		PointOne:   pointOne,
		PointTwo:   pointTwo,
		PointThree: pointThree,
	}
}

func (tri *TriangleRight) Area() float64 {
	distanceSideOne := distance(tri.PointOne, tri.PointTwo)
	distanceSideTwo := distance(tri.PointTwo, tri.PointThree)
	distanceSideThree := distance(tri.PointOne, tri.PointThree)
	maxDist := MaxDistance(distanceSideOne, distanceSideTwo, distanceSideThree)
	a, b := NonMaxDistance(distanceSideOne, distanceSideTwo, distanceSideThree, maxDist)
	return (a * b) / 2
}

func MaxDistance(a, b, c float64) float64 {
	maxDist := a
	if b > a {
		maxDist = b
	}
	if c > a {
		maxDist = c
	}
	return maxDist
}

// NonMaxDistance return a non-max numbers
func NonMaxDistance(a, b, c, max float64) (float64, float64) {
	if a == max {
		return b, c
	} else if b == max {
		return a, c
	}
	return a, b
}

func (tri *TriangleRight) Perimeter() float64 {
	distanceSideOne := distance(tri.PointOne, tri.PointTwo)
	distanceSideTwo := distance(tri.PointTwo, tri.PointThree)
	distanceSideThree := distance(tri.PointOne, tri.PointThree)
	return distanceSideOne + distanceSideTwo + distanceSideThree
}
