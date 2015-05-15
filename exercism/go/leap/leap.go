package leap

import (
	"math"
)

func IsLeapYear(year int) (answer bool) {
	year64 := float64(year)
	answer = false
	if isVanillaLeapYear(year64) {
		if isCentury(year64) {
			answer = isExceptionalCentury(year64)
		} else {
			answer = true
		}
	}
	return answer
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
