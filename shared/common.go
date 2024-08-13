package shared

import "shapelib/types"

func AreaByCoordinates(cords ...types.Point2D) float64 {
	var forPass float64
	var backPass float64
	var numberOfCords = len(cords) - 1
	for i, j := 0, 1; j <= numberOfCords; i, j = i+1, j+1 {
		forPass += cords[i].Y * cords[j].X
		backPass += cords[j].Y * cords[i].X
	}
	var area = -1 * (forPass - backPass) / 2 // results of area are negated due to orientation of types.Point2D
	return area
}
