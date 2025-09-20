package shape

import (
	"testing"

	"github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
)

func TestTriangle(t *testing.T) {
	g := goblin.Goblin(t)
	var triangle = NewTriangle(
		50.0,
		60.0,
		30.0)
	g.Describe("Test Triangle", func() {
		g.It("Area by sides", func() {
			areaBySides := triangle.Area()
			isClose := utils.IsClose(areaBySides, 748.3314773547883, 0.0001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Area by perimeter", func() {
			perimeter := triangle.Perimeter()
			isClose := utils.IsClose(perimeter, 140.0, 0.0001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Area by Type", func() {
			scaleneTriangle := triangle.Type()
			g.Assert(scaleneTriangle).Equal("Scalene")
			equilateralTriangle := NewTriangle(10.0, 10.0, 10.0).Type()
			g.Assert(equilateralTriangle).Equal("Equilateral")
			isoscelesTriangle := NewTriangle(10.0, 10.0, 5.0).Type()
			g.Assert(isoscelesTriangle).Equal("Isosceles")
		})
	})

}
