package tellus

// UTM
type UTM struct {
	Easting  float64
	Northing float64
}

func NewUTM(easting float64, northing float64) *UTM {
	return &UTM{
		Easting:  easting,
		Northing: northing,
	}
}

func (utm *UTM) Distance(utm1 UTM) float64 {
	return 0.0
}
