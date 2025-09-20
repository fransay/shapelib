package shape

import (
	"github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
	"testing"
)

func TestSquare(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Test Square methods", func() {
		square := NewSquare(3.0)
		g.It("Area", func() {
			area := square.Area()
			g.Assert(area).Equal(9.0)
		})

		g.It("Perimeter", func() {
			perimeter := square.Perimeter()
			g.Assert(perimeter).Equal(12.0)
		})

		g.It("Diagonal", func() {
			diagonal := square.Diagonal()
			isClose := utils.IsClose(diagonal, 4.242640687119286, 0.001)
			g.Assert(isClose).IsTrue()
		})

		g.It("CircumRadius", func() {
			circumRadius := square.CircumRadius()
			isClose := utils.IsClose(circumRadius, 2.121320343559643, 0.001)
			g.Assert(isClose).IsTrue()
		})

		g.It("InRadius", func() {
			inRadius := square.InRadius()
			isClose := utils.IsClose(inRadius, 1.5, 0.00001)
			g.Assert(isClose).IsTrue()
		})

		g.It("Apothem", func() {
			apothem := square.Apothem()
			g.Assert(apothem).Equal(1.500000)
		})

		g.It("Golden Ratio", func() {
			goldenRatio := square.GoldenRation()
			isClose := utils.IsClose(goldenRatio, 1.4142135623730951, 0.00001)
			g.Assert(isClose).IsTrue()
		})
	})
}
