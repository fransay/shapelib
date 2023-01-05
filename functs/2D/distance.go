package functs

import (
	"math"
	"shapelib/types"
)

// Distance returns distance
func Distance(pointOne, pointTwo types.Point2D) (dist float64) {
	deltaX := pointTwo.X - pointOne.X
	deltaY := pointOne.Y - pointOne.Y
	dist = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return dist
}
