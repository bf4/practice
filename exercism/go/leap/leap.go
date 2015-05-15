package leap

import (
	"math"
)

func IsLeapYear(year int) bool {
	year64 := float64(year)
	if math.Mod(year64, 4.0) == 0 {
		if math.Mod(year64, 100.0) == 0 {
			return math.Mod(year64, 400.0) == 0
		} else {
			return true
		}
	}
	return false
}
