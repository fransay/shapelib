package shape

import (
	"github.com/franela/goblin"
	"testing"
)

func TestDecagon(t *testing.T) {
	g := goblin.Goblin(t)
	decagon, _ := NewDecagon(20)
	g.Describe("Test Decagon", func() {
		g.It("Area", func() {
			area := decagon.Area()
			g.Assert(area).Equal(3077.68)
		})
		g.It("Perimeter", func() {
			perimeter := decagon.Perimeter()
			g.Assert(perimeter).Equal(200)
		})
	})
}
