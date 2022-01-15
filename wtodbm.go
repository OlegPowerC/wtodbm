package wtodbm

import (
	"fmt"
	"math"
)

const (
	Range_W  = 1
	Range_mW = 1000
	Range_uW = 1000000
)

func roundlocal(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixedlocal(num float64, precision uint) (shortedvalue float64, err error) {
	if precision > 12 {
		return 0, fmt.Errorf("Invalid precision parameter")
	}
	output := math.Pow(10, float64(precision))
	return float64(roundlocal(num*output)) / output, nil
}

func WattTodBm(watt float64, watt_range int, digit_after_point uint) (dBm float64, Err error) {
	var Adder float64 = 0
	switch watt_range {
	case 1:
		Adder = 30
	case 1000:
		break
	case 1000000:
		Adder = -30
	default:
		return 0.0, fmt.Errorf("Error on watt range parameter")
		break
	}
	pDbm := 10*math.Log10(watt) + Adder
	rval, rerr := toFixedlocal(pDbm, digit_after_point)
	return rval, rerr
}

func DBmToWatt(dBm float64, watt_range int, digit_after_point uint) (Watt float64, Err error) {
	var Adder float64 = 0
	switch watt_range {
	case 1:
		Adder = 30
	case 1000:
		break
	case 1000000:
		Adder = -30
	default:
		return 0.0, fmt.Errorf("Error on watt range parameter")
		break
	}
	Watt_in_Watt := math.Pow(10, (dBm-Adder)/10)
	Watt, Werr := toFixedlocal(Watt_in_Watt, digit_after_point)
	return Watt, Werr
}
