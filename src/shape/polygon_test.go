package shape

import (
	"github.com/franela/goblin"
	"testing"

	"github.com/fransay/shapelib/src/geom"
)

func TestPolygon(t *testing.T) {
	g := goblin.Goblin(t)
	points := []geom.Point2D{
		{10, 50},
		{2, 15},
		{15, 10}}
	polygon := NewPolygon(points)
	g.Describe("Polygon", func() {
		g.It("Number of nodes", func() {
			numberOfNodes := polygon.NumberOfNodes()
			g.Assert(numberOfNodes).Equal(3)
		})
		g.It("Area", func() {
			area := polygon.Area()
			g.Assert(area).Equal(247.5)
		})
		g.It("Line segment", func() {})
		g.It("Centroid", func() {})
		g.It("Number of Line Segments", func() {})
		g.It("Rotate", func() {})
		g.It("Scale", func() {})
		g.It("IsEqual", func() {})
		g.It("Contains", func() {})
	})
}
