// Package leap contains logic to determine leap years.
package leap

// TestVersion is the version of the test suite that
// the solution references.
const TestVersion = 1

// IsLeapYear determines whether or not the given year is a leap year.
// This is only valid for the Gregorian calendar.
// While the algorithm works for years that occurred before the Gregorian reform
// in 1582, it is not valid for negative years (representing BC notation).
//
// Check out this delightful 4 minute video by CGP Gray that explains why we have leap years.
// https://www.youtube.com/watch?v=xX96xng7sAE
func IsLeapYear(y int) bool {
	divBy := func(n int) bool {
		return y%n == 0
	}
	return divBy(4) && !divBy(100) || divBy(400)
}
