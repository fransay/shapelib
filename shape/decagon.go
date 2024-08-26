package shape

import (
	"math"
	"shapelib/utils"
)

type Decagon struct {
	sideLength float64
}

// Area returns area of a decagon
func (d *Decagon) Area() (area float64) {
	area = (5 / 2) * math.Pow(d.sideLength, 2) * utils.Cot(math.Pi/10)
	return area
}

// Perimeter returns the perimeter of a decagon
func (d *Decagon) Perimeter() (perimeter float64) {
	return d.sideLength * 10
}

// Circumradius returns the radius of the circumscribed circle
func (d *Decagon) Circumradius() float64 {
	return d.sideLength / 2 * (math.Sin(18.0))
}

// Inradius returns the radius of a regular decagon
func (d *Decagon) Inradius() float64 {
	return (d.sideLength * math.Sqrt(5+2*math.Sqrt(5))) / 2
}
