package shape

import (
	"github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
	"testing"

	"github.com/fransay/shapelib/src/geom"
)

func TestTriangle(t *testing.T) {
	g := goblin.Goblin(t)
	var triangle = NewTriangle(
		50.0,
		60.0,
		30.0,
		30.0,
		10.0)
	g.Describe("Area by sides", func() {
		areaBySides := triangle.AreaBySides()
		isClose := utils.IsClose(areaBySides, 748.3314773547883, 0.0001)
		g.Assert(isClose).Equal(true)
	})

	g.Describe("Area by base and height", func() {
		areaByBaseHeight := triangle.AreaByBaseHeight()
		isClose := utils.IsClose(areaByBaseHeight, 150.0, 0.0001)
		g.Assert(isClose).Equal(true)
	})
	g.Describe("Area by perimeter", func() {
		perimeter := triangle.Perimeter()
		isClose := utils.IsClose(perimeter, 140.0, 0.0001)
		g.Assert(isClose).Equal(true)
	})

	g.Describe("Area by Type", func() {
		ttype := triangle.Type()
		g.Assert(ttype).Equal("Scalene")
	})
}

func TestTriangleCoordinates(t *testing.T) {
	g := goblin.Goblin(t)
	var triangleByCoordinates = NewTriangleByCoordinates(geom.Point2D{X: 50, Y: 160}, geom.Point2D{X: 70, Y: 120}, geom.Point2D{X: 45, Y: 500})

	g.Describe("Center", func() {
		center := triangleByCoordinates.Center()
		g.Assert(center).Equal(geom.Point2D{X: 55, Y: 260})
	})

	g.Describe("Area", func() {
		area := triangleByCoordinates.Area()
		g.Assert(area).Equal(3300.00)
	})
}
