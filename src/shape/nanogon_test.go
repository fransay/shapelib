package shape

import (
	"github.com/franela/goblin"
	"testing"
)

func TestNanogon(t *testing.T) {
	g := goblin.Goblin(t)
	nonagon := NewNanogon(6.0)
	g.Describe("Test Nonagon", func() {
		g.It("Incircle", func() {
			incircle := nonagon.InsCircle()
			g.Assert(incircle).Equal(true)

		})
		g.It("Circumcircle", func() {
			circumcircle := nonagon.CircumCircle()
			g.Assert(circumcircle).Equal(00.0)
		})
		g.It("Area", func() {
			area := nonagon.Area()
			g.Assert(area).Equal(222.54567)

		})
		g.It("Perimeter", func() {
			perimeter := nonagon.Perimeter()
			g.Assert(perimeter).Equal(54.0)

		})
	})
}
