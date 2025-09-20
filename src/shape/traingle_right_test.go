package shape

import (
	"fmt"
	"testing"

	"github.com/franela/goblin"
	"github.com/fransay/shapelib/src/geom"
)

func TestTriangleRight(t *testing.T) {
	g := goblin.Goblin(t)
	triangleRight := NewTriangleRight(
		geom.Point2D{X: 3, Y: 5},
		geom.Point2D{X: 7, Y: 9},
		geom.Point2D{X: 16, Y: 12},
	)
	g.Describe("triangle functions", func() {
		g.It("test area ", func() {
			area := triangleRight.Area()
			fmt.Println("area =", area)
			g.Assert(area).Equal(36.0)

		})
	})
}
