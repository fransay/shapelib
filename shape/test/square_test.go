package test

import (
	"shapelib/shape"
	"testing"
)

func TestSquare(t *testing.T) {
	square := shape.Square{Length: 3.0}

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
	expDiag := 3.0
	if resDiag != expDiag {
		t.Errorf("Got %f, Expected %f", resDiag, expDiag)
	}

	// CircumReference
	resCircumRef := square.CircumRadius()
	expCircumRef := 3.0
	if resCircumRef != expCircumRef {
		t.Errorf("Got %f, Expected %f", resCircumRef, expCircumRef)
	}

	// InRadius
	resRad := square.InRadius()
	expInRad := 3.0
	if resRad != expInRad {
		t.Errorf("Got %f, Expected %f", resRad, expInRad)
	}

	// Apothem
	apotRad := square.Apothem()
	expRad := 3.0
	if apotRad != expRad {
		t.Errorf("Got %f, Expected %f", apotRad, expRad)
	}

	// GoldenRation
	resGoldRat := square.GoldenRation()
	expGoldRat := 3.0
	if resGoldRat != expGoldRat {
		t.Errorf("Got %f, Expected %f", resGoldRat, expGoldRat)
	}

}
