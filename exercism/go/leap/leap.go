package leap

import (
	"math"
)

func IsLeapYear(year int) bool {
	year64 := float64(year)
	if isExceptionalCentury(year64) {
		return true
	} else if isCentury(year64) {
		return false
	} else {
		return isVanillaLeapYear(year64)
	}
}

func isVanillaLeapYear(year64 float64) bool {
	return math.Mod(year64, 4.0) == 0
}

func isCentury(year64 float64) bool {
	return math.Mod(year64, 100.0) == 0
}

func isExceptionalCentury(year64 float64) bool {
	return math.Mod(year64, 400.0) == 0
}
