package leap

const testVersion = 2

// IsLeapYear calculates if the given year is a leap year or not and returns the bool result
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 != 0 {
			return true
		} else if year%400 == 0 {
			return true
		}
	}
	return false
}

/*func IsLeapYear(year int) bool {
	if year%100 != 0 && year%4 == 0 {
		return true
	} else if year%100 == 0 && year%4 == 0 && year%400 == 0 {
		return true
	}
	return false
}*/

/*func IsLeapYear(year int) bool {
	if year%100 != 0 {
		if year%4 == 0 {
			return true
		}
	} else {
		if year%4 == 0 && year%400 == 0 {
			return true
		}
	}
	return false
}
*/
