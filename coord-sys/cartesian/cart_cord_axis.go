package cartesian

// Axis defines an axis in a cartesian plane
type Axis struct {
	Start float64
	End   float64
	Step  float64
}

// New initialises a fresh Axis object
func (a *Axis) New(start float64, end float64, step float64) *Axis {
	return &Axis{start, end, step}
}
