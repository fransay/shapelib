package props

import (
	"fmt"
	"github.com/franela/goblin"
	"github.com/fransay/shapelib/internal/utils"
	"math"
	"testing"

	"github.com/fransay/shapelib/src/geom"
)

func TestArea(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Area", func() {
		g.It("Area of TriangleBH", func() {
			areaOfTriangleBH := AreaOfTriangleGivenBH(15, 20)
			g.Assert(areaOfTriangleBH).Equal(150.0)
		})

		g.It("Area of Triangle given their sides", func() {
			areaOfTriangleGivenSide := AreaOfTriangleGivenSide(12.0, 8.0, 10.0)
			isClose := utils.IsClose(areaOfTriangleGivenSide, 56.57, 0.001)
			g.Assert(isClose).IsTrue()
		})

		g.It("Area of Triangle given their angle", func() {
			areaOfTriangleGivenAngle := AreaOfTriangleGivenAngle(10.0, 5.0, 120.0)
			isClose := utils.IsClose(areaOfTriangleGivenAngle, 14.515280, 0.001)
			g.Assert(isClose).IsTrue()
		})

		g.It("Area of Square", func() {
			areaOfSquare := AreaOfSquare(5.0)
			g.Assert(areaOfSquare).Equal(25.0)
		})

		g.It("Area of a Rectangle", func() {
			areaOfRectangle := AreaOfRectangle(10.0, 5.0)
			g.Assert(areaOfRectangle).Equal(50.0)
		})

		g.It("Area of Parallelogram", func() {
			areaOfParallelogram := AreaOfParallelogram(10.0, 5.0)
			g.Assert(areaOfParallelogram).Equal(50.0)
		})

		g.It("Area of Kite", func() {
			areaOfKite := AreaOfKite(10.0, 5.0)
			g.Assert(areaOfKite).Equal(25.0)
		})

		g.It("Area of Rhombus", func() {
			areaOfRhombus := AreaOfRhombus(10.0, 5.0)
			g.Assert(areaOfRhombus).Equal(25.0)
		})

		g.It("Area of Trapezoid", func() {
			areaOfTrapezoid := AreaOfTrapezoid(10.0, 5.0, 5.0)
			g.Assert(areaOfTrapezoid).Equal(37.5)
		})

		g.It("Area of Pentagon", func() {
			areaOfPentagon := AreaOfPentagon(5.0)
			isClose := utils.IsClose(areaOfPentagon, 43.01194, 0.001)
			g.Assert(isClose).IsTrue()
		})

		g.It("Area of Hexagon", func() {
			areaOfHexagon := AreaOfHexagon(5.0)
			isClose := utils.IsClose(areaOfHexagon, 259.80762, 0.001)
			g.Assert(isClose).IsTrue()
		})
		g.It("Area of Heptagon", func() {
			areaOfHeptagon := AreaOfHeptagon(5.0)
			fmt.Println("area of heptagon", areaOfHeptagon)
			isClose := utils.IsClose(areaOfHeptagon, 90.84781, 0.001)
			g.Assert(isClose).IsTrue()
		})
		g.It("Area of Octagon", func() {
			areaOfOctagon := AreaOfOctagon(5.0)
			isClose := utils.IsClose(areaOfOctagon, 120.71, 0.001)
			g.Assert(isClose).IsTrue()
		})
		g.It("Area of Nonagon", func() {
			areaOfNonagon := AreaOfNonagon(5.0)
			isClose := utils.IsClose(areaOfNonagon, 154.5456, 0.001)
			g.Assert(isClose).IsTrue()
		})
		g.It("Area of Decagon", func() {
			areaOfDecagon := AreaOfDecagon(5.0)
			isClose := utils.IsClose(areaOfDecagon, 192.35522, 0.001)
			g.Assert(isClose).IsTrue()
		})
		g.It("Area of Dodecagon", func() {})
		g.It("Area of Circle", func() {})
		g.It("Area of Semi-Circle", func() {})
		g.It("Area of Quad-Circle", func() {})
		g.It("Area of Oval", func() {})
		g.It("Area of Ellipse", func() {})
		g.It("Area By Coordinates", func() {})

	})
}

// Test area of a regular heptagon
func TestAreaOfRHeptagon(t *testing.T) {
	res := AreaOfHeptagon(5.00)
	exp := math.Trunc(90.85)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a regular dodecagon
func TestAreaOfRDodecagon(t *testing.T) {
	res := AreaOfDodecagon(5)
	exp := math.Trunc(64.9519052838)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a circle
func TestAreaOfCircle(t *testing.T) {
	res := AreaOfCircle(5.00)
	exp := math.Trunc(75.00)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a semicircle
func TestAreaOfSemiCircle(t *testing.T) {
	res := AreaOfSemiCircle(5.00)
	exp := math.Trunc(37.50)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of a quad circle
func TestAreaOfQuadCircle(t *testing.T) {
	res := AreaOfQuadCircle(5.00)
	exp := math.Trunc(18.75)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}

// Test area of an oval
func TestAreaOfOval(t *testing.T) {
	res := AreaOfOval(10.00, 8.0)
	exp := math.Trunc(251.32)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)

	}
}

// Test area of an ellipse
func TestAreaOfEllipse(t *testing.T) {
	res := AreaOfEllipse(10.00, 8.0)
	exp := math.Trunc(251.32)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)

	}
}

// Test area given coordinates
func TestAreaCoordinates(t *testing.T) {
	var res = AreaGivenCoordinates(geom.Point2D{X: 1, Y: 4}, geom.Point2D{X: 3, Y: 6}, geom.Point2D{X: 8, Y: 5})
	var exp = math.Trunc(-19.500)
	if math.Trunc(res) != exp {
		t.Errorf("Expected %f Got %f", exp, res)
	}
}
