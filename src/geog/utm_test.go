package geog

import "testing"

func TestUTM(t *testing.T) {
	utm := UTM{
		Easting:  6000000,
		Northing: 3000000,
	}

	utm1 := UTM{
		Easting:  2000000,
		Northing: 4000000,
	}

	// distance
	observedDistance := utm.Distance(utm1)
	expectedDistance := 0.0 // todo: fix 0.0
	if observedDistance != expectedDistance {
		t.Errorf("distances are not equal")
	}

}
