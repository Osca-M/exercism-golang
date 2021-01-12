package leap

// IsLeapYear function to check whether an year is leap or not.
func IsLeapYear(year int) bool {
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		return true
	}
	return false
}
