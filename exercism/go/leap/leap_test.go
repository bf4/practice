package leap

import "testing"

// Define a function IsLeapYear(int) bool.
//

// Retired testVersions
// (none) 4a9e144a3c5dc0d9773f4cf641ffe3efe48641d8

func TestLeapYears(t *testing.T) {
	for _, test := range testCases {
		observed := IsLeapYear(test.year)
		if observed != test.expected {
			t.Fatalf("IsLeapYear(%d) = %t, want %t (%s)",
				test.year, observed, test.expected, test.description)
		}
	}
}

func BenchmarkLeapYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			IsLeapYear(test.year)
		}
	}
}
