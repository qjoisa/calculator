package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Ведите выражение: ")
	reader := bufio.NewReader(os.Stdin)
	var expression string
	expression, _ = reader.ReadString('\n')
	expression = strings.TrimSuffix(expression, "\n")
	expression = strings.TrimSuffix(expression, "\r")
	err := start(expression)
	if err != nil {
		fmt.Println(err)
	}
}

func start(expression string) error {
	var elements []string = strings.Split(expression, " ")
	if len(elements) > 3 {
		return errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else if len(elements) < 3 {
		return errors.New("Вывод ошибки, так как строка не является математической операцией.")
	}
	if resultArabic, resultRoman, err := calculateExpression(elements); resultRoman != "" {
		fmt.Println(resultRoman)
	} else if resultRoman == "" && err == nil {
		fmt.Println(resultArabic)
	} else {
		fmt.Println(err)
	}
	return nil
}
func calculateExpression(elements []string) (int, string, error) {
	aIsRomanNumeral := false
	bIsRomanNumeral := false
	aIsInt := false
	bIsInt := false
	var a, b int
	a, aIsInt = isInt(elements[0])
	b, bIsInt = isInt(elements[2])
	if aIsInt && bIsInt {
		if result, err := calculate(a, b, elements[1]); err == nil {
			return result, "", nil
		} else {
			return 0, "", err
		}
	}
	a, aIsRomanNumeral = isRomanNumeral(elements[0])
	b, bIsRomanNumeral = isRomanNumeral(elements[2])
	if aIsRomanNumeral && bIsRomanNumeral {
		if result, err := calculate(a, b, elements[1]); err == nil && result > 0 {
			return 0, arabicToRoman(result), nil
		} else if result < 0 {
			return 0, "", errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		} else {
			return 0, "", err
		}

	}

	if aIsInt && !bIsInt || !aIsInt && bIsInt {
		return 0, "", errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}
	return 0, "", nil
}

func isInt(element string) (int, bool) {
	var numbers int
	var wasError error
	if numbers, wasError = strconv.Atoi(element); wasError != nil {
		return 0, false
	} else {
		return numbers, true
	}
}

func isRomanNumeral(element string) (int, bool) {
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

func calculate(a, b int, action string) (int, error) {
	if (a > 10 && a < 0) || (b > 10 && b < 0) {
		return 0, errors.New("Некорректный ввод - переменная должна быть от 1 до 10 включительно")
	}
	switch action {
	case "+":
		return addition(a, b), nil
	case "-":
		return subtraction(a, b), nil
	case "*":
		return multiplication(a, b), nil
	case "/":
		if b != 0 {
			return division(a, b), nil
		} else {
			return 0, errors.New("На 0 делить нельзя")
		}
	default:
		return 0, errors.New("Невозможное действие")
	}
}
func addition(a, b int) int {
	return a + b
}
func subtraction(a, b int) int {
	return a - b
}
func multiplication(a, b int) int {
	return a * b
}
func division(a, b int) int {
	return a / b
}
func arabicToRoman(result int) string {
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
