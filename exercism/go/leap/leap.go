package leap

import (
	"math"
)

func IsLeapYear(year int) (answer bool) {
	year64 := float64(year)
	answer = false
	if math.Mod(year64, 4.0) == 0 {
		if math.Mod(year64, 100.0) == 0 {
			answer = math.Mod(year64, 400.0) == 0
		} else {
			answer = true
		}
	}
	return answer
}
