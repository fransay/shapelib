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
		g.It("Centroid", func() {
			var centroid, err = polygon.Centroid()
			g.Assert(err).Equal(nil)
			g.Assert(centroid).Equal(geom.Point2D{X: 9, Y: 25})
		})

		g.It("Segment Distances", func() {
			segDist := polygon.SegmentDistances()
			g.Assert(segDist).Equal([]float64{}) // todos: manually compute distances
		})

		g.It("Perimeter", func() {
			perimeter := polygon.Perimeter()
			g.Assert(perimeter).Equal(0)
		})

		g.It("Shortest Line Segment", func() {
			shortestLineSegment := polygon.ShortestLineSegment()
			g.Assert(shortestLineSegment).Equal(0.0)
		})

		g.It("IsClosed", func() {
			isClosed := polygon.IsClosed()
			g.Assert(isClosed).Equal(true)
		})
		g.It("IsOpen", func() {
			isOpened := polygon.IsOpened()
			g.Assert(isOpened).Equal(true)
		})
		g.It("Number Of LineSegments", func() {
			numberOfLineSegments := polygon.NumberOfLineSegments()
			g.Assert(numberOfLineSegments).Equal(3)
		})
		g.It("Number of Nodes", func() {
			numberOfNodes := polygon.NumberOfNodes()
			g.Assert(numberOfNodes).Equal(3)
		})
		g.It("Rotate by 60 degrees", func() {
			rotate := polygon.Rotate(60)
			g.Assert(rotate).Equal(geom.Point2D{X: 0, Y: 0})
		})
		g.It("Scale", func() {})
		g.It("IsEqual", func() {})
	})
}
