package shape

import (
	"github.com/franela/goblin"
	"testing"
)

func TestQuadrilateral(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Quadrilateral", func() {
		var rect Quadrilateral = Rectangle{Length: 20.0, Width: 50.0}
		var square Quadrilateral = Square{Length: 20.0}

		g.It("Is Quadrilateral", func() {
			g.Assert(!rect.IsQuad()).IsFalse()
		})

		g.It("Is Quadrilateral", func() {
			g.Assert(!square.IsQuad()).IsFalse()
		})
	})
}
