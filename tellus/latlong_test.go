package tellus

import (
	"shapelib/utils"
	"testing"
)

func TestLatLong(t *testing.T) {
	latlong1 := LatLong{Latitude: 80.00, Longitude: 45.00}
	latlong2 := LatLong{Latitude: 8.00, Longitude: 4.00}

	// Distance
	observedDistance := latlong1.Distance(latlong2)
	expectedDistance := 233.23 // subjected to change
	if observedDistance != expectedDistance {
		t.Errorf("observed distance value not equal to expected")
	}

	// ToDMS
	observedToLatDMS, observedToLongDMS := latlong1.ToDMS()
	expectedToLatDMS, expectedToLongDMS := []float64{80.0, 0.0, 0.0}, []float64{45.0, 0.0, 0.0}
	if utils.EqualSlice(observedToLatDMS, expectedToLatDMS) != true {
		t.Errorf("slices are not equal")
	}
	if utils.EqualSlice(observedToLongDMS, expectedToLongDMS) != true {
		t.Errorf("slices are not equal")
	}

	// UTMZone
	observedUTMZone := latlong1.UTMZone()
	expectedUTMZone := 34 // todo: correctly verify
	if observedUTMZone != float64(expectedUTMZone) {
		t.Errorf("utm zones not equal")
	}
}
