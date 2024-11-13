package proj

import "errors"

// Axis defines an axis in a cartesian plane
type Axis struct {
	Start float64
	End   float64
	Step  float64
	Sign  string // accepts only "+" or "-"
}

// Quadrant defines a quadrant in a cartesian plane
type Quadrant struct {
	X Axis
	Y Axis
}

const (
	FirstQuadrant  string = "first"
	SecondQuadrant string = "second"
	ThirdQuadrant  string = "third"
	FourthQuadrant string = "fourth"
)

// New initialises a fresh Axis object
func (a *Axis) New(start float64, end float64, step float64, sign string) *Axis {
	return &Axis{start, end, step, sign}
}

// GenerateAxisValues returns all values on an axis
func (a *Axis) GenerateAxisValues() []float64 {
	var index int = 0
	var axialValues []float64
	for {
		if a.Sign == "+" {
			axialValues = append(axialValues, a.Start+a.Step)
		} else if a.Sign == "-" {
			axialValues = append(axialValues, -1*a.Start+a.Step)
		} else {
			panic("Invalid Operator")
		}
		if index == int(a.End) {
			break
		}
		index += int(a.Step)
	}
	return axialValues
}

// Type returns type of quadrant
func (q *Quadrant) Type() (string, error) {
	err := errors.New("no quadrant")
	if q.X.Sign == "+" && q.Y.Sign == "+" {
		return FirstQuadrant, nil
	} else if q.X.Sign == "-" && q.Y.Sign == "+" {
		return SecondQuadrant, nil
	} else if q.X.Sign == "-" && q.Y.Sign == "-" {
		return ThirdQuadrant, nil
	} else if q.X.Sign == "+" && q.Y.Sign == "-" {
		return FourthQuadrant, nil
	}
	return "", err
}
