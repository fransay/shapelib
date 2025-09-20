package shape

import (
	"github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
	"testing"
)

func TestHeptagon(t *testing.T) {
	g := goblin.Goblin(t)
	heptagon, _ := NewHeptagon(30)
	g.Describe("Heptagon", func() {
		g.It("Area", func() {
			isClose := utils.IsClose(heptagon.Area(), 3270.52, 0.001)
			g.Assert(isClose).Equal(true)
		})
		g.It("Perimeter", func() {
			perimeter := heptagon.Perimeter()
			g.Assert(perimeter).Equal(210.0)
		})
		g.It("Apothem", func() {
			apothem := heptagon.Apothem()
			isClose := utils.IsClose(apothem, 31.14782094858505, 0.001)
			g.Assert(isClose).Equal(true)
		})
	})
}
