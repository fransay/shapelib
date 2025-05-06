package shape

import (
	"github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
	"testing"
)

func TestRectangle(t *testing.T) {
	g := goblin.Goblin(t)
	var rectangle = Rectangle{Length: 15, Width: 20}
	g.Describe("Test 25Rectangle", func() {
		g.It("Area By Length Width", func() {
			areaByLengthWidth := rectangle.AreaByLengthWidth()
			g.Assert(areaByLengthWidth).Equal(300.0)
		})

		g.It("Area By Diagonal", func() {
			areaByDiagonal := rectangle.AreaByDiagonal(25)
			g.Assert(areaByDiagonal).Equal(300.0)
		})

		g.It("Perimeter", func() {
			perimeter := rectangle.Perimeter()
			g.Assert(perimeter).Equal(70.0)
		})

		g.It("Diagonal", func() {
			diagonal := rectangle.Diagonal()
			isClose := utils.IsClose(diagonal, 28.284271, 0.001)
			g.Assert(isClose).IsTrue()
		})
	})
}
