package shape

import (
	"fmt"
	"github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
	"testing"

	"github.com/fransay/shapelib/src/geom"
)

func TestCircle(t *testing.T) {
	g := goblin.Goblin(t)
	var circle, err = NewCircle(60, geom.Point2D{X: 20, Y: 45})
	g.Assert(err).IsNil(true)
	g.Describe("Test circle methods", func() {
		g.It("Area", func() {
			area := circle.Area()
			fmt.Println(area)
			isClose := utils.IsClose(area, 11309.733552923255, 0.0001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Circumference", func() {
			circle := circle.Circumference()
			isClose := utils.IsClose(circle, 376.99112, 0.0001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Arc length", func() {
			circle := circle.ArcLength(125.0, 60)
			isClose := utils.IsClose(circle, 130.8997, 0.0001)
			g.Assert(isClose).Equal(true)
		})

		// todo: sectorArea
		// todo: chordLength
		// todo: containsPoint
		// todo: circleFromPoints
		// todo: distance
		// todo: gradients
		// todo: lineIntersectGivenEquation
		// todo: asLineString
	})
}
