package tellus

import "math"

// LatLong type
type LatLong struct {
	Latitude  float64 // in degrees
	Longitude float64 // in degrees
}

// Initialise new LatLong object
func NewLatLong(lat, long float64) LatLong {
	return LatLong{Latitude: lat, Longitude: long}
}

// ToDMS return degree minute seconds represent f LatLong
// which is in degree decimals (DD).
func (l *LatLong) ToDMS() ([]float64, []float64) {
	lat := l.Latitude
	long := l.Longitude
	latDMS := conv(lat)
	longDMS := conv(long)
	return latDMS, longDMS

}

// conv returns a conversion of f into whole, mins, secs
func conv(f float64) (fr []float64) {
	fInt := int(f)
	wholeDec := f - float64(fInt)
	mInt := int(wholeDec * 60.0)
	minuteDec := (wholeDec * 60.0) - float64(mInt)
	sInt := minuteDec * 60.0
	fr = []float64{float64(fInt), float64(mInt), float64(sInt)}
	return fr
}

// Distance return the distance between two latlongs based on haversine
func (l *LatLong) Distance(other LatLong) float64 {
	var radiusOfEarth = 6371.0 // in km
	var deltaLat = other.Latitude - l.Latitude
	var deltaLong = other.Longitude - l.Longitude
	var a = (math.Sin(deltaLat/2) * math.Sin(deltaLat/2)) * math.Cos(l.Latitude) * math.Cos(other.Latitude) * math.Sin(deltaLong/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radiusOfEarth * c
}

func (l *LatLong) Geodesic(other LatLong) {}

