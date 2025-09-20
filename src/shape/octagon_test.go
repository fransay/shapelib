package shape

import (
	"github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
	"testing"
)

func TestOctagon(t *testing.T) {
	g := goblin.Goblin(t)
	octagon := NewOctagon(50)
	g.Describe("Test Area of Octagon", func() {
		g.It("Area", func() {
			area := octagon.Area()
			isClose := utils.IsClose(area, 0.0, 0.001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Perimeter", func() {
			perimeter := octagon.Perimeter()
			isClose := utils.IsClose(perimeter, 0.0, 0.001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Apothem", func() {
			apothem := octagon.Apothem()
			isClose := utils.IsClose(apothem, 0.0, 0.001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Internal Diagonal", func() {
			internalDiagonal := octagon.InteriorDiagonal()
			isClose := utils.IsClose(internalDiagonal, 0.0, 0.001)
			g.Assert(isClose).Equal(true)
		})

	})
}
