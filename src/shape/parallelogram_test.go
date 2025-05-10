package shape

import (
	"fmt"
	"github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
	"testing"
)

func TestParallelogram(t *testing.T) {
	g := goblin.Goblin(t)
	parallelogram := NewParallelogram(20, 15, 25)
	g.Describe("Parallelogram", func() {
		g.It("Area By Base and Height", func() {
			areaByBaseHeight := parallelogram.AreaByBaseHeight()
			isClose := utils.IsClose(areaByBaseHeight, 300.0, 0.001)
			g.Assert(isClose).IsTrue()
		})

		g.It("Perimeter", func() {
			perimeter := parallelogram.Perimeter()
			fmt.Println(perimeter)
			isClose := utils.IsClose(perimeter, 90.0, 0.001)
			g.Assert(isClose).IsTrue()
		})
	})

}
