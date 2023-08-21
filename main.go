package main

import (
	"calculator/calculation"
	"calculator/numerals"
	"calculator/workwithIO"
	"errors"
	"fmt"
)

func main() {
	fmt.Print("Ведите выражение: ")
	a, b, operator, err := workwithIO.GetExpression()
	if err != nil {
		workwithIO.PrintError(err)
		return
	}
	resultArabic, resultRoman, err := runCalculate(a, b, operator)
	if err != nil {
		workwithIO.PrintError(err)
	} else {
		workwithIO.PrintResult(resultArabic, resultRoman)
	}
}

func runCalculate(first, second, operator string) (int, string, error) {
	aIsRomanNumeral := false
	bIsRomanNumeral := false
	aIsInt := false
	bIsInt := false
	var a, b int
	a, aIsInt = numerals.IsArabicNumeral(first)
	b, bIsInt = numerals.IsArabicNumeral(second)
	if aIsInt && bIsInt {
		if result, err := calculation.Calculate(a, b, operator); err == nil {
			return result, "", nil
		} else {
			return 0, "", err
		}
	}
	a, aIsRomanNumeral = numerals.IsRomanNumeral(first)
	b, bIsRomanNumeral = numerals.IsRomanNumeral(second)
	if aIsRomanNumeral && bIsRomanNumeral {
		if result, err := calculation.Calculate(a, b, operator); err == nil && result > 0 {
			return 0, numerals.ArabicToRoman(result), nil
		} else if result <= 0 {
			return 0, "", errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел и нуля.")
		} else {
			return 0, "", err
		}
	}
	if aIsInt && !bIsInt || !aIsInt && bIsInt {
		return 0, "", errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}
	return 0, "", nil
}
