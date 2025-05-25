package shape

import (
	"github.com/franela/goblin"
	"github.com/fransay/shapelib/src/geom"
	"testing"
)

func TestTriangleByCoordinates(t *testing.T) {
	g := goblin.Goblin(t)
	var triangleByCoordinates = NewTriangleByCoordinates(
		geom.Point2D{X: 50, Y: 160},
		geom.Point2D{X: 70, Y: 120},
		geom.Point2D{X: 45, Y: 500})

	g.Describe("Test TriangleCoordinates", func() {
		g.It("Center", func() {
			center := triangleByCoordinates.Center()
			g.Assert(center).Equal(geom.Point2D{
				X: 55,
				Y: 260})
		})
		g.It("Area", func() {
			area := triangleByCoordinates.Area()
			g.Assert(area).Equal(3300.00)
		})
		g.It("Perimeter", func() {
			perimeter := triangleByCoordinates.Perimeter(
				10.0,
				20.0,
				30.0)
			g.Assert(perimeter).Equal(60.0)
		})
	})

}
