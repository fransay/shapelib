package functs

import (
	"math"
	"shapelib/types"
)

// EDistance returns euclidean distance between two points/coordinates
func EDistance(pointOne, pointTwo types.Point2D) (dist float64) {
	deltaX := pointTwo.X - pointOne.X
	deltaY := pointOne.Y - pointOne.Y
	dist = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return dist
}
