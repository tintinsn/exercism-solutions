// Package leap provides a function to determine whether a given year is a leap year.
package leap

// IsLeapYear takes year as input amd returns ture if it is a leap year,
// Otherwise returns false.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	return year%4 == 0
}
