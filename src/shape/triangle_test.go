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
		g.It("Test Triangle By Sides", func() {
			areaBySides := triangle.Area()
			isClose := utils.IsClose(areaBySides, 748.3314773547883, 0.0001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Test Triangle Perimeter", func() {
			perimeter := triangle.Perimeter()
			isClose := utils.IsClose(perimeter, 140.0, 0.0001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Test Triangle Type", func() {
			scaleneTriangle := triangle.Type()
			g.Assert(scaleneTriangle).Equal("Scalene")
			equilateralTriangle := NewTriangle(10.0, 10.0, 10.0).Type()
			g.Assert(equilateralTriangle).Equal("Equilateral")
			isoscelesTriangle := NewTriangle(10.0, 10.0, 5.0).Type()
			g.Assert(isoscelesTriangle).Equal("Isosceles")
		})
		g.It("Test SemiPerimeter", func() {
			triangleSemiPerimeter := triangle.SemiPerimeter()
			isClose := utils.IsClose(triangleSemiPerimeter, 70.0, 0.0001)
			g.Assert(isClose).Equal(true)
		})

		g.It("Test InRadius", func() {
			triangleInRadius := triangle.InRadius()
			isClose := utils.IsClose(triangleInRadius, 10.690449676496975, 0.0001)
			g.Assert(isClose).Equal(true)
		})
		g.It("Test CircumRadius", func() {
			triangleCircumRadius := triangle.CircumRadius()
			isClose := utils.IsClose(triangleCircumRadius, 30.066889715147745, 0.0001)
			g.Assert(isClose).Equal(true)
		})
	})

}
