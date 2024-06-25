package types

import (
	"shapelib/utils"
)

// LineString type
type LineString []Point2D // Non-Homogenous 2D point Type

// Length of a linestring
func (l LineString) Length() (totLength float64) {
	for i, j := 0, 1; j < len(l); i, j = i+1, j+1 {
		pointOne := l[i]
		pointTwo := l[j]
		pointOneArr := []float64{pointOne.X, pointOne.Y}
		pointTwoArr := []float64{pointTwo.X, pointTwo.Y}
		dist := utils.Distance(pointOneArr, pointTwoArr)
		totLength = totLength + dist
	}
	return totLength
}

// NumberOfLineSegments returns the total number of elements in an instance of type LineString
func (l *LineString) NumberOfLineSegments() int {
	var indexTracker int
	for index, _ := range *l {
		indexTracker += index
	}
	return indexTracker - 1

}

// Index returns index of linestring in LineString array
func (l *LineString) Index(args Point2D) (index int) {
	var pos int
	for index, value := range *l {
		if value == args {
			pos = index
		}
	}
	return pos

}
