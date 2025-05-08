package geog

import "math"

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

// Distance returns the Euclidean distance between
// two positions in UTM.
func (utm *UTM) Distance(utm1 UTM) float64 {
	deltaEastings := utm1.Easting - utm.Easting
	deltaNorthings := utm1.Northing - utm.Northing
	return math.Sqrt(math.Pow(deltaEastings, 2.0) + math.Pow(deltaNorthings, 2.0))
}

// Geodesic returns the geodesic distance using the Vincenty's distance formulae.
func (utm *UTM) Geodesic(utm1 UTM) {}

// ToLatLong returns latlong representation of UTM
func (utm *UTM) ToLatLong(points []UTM) {}

// Area defines the area of a set of points fr t
func (utm *UTM) Area(points []UTM) {}
