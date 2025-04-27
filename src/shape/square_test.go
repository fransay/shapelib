package shape

import (
	"testing"

	"github.com/fransay/shapelib/internal/utils"
)

func TestSquare(t *testing.T) {
	square := Square{Length: 3.0}

	// Area
	res := square.Area()
	exp := 9.0
	if res != exp {
		t.Errorf("Got %f, Expected %f", res, exp)
	}

	// Perimeter
	resPerim := square.Perimeter()
	expPerim := 12.0
	if resPerim != expPerim {
		t.Errorf("Got %f, Expected %f", resPerim, expPerim)
	}

	// Diagonal
	resDiag := square.Diagional()
	expDiag := 4.242641
	if !utils.IsClose(resDiag, expDiag, 0.1) {
		t.Errorf("Got %f, Expected %f", resDiag, expDiag)
	}

	// CircumReference
	resCircumRef := square.CircumRadius()
	expCircumRef := 4.242641 / 2
	if !utils.IsClose(resCircumRef, expCircumRef, 0.1) {
		t.Errorf("Got %f, Expected %f", resCircumRef, expCircumRef)
	}

	// InRadius
	resRad := square.InRadius()
	expInRad := 1.500000
	if resRad != expInRad {
		t.Errorf("Got %f, Expected %f", resRad, expInRad)
	}

	// Apothem
	apotRad := square.Apothem()
	expRad := 1.500000
	if apotRad != expRad {
		t.Errorf("Got %f, Expected %f", apotRad, expRad)
	}

	// GoldenRation
	resGoldRat := square.GoldenRation()
	expGoldRat := 1.414214

	if !utils.IsClose(resGoldRat, expGoldRat, 0.1) {
		t.Errorf("Got %f, Expected %f", resGoldRat, expGoldRat)
	}

}
