package shape

import (
	"fmt"
	. "github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
	"testing"
)

func TestPentagon(t *testing.T) {
	g := Goblin(t)
	var pent = NewPentagon(20.0)
	g.Describe("Test Pentagon Methods", func() {
		g.It("Area", func() {
			area := pent.Area()
			isClose := utils.IsClose(area, 688.19096, 0.0001)
			g.Assert(isClose).Equal(true)
		})
		g.It("Perimeter", func() {
			perimeter := pent.Perimeter()
			isClose := utils.IsClose(perimeter, 100.00, 0.0001)
			g.Assert(isClose).Equal(true)
		})
		g.It("Diagonal", func() {
			diagonal := pent.Diagonal()
			isClose := utils.IsClose(diagonal, 32.36068, 0.0001)
			g.Assert(isClose).Equal(true)
		})
		g.It("Circumcircle", func() {
			circle := pent.Circumcircle()
			isClose := utils.IsClose(circle, 17.013023469465875, 0.0001)
			g.Assert(isClose).Equal(true)
		})
		g.It("Incircle", func() {
			circle := pent.Incircle()
			fmt.Print(circle)
			isClose := utils.IsClose(circle, 13.763677654669326, 0.0001)
			g.Assert(isClose).Equal(true)
		})
	})
}
