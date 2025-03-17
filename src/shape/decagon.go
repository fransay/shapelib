package shape

import (
	"errors"
	"math"

	"github.com/fransay/shapelib/internal/utils"
)

// Decagon type
type Decagon struct {
	Side float64
}

// NewDecagon initialises a new decagon object
func NewDecagon(side float64) (*Decagon, error) {
	if side < 0.0 {
		return nil, errors.New("can't create new decagon object")
	}
	return &Decagon{
		Side: side,
	}, nil
}

// Area returns area of a decagon
func (d *Decagon) Area() (area float64) {
	area = (5 / 2) * math.Pow(d.Side, 2) * utils.Cot(math.Pi/10)
	return area
}

// Perimeter returns the perimeter of a decagon
func (d *Decagon) Perimeter() (perimeter float64) {
	return d.Side * 10
}

// Circumradius returns the radius of the circumscribed circle
func (d *Decagon) Circumradius() float64 {
	return d.Side / 2 * (math.Sin(18.0))
}

// Inradius returns the radius of a regular decagon
func (d *Decagon) Inradius() float64 {
	return (d.Side * math.Sqrt(5+2*math.Sqrt(5))) / 2
}
