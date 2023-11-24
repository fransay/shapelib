package coord_sys

import "math"

type Polar struct {
	distance float64 // distance from the origin, in dd
	angle    float64 // referenced angle, i.e bearing
}

// converts polar to cartesian coordinates
func (p *Polar) toCartesian() (c Cart) {
	return c
}

// converts distance to degree minutes and seconds representation
func (p *Polar) distance2dms() (dms []float64) {
	var deg = math.Floor(p.distance) // whole number part
	var mins = math.Floor((p.distance - deg) * 60)
	var sec = (p.distance-math.Floor(p.distance))*60 - math.Floor(p.distance-math.Floor(p.distance))*60
	dms[0] = deg
	dms[1] = mins
	dms[2] = sec
	return dms
}

// converts bearing in angles to radians
func (p *Polar) toRadians() {
	return
}

// distance to another point with a polar representation
func (p *Polar) point2PointDistance(point Polar) float64 {
	return 0
}
