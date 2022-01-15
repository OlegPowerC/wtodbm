package wtodbm

import (
	"testing"
)

//uW: 1 dBm: -30
//dBm: -30 uW: 1
//mW: 1 dBm: 0
//dBm: 0 mW: 1
//W: 1 dBm: 30
//dBm: 30 W: 1

func TestWattTodBm(t *testing.T) {
	WattTodBmval, WattToDbmErr := WattTodBm(1, Range_W, 2)
	if WattToDbmErr != nil {
		t.Errorf("WattTodBm return error: %s", WattToDbmErr)
	}
	if WattTodBmval != 30 {
		t.Errorf("Expected -30, got: %.f", WattTodBmval)
	}

	mWattTodBm, WattToDbmErr := WattTodBm(1, Range_mW, 2)
	if WattToDbmErr != nil {
		t.Errorf("WattTodBm return error: %s", WattToDbmErr)
	}
	if mWattTodBm != 0 {
		t.Errorf("Expected 0, got: %.f", mWattTodBm)
	}

	uWattTodBm, WattToDbmErr := WattTodBm(1, Range_uW, 2)
	if WattToDbmErr != nil {
		t.Errorf("WattTodBm return error: %s", WattToDbmErr)
	}
	if uWattTodBm != -30 {
		t.Errorf("Expected -30, got: %.f", uWattTodBm)
	}

	dBmtoWatt, DbmtoWattErr := DBmToWatt(30, Range_W, 2)
	if WattToDbmErr != nil {
		t.Errorf("DBmToWatt return error: %s", DbmtoWattErr)
	}
	if dBmtoWatt != 1 {
		t.Errorf("Expected 1, got: %.f", dBmtoWatt)
	}

	dBmtomWatt, DbmtomWattErr := DBmToWatt(0, Range_mW, 2)
	if WattToDbmErr != nil {
		t.Errorf("DBmToWatt return error: %s", DbmtomWattErr)
	}
	if dBmtomWatt != 1 {
		t.Errorf("Expected 1, got: %.f", dBmtomWatt)
	}

	dBmtouWatt, DbmtouWattErr := DBmToWatt(-30, Range_uW, 2)
	if WattToDbmErr != nil {
		t.Errorf("DBmToWatt return error: %s", DbmtouWattErr)
	}
	if dBmtouWatt != 1 {
		t.Errorf("Expected 1, got: %.f", dBmtouWatt)
	}
}
