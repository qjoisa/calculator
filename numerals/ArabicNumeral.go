package numerals

import "strconv"

func IsArabicNumeral(element string) (int, bool) {
	var numbers int
	var wasError error
	if numbers, wasError = strconv.Atoi(element); wasError != nil {
		return 0, false
	} else {
		return numbers, true
	}
}
