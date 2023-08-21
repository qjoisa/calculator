package numerals

func IsRomanNumeral(element string) (int, bool) {
	switch element {
	case "I":
		return 1, true
	case "II":
		return 2, true
	case "III":
		return 3, true
	case "IV":
		return 4, true
	case "V":
		return 5, true
	case "VI":
		return 6, true
	case "VII":
		return 7, true
	case "VIII":
		return 8, true
	case "IX":
		return 9, true
	case "X":
		return 10, true
	default:
		return 0, false
	}
}

func ArabicToRoman(result int) string {
	tens := result / 10
	switch {
	case tens < 1:
		return numeralsWithinTen(result % 10)
	case tens == 1:
		return "X" + numeralsWithinTen(result%10)
	case tens == 2:
		return "XX" + numeralsWithinTen(result%10)
	case tens == 3:
		return "XXX" + numeralsWithinTen(result%10)
	case tens == 4:
		return "XL" + numeralsWithinTen(result%10)
	case tens == 5:
		return "L" + numeralsWithinTen(result%10)
	case tens == 6:
		return "LX" + numeralsWithinTen(result%10)
	case tens == 7:
		return "LXX" + numeralsWithinTen(result%10)
	case tens == 8:
		return "LXXX" + numeralsWithinTen(result%10)
	case tens == 9:
		return "XC" + numeralsWithinTen(result%10)
	case tens == 10:
		return "C"
	default:
		return ""
	}
}
func numeralsWithinTen(inTen int) string {
	if inTen == 0 {
		return ""
	}
	switch {
	case inTen == 1:
		return "I"
	case inTen == 2:
		return "II"
	case inTen == 3:
		return "III"
	case inTen == 4:
		return "IV"
	case inTen == 5:
		return "V"
	case inTen == 6:
		return "VI"
	case inTen == 7:
		return "VII"
	case inTen == 8:
		return "VIII"
	case inTen == 9:
		return "IX"
	case inTen == 10:
		return "X"
	default:
		return ""
	}
}
