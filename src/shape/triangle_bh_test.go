package shape

import (
	"github.com/franela/goblin"
	"testing"
)

func TestTriangleByBaseHeight(t *testing.T) {
	g := goblin.Goblin(t)
	triangleByBaseHeight := NewTriangleByBaseHeight(10, 10)
	g.Describe("TriangleByBaseHeight", func() {
		g.It("Area", func() {
			areaByBaseHeight := triangleByBaseHeight.Area()
			g.Assert(areaByBaseHeight).Equal(50.0)
		})

		g.It("Perimeter", func() {
			perimeter := triangleByBaseHeight.Perimeter(10, 10, 10)
			g.Assert(perimeter).Equal(30.0)
		})
	})
}
